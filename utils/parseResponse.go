package utils

import (
	"strconv"
	"strings"

	infraestructure "github.com/WelintonJunior/obd-diagnostic-service/infraestructure/bluetooth"
)

func ParseRPM(response string) int {
	parts := strings.Fields(response)
	if len(parts) < 4 {
		return 0
	}
	a, _ := strconv.ParseInt(parts[2], 16, 64)
	b, _ := strconv.ParseInt(parts[3], 16, 64)
	return int((a*256 + b) / 4)
}

func ParseSpeedKPH(response string) int {
	parts := strings.Fields(response)
	if len(parts) < 3 {
		return 0
	}
	v, _ := strconv.ParseInt(parts[2], 16, 64)
	return int(v)
}

func ParseCoolantTemp(response string) float64 {
	parts := strings.Fields(response)
	if len(parts) < 3 {
		return 0
	}
	v, _ := strconv.ParseInt(parts[2], 16, 64)
	return float64(v) - 40
}

func ParseThrottlePosition(response string) float64 {
	parts := strings.Fields(response)
	if len(parts) < 3 {
		return 0
	}
	v, _ := strconv.ParseInt(parts[2], 16, 64)
	return float64(v) * 100.0 / 255.0
}

func ParseMAF(response string) float64 {
	parts := strings.Fields(response)
	if len(parts) < 4 {
		return 0
	}
	a, _ := strconv.ParseInt(parts[2], 16, 64)
	b, _ := strconv.ParseInt(parts[3], 16, 64)
	return float64(a*256+b) / 100.0 // g/s
}

func ParseFuelPressure(response string) float64 {
	parts := strings.Fields(response)
	if len(parts) < 3 {
		return 0
	}
	v, _ := strconv.ParseInt(parts[2], 16, 64)
	return float64(v) * 3.0 // kPa
}

func ParseBatteryVoltage(response string) float64 {
	resp := infraestructure.SendCommand("ATRV")
	resp = strings.TrimSuffix(resp, "V")
	v, err := strconv.ParseFloat(resp, 64)
	if err != nil {
		return 0
	}
	return v
}

func ParseEngineRunTime(response string) int {
	parts := strings.Fields(response)
	if len(parts) < 4 {
		return 0
	}
	a, _ := strconv.ParseInt(parts[2], 16, 64)
	b, _ := strconv.ParseInt(parts[3], 16, 64)
	return int(a*256 + b)
}
