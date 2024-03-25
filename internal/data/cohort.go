package data

import "time"

type Cohort struct {
	ID                    uint   `gorm:"primaryKey"`
	Name                  string `gorm:"size:120"`
	Pool                  *Pool
	Group                 *Group
	StartDate             *time.Time
	StudentSlots          uint
	StudentSlotsRemaining uint
	Students              []CohortStudent
}

type StudentStatus uint

const (
	StudentStatus_Pending StudentStatus = iota
	StudentStatus_Complete
	StudentStatus_Incomplete
)

type CohortStudent struct {
	ID         uint `gorm:"primaryKey"`
	Cohort     *Cohort
	Controller *Controller
}
