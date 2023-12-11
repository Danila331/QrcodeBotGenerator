package models

type Update struct {
	UpadateId int     `json:"update_id"`
	Message   Message `json:"message"`
}

type RestResponse struct {
	Result []Update `json:"result"`
}
