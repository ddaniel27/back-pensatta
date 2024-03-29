package domain

import "time"

type User struct {
	Username      string    `json:"username,omitempty"    gorm:"unique;not null"`
	Password      string    `json:"password,omitempty"    gorm:"not null"`
	FirstName     string    `json:"first_name"            gorm:"not null"`
	LastName      string    `json:"last_name"             gorm:"not null"`
	ListNumber    uint64    `json:"list_number"           gorm:"not null"`
	Role          string    `json:"role"                  gorm:"not null"`
	InstitutionID string    `json:"institution_code"      gorm:"not null,name:institucion_id"`
	DateJoined    time.Time `json:"date_joined,omitempty" gorm:"not null"`
}
