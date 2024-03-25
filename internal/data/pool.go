package data

type Pool struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"size:120"`
	Facility *Facility
}

type PoolStudent struct {
	ID         uint `gorm:"primaryKey"`
	Pool       *Pool
	Controller *Controller
}

type PoolGroup struct {
	ID    uint `gorm:"primaryKey"`
	Pool  *Pool
	Group *Group
}
