package core

import "github.com/vatusa/training/internal/data"

func SyncController(cid uint, name string, email string, rating int, facilityCode string) (*data.Controller, error) {
	controller, err := data.GetController(cid)
	// TODO: Check if err is not exists
	if controller == nil {
		controller = &data.Controller{
			CID:                 cid,
			Name:                name,
			Email:               email,
			NetworkRating:       rating,
			ATCRating:           0,
			Facility:            nil,
			LastPromotion:       nil,
			LastTrainingSession: nil,
			BasicStageStatus:    0,
			LocalStageStatus:    0,
			ApproachStageStatus: 0,
			EnrouteStageStatus:  0,
			CurrencyStatus:      0,
		}
	}
}
