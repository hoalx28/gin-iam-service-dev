package dto

import jwt "github.com/golang-jwt/jwt/v5"

type Payload struct {
	UserId  uint   `json:"userId,omitempty"`
	ReferId string `json:"referId,omitempty"`
	Scope   string `json:"scope,omitempty"`
}

type AuthClaims struct {
	Payload Payload `json:"payload"`
	jwt.RegisteredClaims
}

type Token struct {
	AccessTokenSecret      string
	AccessTokenTimeToLive  int
	RefreshTokenSecret     string
	RefreshTokenTimeToLive int
}
