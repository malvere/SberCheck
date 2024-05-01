package model

import "time"

type ResponseModel struct {
	Success bool `json:"success"`
	Meta    struct {
		Time       time.Time `json:"time"`
		TraceID    string    `json:"traceId"`
		RequestID  string    `json:"requestId"`
		AppVersion string    `json:"appVersion"`
	} `json:"meta"`
	Errors     []interface{} `json:"errors"`
	PromoCodes []struct {
		Key           string `json:"key"`
		Description   string `json:"description"`
		Amount        int    `json:"amount"`
		MinOrderPrice int    `json:"minOrderPrice"`
		StartAt       struct {
			Year    int `json:"year"`
			Month   int `json:"month"`
			Day     int `json:"day"`
			Hours   int `json:"hours"`
			Minutes int `json:"minutes"`
			Seconds int `json:"seconds"`
			Nanos   int `json:"nanos"`
		} `json:"startAt"`
		ExpireAt struct {
			Year    int `json:"year"`
			Month   int `json:"month"`
			Day     int `json:"day"`
			Hours   int `json:"hours"`
			Minutes int `json:"minutes"`
			Seconds int `json:"seconds"`
			Nanos   int `json:"nanos"`
		} `json:"expireAt"`
		LinkToTerms string `json:"linkToTerms"`
	} `json:"promoCodes"`
}
