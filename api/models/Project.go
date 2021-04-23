package models

import "time"

type Project struct {
	ID          string    `json:"id"`
	ProjectName string    `json:"ProjectName"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
