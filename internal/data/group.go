package data

type Group struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"size:120"`
	Facility *Facility
	Teachers []GroupTeacher
}

type TeacherStatus = uint

const (
	TeacherStatus_Active TeacherStatus = iota
	TeacherStatus_Inactive
)

type GroupTeacher struct {
	ID            uint `gorm:"primaryKey"`
	Group         *Group
	Controller    *Controller
	TeacherStatus TeacherStatus
}
