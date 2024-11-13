package token

import (
	"errors"
	"iam/src/v1/abstraction"
	"iam/src/v1/config"
	"iam/src/v1/constant"
	"iam/src/v1/domain"
	"iam/src/v1/domain/dto"
	"iam/src/v1/exception"
	"iam/src/v1/storage"
	"os"
	"strconv"
	"strings"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/samber/lo"
)

type jwtAuthProvider struct {
	badCredentialS abstraction.BadCredentialS
	Token          dto.Token
}

func NewJWTAuthProvider(appCtx config.AppContext) abstraction.JwtAuthProvider[dto.AuthClaims, domain.UserResponse] {
	badCredentialS := storage.NewGormBadCredentialS(appCtx)
	accessTokenSecret := os.Getenv("ACCESS_TOKEN_SECRET")
	accessTokenTimeToLive, _ := strconv.Atoi(os.Getenv("ACCESS_TOKEN_TIME_TO_LIVE"))
	refreshTokenSecret := os.Getenv("REFRESH_TOKEN_SECRET")
	refreshTokenTimeToLive, _ := strconv.Atoi(os.Getenv("ACCESS_TOKEN_TIME_TO_LIVE"))
	token := dto.Token{AccessTokenSecret: accessTokenSecret, AccessTokenTimeToLive: accessTokenTimeToLive, RefreshTokenSecret: refreshTokenSecret, RefreshTokenTimeToLive: refreshTokenTimeToLive}
	return jwtAuthProvider{badCredentialS: badCredentialS, Token: token}
}

func (p jwtAuthProvider) BuildScope(user domain.UserResponse) string {
	var scope = []string{}
	if len(user.Roles) > 0 {
		lo.ForEach(user.Roles, func(r domain.RoleResponse, i int) {
			scope = append(scope, "ROLE_"+r.Name)
			if len(r.Privileges) > 0 {
				lo.ForEach(r.Privileges, func(p domain.PrivilegeResponse, i int) {
					scope = append(scope, p.Name)
				})
			}
		})
	}
	return strings.Join(scope, " ")
}

func (p jwtAuthProvider) Sign(user domain.UserResponse, expiredTime int, secretKey string, id string, referId string) (*string, exception.ServiceException) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, &dto.AuthClaims{
		Payload: dto.Payload{UserId: user.ID, Scope: p.BuildScope(user), ReferId: referId},
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

func (p jwtAuthProvider) Verify(token string, secretKey string) (*dto.AuthClaims, exception.ServiceException) {
	t, parseErr := jwt.ParseWithClaims(token, &dto.AuthClaims{}, func(t *jwt.Token) (interface{}, error) { return []byte(secretKey), nil })
	if parseErr != nil {
		if errors.Is(parseErr, jwt.ErrTokenExpired) {
			return nil, exception.NewServiceException(parseErr, constant.JwtTokenExpiredF)
		}
		return nil, exception.NewServiceException(parseErr, constant.ParseJwtTokenF)
	}
	if claims, ok := t.Claims.(*dto.AuthClaims); ok {
		return claims, nil
	}
	return nil, exception.NewServiceException(nil, constant.IllLegalJwtTokenF)
}

func (p jwtAuthProvider) EnsureNotBadCredential(token string) (*dto.AuthClaims, exception.ServiceException) {
	claims, verifiedErr := p.Verify(token, p.Token.AccessTokenSecret)
	if verifiedErr != nil {
		return nil, verifiedErr
	}
	accessTokenID := claims.ID
	badCredential, queriedErr := p.badCredentialS.FindByAccessTokenId(accessTokenID)
	if badCredential != nil {
		return nil, exception.NewServiceException(nil, constant.TokenBlockedF)
	}
	isNotBadCredential := queriedErr != nil && queriedErr.GetFailed() == constant.FindByIdNoContentF
	if isNotBadCredential {
		return claims, nil
	}
	return nil, exception.NewServiceException(queriedErr, constant.EnsureTokenNotBadCredentialF)
}

func (p jwtAuthProvider) GetConstant() dto.Token {
	return p.Token
}
