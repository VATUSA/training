package data

import "time"

type ControllerAvailability struct {
	ID         uint64 `gorm:"primaryKey"`
	Controller *Controller
	StartTime  *time.Time
	EndTime    *time.Time
}
