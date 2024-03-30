package models

import "time"

type UserModel struct {
	ID            uint      `gorm:"primary_key"`
	Username      string    `gorm:"unique;not null"`
	Password      string    `gorm:"not null"`
	FirstName     string    `gorm:"not null"`
	LastName      string    `gorm:"not null"`
	Role          string    `gorm:"not null"`
	InstitutionID uint64    `gorm:"not null"`
	ListNumber    uint64    `gorm:"not null"`
	DateJoined    time.Time `gorm:"not null"`
	LastLogin     time.Time
}

func (UserModel) TableName() string {
	return "pensatta_user"
}
