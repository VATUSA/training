package data

type Facility struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"size:120"`
	Code string `gorm:"size:4"`
}

var facilityCache = NewBulkCache[string, Facility](fetchFacilities, facilityExtractKey, 300)

func GetFacilityByCode(code string) (*Facility, error) {
	return facilityCache.Get(code)
}

func fetchFacilities() ([]Facility, error) {
	var facilities []Facility
	result := DB.Model(Facility{}).Find(facilities)
	return facilities, result.Error
}

func facilityExtractKey(facility Facility) string {
	return facility.Code
}
