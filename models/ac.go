package models

type AC struct {
	UserId string `json:"user_id"`
	QueId  string `json:"question_id"`
	IfAc   bool   `json:"if_ac"`
}
