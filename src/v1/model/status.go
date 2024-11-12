package model

import "gorm.io/gorm"

type Status struct {
	gorm.Model
	Content string `gorm:"column:content;unique;not null"`
	UserID  uint
	User    *User
}

type StatusCreation struct {
	Content *string `json:"content,omitempty" binding:"required"`
	UserId  *uint   `json:"userId,omitempty" binding:"required"`
}

type StatusUpdate struct {
	Content *string `json:"content,omitempty"`
}

type StatusResponse struct {
	gorm.Model
	Content string       `json:"content,omitempty"`
	User    UserResponse `json:"user,omitempty"`
}

type Statuses []Status
type StatusResponses []StatusResponse

func (Status) TableName() string          { return "statuses" }
func (Statuses) TableName() string        { return Status{}.TableName() }
func (StatusCreation) TableName() string  { return Status{}.TableName() }
func (StatusUpdate) TableName() string    { return Status{}.TableName() }
func (StatusResponse) TableName() string  { return Status{}.TableName() }
func (StatusResponses) TableName() string { return Status{}.TableName() }

func (p StatusCreation) AsModel() *Status {
	return &Status{Model: gorm.Model{}, Content: *p.Content, UserID: *p.UserId}
}

func (p Status) AsResponse() *StatusResponse {
	return &StatusResponse{Model: p.Model, Content: p.Content, User: *p.User.AsResponse()}
}

func (p Statuses) AsCollectionResponse() StatusResponses {
	result := StatusResponses{}
	for _, model := range p {
		response := model.AsResponse()
		result = append(result, *response)
	}
	return result
}
