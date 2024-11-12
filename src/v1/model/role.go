package model

import "gorm.io/gorm"

type Role struct {
	gorm.Model   `json:"-"`
	Name         string     `gorm:"column:name;unique;not null"`
	Description  string     `gorm:"column:description;not null"`
	Privileges   Privileges `gorm:"many2many:role_privilege"`
	Users        Users      `gorm:"many2many:user_role"`
	PrivilegeIds []uint     `gorm:"-"`
}

type RoleCreation struct {
	Name         *string `json:"name,omitempty" binding:"required"`
	Description  *string `json:"description,omitempty" binding:"required"`
	PrivilegeIds []uint  `json:"privilegeIds,omitempty" binding:"required"`
}

type RoleUpdate struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

type RoleResponse struct {
	Name        string             `json:"name,omitempty"`
	Description string             `json:"description,omitempty"`
	Privileges  PrivilegeResponses `json:"privileges,omitempty"`
}

type Roles []Role
type RoleResponses []RoleResponse

func (Role) TableName() string          { return "roles" }
func (Roles) TableName() string         { return Role{}.TableName() }
func (RoleCreation) TableName() string  { return Role{}.TableName() }
func (RoleUpdate) TableName() string    { return Role{}.TableName() }
func (RoleResponse) TableName() string  { return Role{}.TableName() }
func (RoleResponses) TableName() string { return Role{}.TableName() }

func (p RoleCreation) AsModel() *Role {
	return &Role{Model: gorm.Model{}, Name: *p.Name, Description: *p.Description, PrivilegeIds: p.PrivilegeIds}
}

func (p Role) AsResponse() *RoleResponse {
	return &RoleResponse{Name: p.Name, Description: p.Description, Privileges: p.Privileges.AsCollectionResponse()}
}

func (p Roles) AsCollectionResponse() RoleResponses {
	result := RoleResponses{}
	for _, model := range p {
		response := model.AsResponse()
		result = append(result, *response)
	}
	return result
}
