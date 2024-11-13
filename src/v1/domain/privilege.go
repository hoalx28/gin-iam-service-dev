package domain

import "gorm.io/gorm"

type Privilege struct {
	gorm.Model
	Name        string `gorm:"column:name;unique;not null"`
	Description string `gorm:"column:description;not null"`
	Roles       *Roles `gorm:"many2many:role_privilege;"`
}

type PrivilegeCreation struct {
	Name        *string `json:"name,omitempty" binding:"required"`
	Description *string `json:"description,omitempty" binding:"required"`
}

type PrivilegeUpdate struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

type PrivilegeResponse struct {
	gorm.Model
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type Privileges []Privilege
type PrivilegeResponses []PrivilegeResponse

func (Privilege) TableName() string          { return "privileges" }
func (Privileges) TableName() string         { return Privilege{}.TableName() }
func (PrivilegeCreation) TableName() string  { return Privilege{}.TableName() }
func (PrivilegeUpdate) TableName() string    { return Privilege{}.TableName() }
func (PrivilegeResponse) TableName() string  { return Privilege{}.TableName() }
func (PrivilegeResponses) TableName() string { return Privilege{}.TableName() }

func (p PrivilegeCreation) AsModel() Privilege {
	return Privilege{Model: gorm.Model{}, Name: *p.Name, Description: *p.Description}
}

func (p Privilege) AsResponse() PrivilegeResponse {
	return PrivilegeResponse{Model: p.Model, Name: p.Name, Description: p.Description}
}

func (p Privileges) AsCollectionResponse() PrivilegeResponses {
	result := PrivilegeResponses{}
	for _, model := range p {
		response := model.AsResponse()
		result = append(result, response)
	}
	return result
}
