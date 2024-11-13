package dto

import (
	"iam/src/v1/domain"

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
	Password string `json:"-"`
	Roles    domain.RoleResponses
}

func (r RegisterRequest) AsUserCreation() domain.UserCreation {
	return domain.UserCreation{Username: &r.Username, Password: &r.Password, RoleIds: r.RoleIds}
}

func (r *RegisterResponse) FromUserResponse(user domain.UserResponse) {
	r.Model = user.Model
	r.Username = user.Username
	r.Password = user.Password
	r.Roles = user.Roles
}
