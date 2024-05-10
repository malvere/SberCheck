package model

type CartSearchResponse struct {
	Elements []struct {
		Identification          CartIdentification `json:"identification"`
		ItemGroups              []any              `json:"itemGroups"`
		Errors                  []any              `json:"errors"`
		Discounts               []any              `json:"discounts"`
		PrivilegesObtained      []any              `json:"privilegesObtained"`
		TotalSum                int                `json:"totalSum"`
		TotalFinalSum           int                `json:"totalFinalSum"`
		TotalDisplaySum         int                `json:"totalDisplaySum"`
		TotalOldPriceSum        int                `json:"totalOldPriceSum"`
		TotalPriceChange        int                `json:"totalPriceChange"`
		TotalEconomy            int                `json:"totalEconomy"`
		MinCartSum              int                `json:"minCartSum"`
		MinCartSumLack          int                `json:"minCartSumLack"`
		CreatedAt               any                `json:"createdAt"`
		UpdatedAt               any                `json:"updatedAt"`
		TotalOffers             int                `json:"totalOffers"`
		TotalMerchants          int                `json:"totalMerchants"`
		UserBonusBalance        int                `json:"userBonusBalance"`
		Type                    string             `json:"type"`
		LocationID              string             `json:"locationId"`
		LocationName            string             `json:"locationName"`
		LocationAddress         string             `json:"locationAddress"`
		Success                 bool               `json:"success"`
		CashBonusInfo           any                `json:"cashBonusInfo"`
		OnlineBonusInfo         any                `json:"onlineBonusInfo"`
		IsBpg2DiscountAvailable bool               `json:"isBpg2DiscountAvailable"`
		ClientAddress           any                `json:"clientAddress"`
		DeliveryDiscountPrice   int                `json:"deliveryDiscountPrice"`
		ServiceScheme           int                `json:"serviceScheme"`
		MinCartSumForDelivery   int                `json:"minCartSumForDelivery"`
		Pochta                  any                `json:"pochta"`
		ObtainedBonusInfo       []any              `json:"obtainedBonusInfo"`
		PossibleBonusInfo       []any              `json:"possibleBonusInfo"`
	} `json:"elements"`
}

type CartIdentification struct {
	ID string `json:"id"`
}
