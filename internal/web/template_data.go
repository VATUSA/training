package web

import "github.com/vatusa/training/internal/data"

type MetaData struct {
	Facilities []data.Facility
	Ratings    []Rating
}

func MakeMetaData() MetaData {
	return MetaData{
		Facilities: nil,
		Ratings:    nil,
	}
}

type Rating struct {
	Int   int
	Short string
	Long  string
}

type UserData struct {
	CID      uint64
	Name     string
	Rating   Rating
	ACLFlags UserACLFlags
}

type UserACLFlags struct {
	Administrator bool
}
