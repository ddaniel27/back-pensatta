package domain

import "time"

type User struct {
	ID              uint64    `json:"id"`
	Username        string    `json:"username,omitempty"`
	Password        string    `json:"password,omitempty"`
	FirstName       string    `json:"first_name"`
	LastName        string    `json:"last_name"`
	ListNumber      uint64    `json:"list_number"`
	Role            string    `json:"role"`
	InstitutionID   uint64    `json:"institution_id"`
	InstitutionCode string    `json:"institution_code,omitempty"`
	Language        string    `json:"language,omitempty"`
	DateJoined      time.Time `json:"date_joined,omitempty"`
	LastLogin       time.Time `json:"last_login,omitempty"`
}
