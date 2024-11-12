package dto

import (
	"iam/src/v1/model"

	"gorm.io/gorm"
)

type RegisterRequest struct {
	Username string `json:"username,omitempty" binding:"required"`
	Password string `json:"password,omitempty" binding:"required"`
	RoleIds  []uint `json:"roleIds,omitempty" binding:"required"`
}

type RegisterResponse struct {
	gorm.Model
	Username string `json:"username,omitempty"`
	Password string `json:"-,omitempty"`
	Roles    model.RoleResponses
}

func (r RegisterRequest) AsUserCreation() *model.UserCreation {
	return &model.UserCreation{Username: &r.Username, Password: &r.Password, RoleIds: r.RoleIds}
}

func (r *RegisterResponse) FromUserResponse(user model.UserResponse) {
	r.Model = user.Model
	r.Username = user.Username
	r.Password = user.Password
	r.Roles = user.Roles
}
