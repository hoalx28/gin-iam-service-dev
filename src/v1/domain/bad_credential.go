package domain

import (
	"time"

	"gorm.io/gorm"
)

type BadCredential struct {
	gorm.Model
	AccessTokenId        string    `gorm:"column:access_token_id;unique;not null"`
	AccessTokenExpiredAt time.Time `gorm:"column:access_token_expired_at;not null"`
	UserId               uint      `gorm:"column:user_id;not null"`
}

type BadCredentialCreation struct {
	AccessTokenId        *string    `json:"accessTokenId,omitempty"`
	AccessTokenExpiredAt *time.Time `json:"accessTokenExpiredAt,omitempty"`
	UserId               *uint      `json:"userId,omitempty"`
}

type BadCredentialResponse struct {
	gorm.Model
	AccessTokenId        string    `json:"accessTokenId,omitempty"`
	AccessTokenExpiredAt time.Time `json:"accessTokenExpiredAt,omitempty"`
	UserId               uint      `json:"userId,omitempty"`
}

type BadCredentials []BadCredential
type BadCredentialResponses []BadCredentialResponse

func (BadCredential) TableName() string          { return "bad_credentials" }
func (BadCredentials) TableName() string         { return BadCredential{}.TableName() }
func (BadCredentialCreation) TableName() string  { return BadCredential{}.TableName() }
func (BadCredentialResponse) TableName() string  { return BadCredential{}.TableName() }
func (BadCredentialResponses) TableName() string { return BadCredential{}.TableName() }

func (p BadCredentialCreation) AsModel() BadCredential {
	return BadCredential{Model: gorm.Model{}, AccessTokenId: *p.AccessTokenId, AccessTokenExpiredAt: *p.AccessTokenExpiredAt, UserId: *p.UserId}
}

func (p BadCredential) AsResponse() BadCredentialResponse {
	return BadCredentialResponse{Model: p.Model, AccessTokenId: p.AccessTokenId, AccessTokenExpiredAt: p.AccessTokenExpiredAt, UserId: p.ID}
}

func (p BadCredentials) AsCollectionResponse() BadCredentialResponses {
	result := BadCredentialResponses{}
	for _, model := range p {
		response := model.AsResponse()
		result = append(result, response)
	}
	return result
}
