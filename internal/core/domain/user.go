package domain

import "time"

type User struct {
	Username      string    `json:"username,omitempty"`
	Password      string    `json:"password,omitempty"`
	FirstName     string    `json:"first_name"`
	LastName      string    `json:"last_name"`
	ListNumber    uint64    `json:"list_number"`
	Role          string    `json:"role"`
	InstitutionID string    `json:"institution_id"`
	DateJoined    time.Time `json:"date_joined,omitempty"`
}
