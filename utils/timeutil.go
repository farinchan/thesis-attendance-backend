package utils

import (
	"time"
)

func GetIndonesianDayName(t time.Time) string {
	days := map[time.Weekday]string{
		time.Sunday:    "minggu",
		time.Monday:    "senin",
		time.Tuesday:   "selasa",
		time.Wednesday: "rabu",
		time.Thursday:  "kamis",
		time.Friday:    "jumat",
		time.Saturday:  "sabtu",
	}

	return days[t.Weekday()]
}

func GetCurrentDate() string {
	return time.Now().Format("2006-01-02")
}

func GetCurrentTime() string {
	return time.Now().Format("15:04:05")
}
