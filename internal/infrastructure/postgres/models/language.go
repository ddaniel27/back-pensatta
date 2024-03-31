package models

type LanguageModel struct {
	ID            uint64 `gorm:"primary_key"`
	InstitutionID uint64 `gorm:"not null"`
	Value         string `gorm:"not null"`
}

func (LanguageModel) TableName() string {
	return "pensatta_languages"
}
