package domain

import "gorm.io/gorm"

// ! One To One not working right
type Device struct {
	gorm.Model
	IpAddress string `gorm:"column:ip_address;unique;not null"`
	UserAgent string `gorm:"column:user_agent;not null"`
	UserID    uint
	User      *User
}

type DeviceCreation struct {
	IpAddress *string `json:"ipAddress,omitempty" binding:"required"`
	UserAgent *string `json:"userAgent,omitempty" binding:"required"`
	UserID    *uint   `json:"userId,omitempty" binding:"required"`
}

type DeviceUpdate struct {
	IpAddress *string `json:"ipAddress,omitempty"`
	UserAgent *string `json:"userAgent,omitempty"`
}

type DeviceResponse struct {
	gorm.Model
	IpAddress string `json:"ipAddress,omitempty"`
	UserAgent string `json:"userAgent,omitempty"`
}

type Devices []Device
type DeviceResponses []DeviceResponse

func (Device) TableName() string          { return "devices" }
func (Devices) TableName() string         { return Device{}.TableName() }
func (DeviceCreation) TableName() string  { return Device{}.TableName() }
func (DeviceUpdate) TableName() string    { return Device{}.TableName() }
func (DeviceResponse) TableName() string  { return Device{}.TableName() }
func (DeviceResponses) TableName() string { return Device{}.TableName() }

func (p DeviceCreation) AsModel() Device {
	return Device{Model: gorm.Model{}, IpAddress: *p.IpAddress, UserAgent: *p.UserAgent, UserID: *p.UserID}
}

func (p Device) AsResponse() DeviceResponse {
	return DeviceResponse{Model: p.Model, IpAddress: p.IpAddress, UserAgent: p.UserAgent}
}

func (p Devices) AsCollectionResponse() DeviceResponses {
	result := DeviceResponses{}
	for _, model := range p {
		response := model.AsResponse()
		result = append(result, response)
	}
	return result
}
