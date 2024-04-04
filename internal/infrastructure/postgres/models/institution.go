package models

type CalificationModel struct {
	ID           uint64  `gorm:"primaryKey"`
	AverageScore float64 `gorm:"not null"`
	AverageTime  float64 `gorm:"not null"`
	UserID       uint64  `gorm:"not null"`
}

func (CalificationModel) TableName() string {
	return "pensatta_calification"
}
