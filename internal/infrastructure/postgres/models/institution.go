package models

type InstitutionModel struct {
	ID       uint64 `gorm:"primary_key"`
	Name     string `gorm:"unique;not null"`
	Email    string `gorm:"not null"`
	Country  string `gorm:"not null"`
	Province string `gorm:"not null"`
	City     string `gorm:"not null"`
	Code     string `gorm:"unique;not null"`
}

func (InstitutionModel) TableName() string {
	return "pensatta_institution"
}
