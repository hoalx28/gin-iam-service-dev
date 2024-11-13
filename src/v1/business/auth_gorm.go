package business

import (
	"iam/src/v1/abstraction"
	"iam/src/v1/config"
	"iam/src/v1/constant"
	"iam/src/v1/domain"
	"iam/src/v1/domain/dto"
	"iam/src/v1/exception"
	"iam/src/v1/token"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type gormAuthB struct {
	userB           abstraction.UserB
	badCredentialB  abstraction.BadCredentialB
	jwtAuthProvider abstraction.JwtAuthProvider[dto.AuthClaims, domain.UserResponse]
}

func NewGormAuthB(appCtx config.AppContext) abstraction.AuthB {
	userB := NewGormUserB(appCtx)
	badCredentialB := NewGormBadCredentialB(appCtx)
	jwtAuthProvider := token.NewJWTAuthProvider(appCtx)
	return gormAuthB{userB: userB, jwtAuthProvider: jwtAuthProvider, badCredentialB: badCredentialB}
}

func (b gormAuthB) newCredentials(user domain.UserResponse) (*dto.CredentialResponse, exception.ServiceException) {
	accessTokenId := uuid.NewString()
	refreshTokenId := uuid.NewString()
	accessTokenSecret := b.jwtAuthProvider.GetConstant().AccessTokenSecret
	accessTokenTimeToLive := b.jwtAuthProvider.GetConstant().AccessTokenTimeToLive
	refreshTokenSecret := b.jwtAuthProvider.GetConstant().RefreshTokenSecret
	refreshTokenTimeToLive := b.jwtAuthProvider.GetConstant().RefreshTokenTimeToLive
	accessToken, signAccessTokenErr := b.jwtAuthProvider.Sign(user, accessTokenTimeToLive, accessTokenSecret, accessTokenId, refreshTokenId)
	if signAccessTokenErr != nil {
		return nil, exception.NewServiceException(signAccessTokenErr, constant.SignJwtTokenF)
	}
	accessTokenIssuedAt := time.Now().Unix()
	refreshToken, signRefreshTokenErr := b.jwtAuthProvider.Sign(user, refreshTokenTimeToLive, refreshTokenSecret, refreshTokenId, accessTokenId)
	if signRefreshTokenErr != nil {
		return nil, exception.NewServiceException(signAccessTokenErr, constant.SignJwtTokenF)
	}
	refreshTokenIssuedAt := time.Now().Unix()
	credential := dto.NewCredentialResponse(*accessToken, int(accessTokenIssuedAt), *refreshToken, int(refreshTokenIssuedAt))
	return &credential, nil
}

func (b gormAuthB) asBadCredentialCacheCreation(accessToken string) (*domain.BadCredentialCreation, exception.ServiceException) {
	claims, parseErr := b.Identity(accessToken)
	if parseErr != nil {
		return nil, parseErr
	}
	accessTokenId := claims.ID
	accessTokenExpiredAt := claims.ExpiresAt.Time
	userId := claims.Payload.UserId
	badCredential := domain.BadCredentialCreation{AccessTokenId: &accessTokenId, AccessTokenExpiredAt: &accessTokenExpiredAt, UserId: &userId}
	return &badCredential, nil
}

func (b gormAuthB) SignUp(request dto.RegisterRequest) (*dto.CredentialResponse, exception.ServiceException) {
	creation := request.AsUserCreation()
	user, signUpErr := b.userB.SaveB(&creation)
	if signUpErr != nil {
		return nil, signUpErr
	}
	credential, newCredentialErr := b.newCredentials(*user)
	if newCredentialErr != nil {
		return nil, newCredentialErr
	}
	return credential, nil
}

func (b gormAuthB) SignIn(request dto.CredentialRequest) (*dto.CredentialResponse, exception.ServiceException) {
	user, queriedErr := b.userB.FindByUsernameB(request.Username)
	if queriedErr != nil {
		return nil, exception.NewServiceException(queriedErr, constant.BadCredentialF)
	}
	authorizationErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if authorizationErr != nil {
		return nil, exception.NewServiceException(nil, constant.BadCredentialF)
	}
	credential, newCredentialErr := b.newCredentials(*user)
	if newCredentialErr != nil {
		return nil, newCredentialErr
	}
	return credential, nil
}

func (b gormAuthB) Identity(accessToken string) (*dto.AuthClaims, exception.ServiceException) {
	claims, verifiedErr := b.jwtAuthProvider.EnsureNotBadCredential(accessToken)
	if verifiedErr != nil {
		return nil, verifiedErr
	}
	return claims, nil
}

func (b gormAuthB) Me(accessToken string) (*dto.RegisterResponse, exception.ServiceException) {
	claims, verifiedErr := b.Identity(accessToken)
	if verifiedErr != nil {
		return nil, verifiedErr
	}
	username := claims.Subject
	user, queriedErr := b.userB.FindByUsernameB(username)
	if queriedErr != nil {
		return nil, exception.NewServiceException(queriedErr, constant.RetrieveProfileF)
	}
	register := &dto.RegisterResponse{}
	register.FromUserResponse(*user)
	return register, nil
}

func (b gormAuthB) SignOut(accessToken string) (*uint, exception.ServiceException) {
	creation, parseErr := b.asBadCredentialCacheCreation(accessToken)
	if parseErr != nil {
		return nil, parseErr
	}
	saved, savedErr := b.badCredentialB.SaveB(creation)
	if savedErr != nil {
		return nil, exception.NewServiceException(savedErr, constant.SignOutF)
	}
	return &saved.ID, nil
}

func (b gormAuthB) Refresh(accessToken string, refreshToken string) (*dto.CredentialResponse, exception.ServiceException) {
	refreshTokenSecret := b.jwtAuthProvider.GetConstant().RefreshTokenSecret
	accessClaims, accessVerifiedErr := b.Identity(accessToken)
	if accessVerifiedErr != nil {
		return nil, accessVerifiedErr
	}
	refreshClaims, refreshVerifiedErr := b.jwtAuthProvider.Verify(refreshToken, refreshTokenSecret)
	if refreshVerifiedErr != nil {
		return nil, refreshVerifiedErr
	}
	accessTokenReferId := accessClaims.Payload.ReferId
	refreshTokenId := refreshClaims.ID
	if accessTokenReferId != refreshTokenId {
		return nil, exception.NewServiceException(nil, constant.JwtTokenNotSuitableF)
	}
	creation, parseErr := b.asBadCredentialCacheCreation(accessToken)
	if parseErr != nil {
		return nil, parseErr
	}
	_, savedErr := b.badCredentialB.SaveB(creation)
	if savedErr != nil {
		return nil, exception.NewServiceException(savedErr, constant.RecallJwtTokenF)
	}
	user, queriedErr := b.userB.FindByUsernameB(accessClaims.Subject)
	if queriedErr != nil {
		return nil, exception.NewServiceException(queriedErr, constant.RefreshTokenF)
	}
	credential, newCredentialErr := b.newCredentials(*user)
	if newCredentialErr != nil {
		return nil, newCredentialErr
	}
	return credential, nil
}
