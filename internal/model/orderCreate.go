package model

import "time"

type OrderCreateRequest struct {
	Identification struct {
		ID string `json:"id"`
	} `json:"identification"`
	PaymentType                  string          `json:"paymentType"`
	Customer                     CustomerProfile `json:"customer"`
	Flags                        []any           `json:"flags"`
	IsSelectedCartItemGroupsOnly bool            `json:"isSelectedCartItemGroupsOnly"`
	Discounts                    []OrderDiscount `json:"discounts"`
	PaymentTypeOptions           []any           `json:"paymentTypeOptions"`
	Auth                         AuthFiled       `json:"auth"`
}

type OrderDiscount struct {
	Type    string `json:"type"` // PROMO_CODE
	Voucher string `json:"voucher"`
}

type CustomerProfile struct {
	NotMe        bool   `json:"notMe"` // false
	ThirdName    string `json:"thirdName"`
	Comment      string `json:"comment"`
	FirstName    string `json:"firstName"` // Володя
	LastName     string `json:"lastName"`  // Интеров
	Email        string `json:"email"`     // TLOTOL.OLTO@MAIL.RU
	Phone        string `json:"phone"`     // 79013912737
	PhoneMisc    string `json:"phoneMisc"`
	Restored     bool   `json:"restored"`
	BonusBalance int64  `json:"bonusBalance"`
}

type OrderCreateResponse struct {
	Success bool `json:"success"`
	Meta    struct {
		Time       time.Time `json:"time"`
		TraceID    string    `json:"traceId"`
		RequestID  string    `json:"requestId"`
		AppVersion string    `json:"appVersion"`
	} `json:"meta"`
	Errors []struct {
		Title  string `json:"title"`
		Detail string `json:"detail"`
		Field  string `json:"field"`
		Code   int    `json:"code"`
	} `json:"errors"`
	Message string `json:"message"`
	OrderID string `json:"orderId"`
}
