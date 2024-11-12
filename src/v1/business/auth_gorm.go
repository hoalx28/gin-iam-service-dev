package business

import (
	"fmt"
	"iam/src/v1/config"
	"iam/src/v1/constant"
	"iam/src/v1/exception"
	"iam/src/v1/model"
	"iam/src/v1/model/dto"
	"iam/src/v1/token"
	"os"
	"strconv"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type gormAuthBusiness struct {
	userBusiness          UserBusiness
	badCredentialBusiness BadCredentialBusiness
	jwtAuthProvider       token.TokenAuthProvider[token.AuthClaims, model.UserResponse]
}

func NewGormAuthBusiness(appCtx config.AppContext) AuthBusiness {
	userBusiness := NewGormUserBusiness(appCtx)
	badCredentialBusiness := NewGormBadCredentialBusiness(appCtx)
	jwtAuthProvider := token.NewJWTAuthProvider(appCtx)
	return gormAuthBusiness{userBusiness: userBusiness, jwtAuthProvider: jwtAuthProvider, badCredentialBusiness: badCredentialBusiness}
}

func (b gormAuthBusiness) newCredentials(user *model.UserResponse) (*dto.CredentialResponse, exception.ServiceException) {
	accessTokenSecret := os.Getenv("ACCESS_TOKEN_SECRET")
	accessTokenTimeToLive, _ := strconv.Atoi(os.Getenv("ACCESS_TOKEN_TIME_TO_LIVE"))
	refreshTokenSecret := os.Getenv("REFRESH_TOKEN_SECRET")
	refreshTokenTimeToLive, _ := strconv.Atoi(os.Getenv("REFRESH_TOKEN_TIME_TO_LIVE"))
	accessTokenId := uuid.NewString()
	refreshTokenId := uuid.NewString()
	accessToken, signAccessTokenErr := b.jwtAuthProvider.Sign(*user, accessTokenTimeToLive, accessTokenSecret, accessTokenId, refreshTokenId)
	if signAccessTokenErr != nil {
		return nil, exception.NewServiceException(signAccessTokenErr, constant.SignJwtTokenF)
	}
	accessTokenIssuedAt := time.Now().Unix()
	refreshToken, signRefreshTokenErr := b.jwtAuthProvider.Sign(*user, refreshTokenTimeToLive, refreshTokenSecret, refreshTokenId, accessTokenId)
	if signRefreshTokenErr != nil {
		return nil, exception.NewServiceException(signAccessTokenErr, constant.SignJwtTokenF)
	}
	refreshTokenIssuedAt := time.Now().Unix()
	credential := dto.NewCredentialResponse(*accessToken, int(accessTokenIssuedAt), *refreshToken, int(refreshTokenIssuedAt))
	return &credential, nil
}

func (b gormAuthBusiness) asBadCredentialCacheCreation(accessToken string) (*model.BadCredentialCreation, exception.ServiceException) {
	claims, parseErr := b.Identity(accessToken)
	if parseErr != nil {
		return nil, parseErr
	}
	accessTokenId := claims.ID
	accessTokenExpiredAt := claims.ExpiresAt.Time
	userId := claims.Payload.UserId
	badCredential := model.BadCredentialCreation{AccessTokenId: &accessTokenId, AccessTokenExpiredAt: &accessTokenExpiredAt, UserId: &userId}
	return &badCredential, nil
}

func (b gormAuthBusiness) SignUp(request dto.RegisterRequest) (*dto.CredentialResponse, exception.ServiceException) {
	creation := request.AsUserCreation()
	user, signUpErr := b.userBusiness.SaveBusiness(creation)
	if signUpErr != nil {
		return nil, signUpErr
	}
	credential, newCredentialErr := b.newCredentials(user)
	if newCredentialErr != nil {
		return nil, newCredentialErr
	}
	return credential, nil
}

func (b gormAuthBusiness) SignIn(request dto.CredentialRequest) (*dto.CredentialResponse, exception.ServiceException) {
	user, queriedErr := b.userBusiness.FindByUsernameBusiness(request.Username)
	if queriedErr != nil {
		return nil, exception.NewServiceException(queriedErr, constant.BadCredentialF)
	}
	authorizationErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if authorizationErr != nil {
		return nil, exception.NewServiceException(nil, constant.BadCredentialF)
	}
	credential, newCredentialErr := b.newCredentials(user)
	if newCredentialErr != nil {
		return nil, newCredentialErr
	}
	return credential, nil
}

func (b gormAuthBusiness) Identity(accessToken string) (*token.AuthClaims, exception.ServiceException) {
	claims, verifiedErr := b.jwtAuthProvider.EnsureNotBadCredential(accessToken)
	if verifiedErr != nil {
		return nil, verifiedErr
	}
	return claims, nil
}

func (b gormAuthBusiness) Me(accessToken string) (*dto.RegisterResponse, exception.ServiceException) {
	claims, verifiedErr := b.Identity(accessToken)
	fmt.Println(verifiedErr)
	if verifiedErr != nil {
		return nil, verifiedErr
	}
	username := claims.Subject
	user, queriedErr := b.userBusiness.FindByUsernameBusiness(username)
	if queriedErr != nil {
		return nil, exception.NewServiceException(queriedErr, constant.RetrieveProfileF)
	}
	register := &dto.RegisterResponse{}
	register.FromUserResponse(*user)
	return register, nil
}

func (b gormAuthBusiness) SignOut(accessToken string) (*uint, exception.ServiceException) {
	creation, parseErr := b.asBadCredentialCacheCreation(accessToken)
	if parseErr != nil {
		return nil, parseErr
	}
	saved, savedErr := b.badCredentialBusiness.SaveBusiness(creation)
	if savedErr != nil {
		return nil, exception.NewServiceException(savedErr, constant.SignOutF)
	}
	return &saved.ID, nil
}

func (b gormAuthBusiness) Refresh(accessToken string, refreshToken string) exception.ServiceException {
	panic("not implemented") // TODO: Implement
}
