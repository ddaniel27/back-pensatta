package domain

import "errors"

var ErrInvalidLanguage = errors.New("invalid language")

type Institution struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Country  string `json:"country"`
	Province string `json:"province"`
	City     string `json:"city"`
	Code     string `json:"code"`
	Language string `json:"language"`
}
