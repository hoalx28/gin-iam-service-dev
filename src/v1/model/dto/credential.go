package dto

type CredentialRequest struct {
	Username string `json:"username,omitempty" binding:"required"`
	Password string `json:"password,omitempty" binding:"required"`
}

type CredentialResponse struct {
	AccessToken          string `json:"access_token,omitempty"`
	AccessTokenIssuedAt  int    `json:"access_token_issued_at,omitempty"`
	RefreshToken         string `json:"refresh_token,omitempty"`
	RefreshTokenIssuedAt int    `json:"refresh_token_issued_at,omitempty"`
}

func NewCredentialResponse(accessToken string, accessTokenIssuedAt int, refreshToken string, refreshTokenIssuedAt int) CredentialResponse {
	return CredentialResponse{AccessToken: accessToken, AccessTokenIssuedAt: accessTokenIssuedAt, RefreshToken: refreshToken, RefreshTokenIssuedAt: refreshTokenIssuedAt}
}
