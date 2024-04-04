package models

type CalificationModel struct {
	ID           uint64 `gorm:"primaryKey"`
	AverageScore uint64 `gorm:"not null"`
	AverageTime  uint64 `gorm:"not null"`
	UserID       uint64 `gorm:"not null"`
}

func (CalificationModel) TableName() string {
	return "pensatta_calification"
}

func (c *CalificationModel) NewAverages(score float64, time uint64, totalRecords int64) {
	c.AverageScore = (c.AverageScore*uint64(totalRecords-1) + uint64(score)) / uint64(totalRecords)
	c.AverageTime = (c.AverageTime*uint64(totalRecords-1) + time) / uint64(totalRecords)
}
