package domain

import "gopkg.in/mgo.v2/bson"

type (
	Register struct {
		Id bson.ObjectId "_id"
		Key string "key"
		LastId string "lastId"
	}
	Status struct {
		Status string
	}
	CallVO struct{
		Items []Call "items"
		LastId bson.ObjectId "lastId"
		Pageable SimplePageable "pageable"
	}
	SimplePageable struct {
		Total int "total"
		Limit int "limit"
		Remain int "remain"
		Pages int "pages"
	}
	Call struct {
		Id bson.ObjectId "_id"
		CallDuration int32 "call_duration"
		TranscriptionText string "transcription_text"
	}
)



