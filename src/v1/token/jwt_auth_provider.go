package token

import (
	"errors"
	"iam/src/v1/config"
	"iam/src/v1/constant"
	"iam/src/v1/exception"
	"iam/src/v1/model"
	"iam/src/v1/storage"
	"os"
	"strings"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/samber/lo"
)

type payload struct {
	UserId  uint   `json:"userId,omitempty"`
	ReferId string `json:"referId,omitempty"`
	Scope   string `json:"scope,omitempty"`
}

type AuthClaims struct {
	Payload payload `json:"payload"`
	jwt.RegisteredClaims
}

func NewAuthClaims(payload payload) AuthClaims {
	return AuthClaims{Payload: payload}
}

type jwtAuthProvider struct {
	badCredentialStorage storage.BadCredentialStorage
}

func NewJWTAuthProvider(appCtx config.AppContext) TokenAuthProvider[AuthClaims, model.UserResponse] {
	badCredentialStorage := storage.NewGormBadCredentialStorage(appCtx)
	return jwtAuthProvider{badCredentialStorage: badCredentialStorage}
}

func (p jwtAuthProvider) BuildScope(user model.UserResponse) string {
	var scope = []string{}
	if len(user.Roles) > 0 {
		lo.ForEach(user.Roles, func(r model.RoleResponse, i int) {
			scope = append(scope, "ROLE_"+r.Name)
			if len(r.Privileges) > 0 {
				lo.ForEach(r.Privileges, func(p model.PrivilegeResponse, i int) {
					scope = append(scope, p.Name)
				})
			}
		})
	}
	return strings.Join(scope, " ")
}

func (p jwtAuthProvider) Sign(user model.UserResponse, expiredTime int, secretKey string, id string, referId string) (*string, exception.ServiceException) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, &AuthClaims{
		Payload: payload{UserId: user.ID, Scope: p.BuildScope(user), ReferId: referId},
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   user.Username,
			ID:        id,
			ExpiresAt: &jwt.NumericDate{Time: time.Now().Local().Add(time.Second * time.Duration(expiredTime))},
			IssuedAt:  &jwt.NumericDate{Time: time.Now().Local()},
		},
	})
	token, signErr := t.SignedString([]byte(secretKey))
	if signErr != nil {
		return nil, exception.NewServiceException(signErr, constant.SignJwtTokenF)
	}
	return &token, nil
}

func (p jwtAuthProvider) Verify(token string, secretKey string) (*AuthClaims, exception.ServiceException) {
	t, parseErr := jwt.ParseWithClaims(token, &AuthClaims{}, func(t *jwt.Token) (interface{}, error) { return []byte(secretKey), nil })
	if parseErr != nil {
		if errors.Is(parseErr, jwt.ErrTokenExpired) {
			return nil, exception.NewServiceException(parseErr, constant.JwtTokenExpiredF)
		}
		return nil, exception.NewServiceException(parseErr, constant.ParseJwtTokenF)
	}
	if claims, ok := t.Claims.(*AuthClaims); ok {
		return claims, nil
	}
	return nil, exception.NewServiceException(nil, constant.IllLegalJwtTokenF)
}

func (p jwtAuthProvider) EnsureNotBadCredential(token string) (*AuthClaims, exception.ServiceException) {
	accessTokenSecret := os.Getenv("ACCESS_TOKEN_SECRET")
	claims, verifiedErr := p.Verify(token, accessTokenSecret)
	if verifiedErr != nil {
		return nil, verifiedErr
	}
	accessTokenID := claims.ID
	badCredential, queriedErr := p.badCredentialStorage.FindByAccessTokenId(accessTokenID)
	if badCredential != nil {
		return nil, exception.NewServiceException(nil, constant.TokenBlockedF)
	}
	isNotBadCredential := queriedErr != nil && queriedErr.GetFailed() == constant.FindByIdNoContentF
	if isNotBadCredential {
		return claims, nil
	}
	return nil, exception.NewServiceException(queriedErr, constant.EnsureTokenNotBadCredentialF)
}
