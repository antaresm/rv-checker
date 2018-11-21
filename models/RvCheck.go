package models

type RvCheck struct {
	Id 			int    	`json:"id"`
	Name 		string 	`json:"name"`
	AppID 		string 	`json:"app_id"`
	StoreType	string 	`json:"store_type"`
	Country		string 	`json:"country"`
	LastStatus 	int  	`json:"last_status"`
}

func (RvCheck) TableName() string {
	return "checkpoint"
}