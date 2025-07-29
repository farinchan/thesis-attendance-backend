package model

import "time"

type StudentAttendanceRule struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	Day       string    `json:"day" validate:"required"`                    // e.g., "senin", "selasa", etc.
	Start     time.Time `gorm:"type:time" json:"start" validate:"required"` // format: "HH:MM"
	End       time.Time `gorm:"type:time" json:"end" validate:"required"`   // format: "HH:MM"
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	// Additional fields can be added as needed
}

