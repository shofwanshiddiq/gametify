package models

import (
	"time"

	"gorm.io/gorm"
)

/*
RoomId uint
RoomName string
Description string
Town string
Address string
Price int
Availability bool
ConsoleType string
PackageType string
OpenHrsStart string
OpenHrsEnd string
Image string
Ratings int

*/

type ConsoleType string

const (
	TypePC  ConsoleType = "PC"
	TypePS5 ConsoleType = "PS5"
	TypePS4 ConsoleType = "PS4"
)

type Room struct {
	gorm.Model
	PlaceID      uint        `json:"place_id"`
	Availability bool        `json:"availability"`
	ConsoleType  ConsoleType `json:"console_type" gorm:"type:varchar(20)"`
}

type Place struct {
	gorm.Model
	Name    string `json:"name"`
	Town    string `json:"town"`
	Address string `json:"address"`
	Image   string `json:"image"`
	Ratings int    `json:"ratings"`

	Rooms []Room `gorm:"foreignKey:PlaceID" json:"rooms"`
}

type Package struct {
	gorm.Model
	RoomID       uint      `json:"room_id"`
	PackageType  string    `json:"package_type" gorm:"type:varchar(20)"`
	Price        float32   `json:"price"`
	OpenHrsStart time.Time `json:"open_hrs_start" gorm:"type:time"`
	OpenHrsEnd   time.Time `json:"open_hrs_end" gorm:"type:time"`
}
