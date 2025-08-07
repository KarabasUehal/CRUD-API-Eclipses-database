package models

type Eclipse struct {
	ID          int     `json:"id" gorm:"primaryKey;autoIncrement"`
	Description string  `json:"description"`
	Duration    float64 `json:"duration"`
}
