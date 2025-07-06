package types

import "time"

type ReadingSession struct {
	Base

	SessionID string    `gorm:"uniqueIndex;not null" json:"session_id"`
	StartTime time.Time `gorm:"not null" json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	TripName  string    `json:"trip_name"`
}
