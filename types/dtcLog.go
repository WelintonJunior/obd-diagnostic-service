package types

import "time"

type DTCLog struct {
	Base

	TimestampDetected time.Time  `gorm:"not null" json:"timestamp_detected"`
	DTCCode           string     `gorm:"not null" json:"dtc_code"`
	Description       string     `json:"description"`
	Status            string     `gorm:"type:enum('ativo','resolvido','apagado');default:'ativo'" json:"status"`
	ClearedAt         *time.Time `json:"cleared_at,omitempty"` // pode ser nulo
}
