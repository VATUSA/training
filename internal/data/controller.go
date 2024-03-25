package data

import "time"

type StageStatus = uint

const (
	StageStatus_NotEligible StageStatus = iota
	StageStatus_Eligible
	StageStatus_PendingWritten
	StageStatus_PendingCohort
	StageStatus_CohortAssigned
	StageStatus_CohortNoShow
)

type CurrencyStatus = uint

const (
	CurrencyStatus_NoRating CurrencyStatus = iota
	CurrencyStatus_Active
	CurrencyStatus_RecentlyActive
	CurrencyStatus_NotCurrent
	CurrencyStatus_Exception
)

type Controller struct {
	CID                 uint   `gorm:"primaryKey"`
	Name                string `gorm:"size:120"`
	Email               string `gorm:"size:120"`
	NetworkRating       int
	ATCRating           int
	Facility            *Facility
	LastPromotion       *time.Time
	LastTrainingSession *time.Time
	BasicStageStatus    StageStatus
	LocalStageStatus    StageStatus
	ApproachStageStatus StageStatus
	EnrouteStageStatus  StageStatus
	CurrencyStatus      CurrencyStatus
}

var controllerCache = NewSingleCache(fetchControllerByCID, controllerExtractKey, 300)

func GetController(cid uint) (*Controller, error) {
	return controllerCache.Get(cid)
}

func fetchControllerByCID(cid uint) (*Controller, error) {
	var controller *Controller
	result := DB.Model(Controller{}).Where("cid", cid).First(controller)

	if result.Error != nil {
		return nil, result.Error
	}
	return controller, nil
}

func controllerExtractKey(controller Controller) uint {
	return controller.CID
}
