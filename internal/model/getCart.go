package model

import "time"

// Cart ID
type GetCartRequest struct {
	Identification struct {
		ID string `json:"id"`
	} `json:"identification"`
	IsCartStateValidationRequired bool `json:"isCartStateValidationRequired"`
	IsSelectedItemGroupsOnly      bool `json:"isSelectedItemGroupsOnly"`
	LoyaltyCalculationRequired    bool `json:"loyaltyCalculationRequired"`
	IsSkipPersonalDiscounts       bool `json:"isSkipPersonalDiscounts"`
	Auth                          AuthFiled
}

type GetCartResponse struct {
	Identification struct {
		ID string `json:"id"`
	} `json:"identification"`
	ItemGroups []struct {
		ID    string `json:"id"`
		Goods struct {
			GoodsID    string `json:"goodsId"`
			Title      string `json:"title"`
			TitleImage string `json:"titleImage"`
			Attributes []any  `json:"attributes"`
			WebURL     string `json:"webUrl"`
			Slug       string `json:"slug"`
			Boxes      []struct {
				Box           string  `json:"box"`
				PackagingUnit string  `json:"packagingUnit"`
				Width         int     `json:"width"`
				Height        int     `json:"height"`
				Length        int     `json:"length"`
				WeightUnit    string  `json:"weightUnit"`
				Weight        float64 `json:"weight"`
			} `json:"boxes"`
			CategoryID       string   `json:"categoryId"`
			Brand            string   `json:"brand"`
			ContentFlags     []string `json:"contentFlags"`
			ContentFlagsStr  []string `json:"contentFlagsStr"`
			Stocks           int      `json:"stocks"`
			PhotosCount      int      `json:"photosCount"`
			OffersCount      int      `json:"offersCount"`
			LogisticTags     []string `json:"logisticTags"`
			Images           []any    `json:"images"`
			Documents        []any    `json:"documents"`
			Description      string   `json:"description"`
			Videos           []any    `json:"videos"`
			MainCollectionID string   `json:"mainCollectionId"`
			Package          struct {
				PackageType string `json:"packageType"`
				MinQuantity string `json:"minQuantity"`
				WeightUnit  string `json:"weightUnit"`
			} `json:"package"`
		} `json:"goods"`
		Offer struct {
			ID                       string `json:"id"`
			Price                    int    `json:"price"`
			Score                    int    `json:"score"`
			IsFavorite               bool   `json:"isFavorite"`
			MerchantID               string `json:"merchantId"`
			DeliveryPossibilities    []any  `json:"deliveryPossibilities"`
			FinalPrice               int    `json:"finalPrice"`
			BonusPercent             int    `json:"bonusPercent"`
			BonusAmount              int    `json:"bonusAmount"`
			AvailableQuantity        int    `json:"availableQuantity"`
			BonusAmountFinalPrice    int    `json:"bonusAmountFinalPrice"`
			Discounts                []any  `json:"discounts"`
			PriceAdjustments         []any  `json:"priceAdjustments"`
			DeliveryDate             string `json:"deliveryDate"`
			PickupDate               string `json:"pickupDate"`
			MerchantOfferID          string `json:"merchantOfferId"`
			MerchantName             string `json:"merchantName"`
			MerchantLogoURL          string `json:"merchantLogoUrl"`
			MerchantURL              string `json:"merchantUrl"`
			MerchantSummaryRating    int    `json:"merchantSummaryRating"`
			IsBpgByMerchant          bool   `json:"isBpgByMerchant"`
			Nds                      int    `json:"nds"`
			OldPrice                 int    `json:"oldPrice"`
			OldPriceChangePercentage int    `json:"oldPriceChangePercentage"`
			MaxDeliveryDays          any    `json:"maxDeliveryDays"`
			BpgType                  string `json:"bpgType"`
			CreditPaymentAmount      int    `json:"creditPaymentAmount"`
			InstallmentPaymentAmount int    `json:"installmentPaymentAmount"`
			ShowMerchant             any    `json:"showMerchant"`
			SalesBlockInfo           any    `json:"salesBlockInfo"`
			DueDate                  string `json:"dueDate"`
			DueDateText              string `json:"dueDateText"`
			LocationID               string `json:"locationId"`
			SpasiboIsAvailable       bool   `json:"spasiboIsAvailable"`
			IsShowcase               bool   `json:"isShowcase"`
			LoyaltyPromotionFlags    []any  `json:"loyaltyPromotionFlags"`
			PricesPerMeasurement     []any  `json:"pricesPerMeasurement"`
			AvailablePaymentMethods  []any  `json:"availablePaymentMethods"`
			SuperPrice               int    `json:"superPrice"`
			WarehouseID              string `json:"warehouseId"`
			BnplPaymentParams        any    `json:"bnplPaymentParams"`
			InstallmentPaymentParams any    `json:"installmentPaymentParams"`
			BonusInfoGroups          []any  `json:"bonusInfoGroups"`
			CalculatedDeliveryDate   string `json:"calculatedDeliveryDate"`
		} `json:"offer"`
		Merchant struct {
			ID                    string `json:"id"`
			Name                  string `json:"name"`
			MinDeliveryDate       string `json:"minDeliveryDate"`
			PartnerID             string `json:"partnerId"`
			LegalInfo             any    `json:"legalInfo"`
			Design                any    `json:"design"`
			Slug                  string `json:"slug"`
			URL                   string `json:"url"`
			SiteURL               string `json:"siteUrl"`
			ConfirmationTimeLimit string `json:"confirmationTimeLimit"`
			FullName              string `json:"fullName"`
			Categories            []any  `json:"categories"`
		} `json:"merchant"`
		GoodsIDAlt         string `json:"goodsIdAlt"`
		Quantity           int    `json:"quantity"`
		Price              int    `json:"price"`
		FinalPrice         int    `json:"finalPrice"`
		DisplayPrice       int    `json:"displayPrice"`
		Sum                int    `json:"sum"`
		FinalSum           int    `json:"finalSum"`
		DisplaySum         int    `json:"displaySum"`
		OldPriceSum        int    `json:"oldPriceSum"`
		FinalWeight        string `json:"finalWeight"`
		ItemGroupID        string `json:"itemGroupId"`
		Discounts          []any  `json:"discounts"`
		PrivilegesObtained []any  `json:"privilegesObtained"`
		Errors             []any  `json:"errors"`
		AdditionalData     string `json:"additionalData"`
		Collection         any    `json:"collection"`
		Replacements       []any  `json:"replacements"`
		CashBonusInfo      struct {
			ChargedBonus           int    `json:"chargedBonus"`
			ChargedBonusPercent    int    `json:"chargedBonusPercent"`
			AdditionalAmount       int    `json:"additionalAmount"`
			AdditionalPercent      int    `json:"additionalPercent"`
			ChargedBonusPerOneItem int    `json:"chargedBonusPerOneItem"`
			BonusType              string `json:"bonusType"`
		} `json:"cashBonusInfo"`
		OnlineBonusInfo struct {
			ChargedBonus           int    `json:"chargedBonus"`
			ChargedBonusPercent    int    `json:"chargedBonusPercent"`
			AdditionalAmount       int    `json:"additionalAmount"`
			AdditionalPercent      int    `json:"additionalPercent"`
			ChargedBonusPerOneItem int    `json:"chargedBonusPerOneItem"`
			BonusType              string `json:"bonusType"`
		} `json:"onlineBonusInfo"`
		CrossedPrice      int   `json:"crossedPrice"`
		ObtainedBonusInfo []any `json:"obtainedBonusInfo"`
		IsSelected        bool  `json:"isSelected"`
		Deliveries        []struct {
			Type     string `json:"type"`
			DateFrom struct {
				Year  int `json:"year"`
				Month int `json:"month"`
				Day   int `json:"day"`
			} `json:"dateFrom"`
			DateTo          any    `json:"dateTo"`
			TimeFrom        any    `json:"timeFrom"`
			TimeTo          any    `json:"timeTo"`
			DateTimeString  string `json:"dateTimeString"`
			DeliveryOptions []any  `json:"deliveryOptions"`
		} `json:"deliveries"`
		PossibleBonusInfo []struct {
			Type    string `json:"type"`
			Amount  int    `json:"amount"`
			Percent int    `json:"percent"`
			Context string `json:"context"`
		} `json:"possibleBonusInfo"`
		Notifications []any `json:"notifications"`
	} `json:"itemGroups"`
	Errors             []any     `json:"errors"`
	Discounts          []any     `json:"discounts"`
	PrivilegesObtained []any     `json:"privilegesObtained"`
	TotalSum           int       `json:"totalSum"`
	TotalFinalSum      int       `json:"totalFinalSum"`
	TotalDisplaySum    int       `json:"totalDisplaySum"`
	TotalOldPriceSum   int       `json:"totalOldPriceSum"`
	TotalPriceChange   int       `json:"totalPriceChange"`
	TotalEconomy       int       `json:"totalEconomy"`
	MinCartSum         int       `json:"minCartSum"`
	MinCartSumLack     int       `json:"minCartSumLack"`
	CreatedAt          time.Time `json:"createdAt"`
	UpdatedAt          time.Time `json:"updatedAt"`
	TotalOffers        int       `json:"totalOffers"`
	TotalMerchants     int       `json:"totalMerchants"`
	UserBonusBalance   int       `json:"userBonusBalance"`
	Type               string    `json:"type"`
	LocationID         string    `json:"locationId"`
	LocationName       string    `json:"locationName"`
	LocationAddress    string    `json:"locationAddress"`
	Success            bool      `json:"success"`
	CashBonusInfo      struct {
		ChargedBonus           int    `json:"chargedBonus"`
		ChargedBonusPercent    int    `json:"chargedBonusPercent"`
		AdditionalAmount       int    `json:"additionalAmount"`
		AdditionalPercent      int    `json:"additionalPercent"`
		ChargedBonusPerOneItem int    `json:"chargedBonusPerOneItem"`
		BonusType              string `json:"bonusType"`
	} `json:"cashBonusInfo"`
	OnlineBonusInfo struct {
		ChargedBonus           int    `json:"chargedBonus"`
		ChargedBonusPercent    int    `json:"chargedBonusPercent"`
		AdditionalAmount       int    `json:"additionalAmount"`
		AdditionalPercent      int    `json:"additionalPercent"`
		ChargedBonusPerOneItem int    `json:"chargedBonusPerOneItem"`
		BonusType              string `json:"bonusType"`
	} `json:"onlineBonusInfo"`
	IsBpg2DiscountAvailable bool  `json:"isBpg2DiscountAvailable"`
	ClientAddress           any   `json:"clientAddress"`
	DeliveryDiscountPrice   int   `json:"deliveryDiscountPrice"`
	ServiceScheme           int   `json:"serviceScheme"`
	MinCartSumForDelivery   int   `json:"minCartSumForDelivery"`
	Pochta                  any   `json:"pochta"`
	ObtainedBonusInfo       []any `json:"obtainedBonusInfo"`
	PossibleBonusInfo       []struct {
		Type    string `json:"type"`
		Amount  int    `json:"amount"`
		Percent int    `json:"percent"`
		Context string `json:"context"`
	} `json:"possibleBonusInfo"`
}
