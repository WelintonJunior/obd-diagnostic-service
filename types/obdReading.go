package types

import "time"

type OBDReading struct {
	Base

	Timestamp time.Time `gorm:"not null" json:"timestamp"`
	SessionID string    `gorm:"index" json:"session_id"`

	RPM                 int     `json:"rpm"`
	SpeedKPH            int     `json:"speed_kph"`
	CoolantTempC        float64 `json:"coolant_temp_c"`
	ThrottlePositionPct float64 `json:"throttle_position_pct"`
	MAFGps              float64 `json:"maf_gps"`
	FuelPressureKpa     float64 `json:"fuel_pressure_kpa"`
	BatteryVoltage      float64 `json:"battery_voltage"`
	FuelTrimShort       float64 `json:"fuel_trim_short"`
	FuelTrimLong        float64 `json:"fuel_trim_long"`
	EngineRunTimeSec    int     `json:"engine_run_time_sec"`
}
