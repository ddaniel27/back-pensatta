package domain

import "time"

type User struct {
	ID            uint64    `json:"id"                 gorm:"primary_key,auto_increment"`
	Username      string    `json:"username"           gorm:"unique;not null"`
	Password      string    `json:"password,omitempty" gorm:"not null"`
	FirstName     string    `json:"first_name"         gorm:"not null"`
	LastName      string    `json:"last_name"          gorm:"not null"`
	ListNumber    uint64    `json:"list_number"        gorm:"not null"`
	Role          string    `json:"role"               gorm:"not null"`
	InstitutionID uint64    `json:"institution_id"     gorm:"not null,name:institucion_id"`
	DateJoined    time.Time `json:"date_joined"        gorm:"not null"`
}
