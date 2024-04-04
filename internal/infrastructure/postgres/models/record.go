package models

import "time"

type RecordModel struct {
	ID         uint64    `gorm:"primaryKey"`
	Score      uint64    `gorm:"not null"`
	Time       uint64    `gorm:"not null"`
	Date       time.Time `gorm:"not null"`
	ExerciseID uint64    `gorm:"not null"`
	UserID     uint64    `gorm:"not null"`
}

func (RecordModel) TableName() string {
	return "pensatta_record"
}
