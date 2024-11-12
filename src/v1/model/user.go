package model

import "gorm.io/gorm"

type User struct {
	gorm.Model `json:"-"`
	Username   string `gorm:"column:username;unique;not null"`
	Password   string `gorm:"column:password;not null"`
	Roles      Roles  `gorm:"many2many:user_role"`
	Device     *Device
	Statuses   Statuses
	RoleIds    []uint `gorm:"-"`
}

type UserCreation struct {
	Username *string `json:"username,omitempty" binding:"required"`
	Password *string `json:"password,omitempty" binding:"required"`
	RoleIds  []uint  `json:"roleIds,omitempty" binding:"required"`
}

type UserUpdate struct {
	Username *string `json:"username,omitempty"`
	Password *string `json:"password,omitempty"`
}

type UserResponse struct {
	Username string        `json:"username,omitempty"`
	Password string        `json:"password,omitempty"`
	Roles    RoleResponses `json:"roles,omitempty"`
}

type Users []User
type UserResponses []UserResponse

func (User) TableName() string          { return "users" }
func (Users) TableName() string         { return User{}.TableName() }
func (UserCreation) TableName() string  { return User{}.TableName() }
func (UserUpdate) TableName() string    { return User{}.TableName() }
func (UserResponse) TableName() string  { return User{}.TableName() }
func (UserResponses) TableName() string { return User{}.TableName() }

func (p UserCreation) AsModel() *User {
	return &User{Model: gorm.Model{}, Username: *p.Username, Password: *p.Password, RoleIds: p.RoleIds}
}

func (p User) AsResponse() *UserResponse {
	return &UserResponse{Username: p.Username, Password: p.Password, Roles: p.Roles.AsCollectionResponse()}
}

func (p Users) AsCollectionResponse() UserResponses {
	result := UserResponses{}
	for _, model := range p {
		response := model.AsResponse()
		result = append(result, *response)
	}
	return result
}
