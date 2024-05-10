package model

type CNCServiceRequest struct {
	LocationIds []string  `json:"locationIds"`
	Auth        AuthFiled `json:"auth"`
}

type CNCServiceResponse struct {
	Auth      any `json:"auth"`
	Locations []struct {
		Name                  string `json:"name"`
		Address               string `json:"address"`
		WorkingHours          string `json:"workingHours"`
		ConfirmationTimeLimit string `json:"confirmationTimeLimit"`
		LegalInfo             string `json:"legalInfo"`
		PaymentTypes          struct {
			PrepayOnline bool `json:"prepayOnline"`
			Cash         bool `json:"cash"`
			BankCards    bool `json:"bankCards"`
			Spasibo      bool `json:"spasibo"`
		} `json:"paymentTypes"`
		Geo struct {
			Lon string `json:"lon"`
			Lat string `json:"lat"`
		} `json:"geo"`
		ShopDesign struct {
			LogoURL               string `json:"logoUrl"`
			MapLogoURL            string `json:"mapLogoUrl"`
			HeaderBackgroundColor string `json:"headerBackgroundColor"`
			PageBackgroundColor   string `json:"pageBackgroundColor"`
			MerchantColorScheme   string `json:"merchantColorScheme"`
		} `json:"shopDesign"`
		MerchantID string `json:"merchantId"`
		Location   struct {
			Address struct {
				Source     string `json:"Source"`
				Cleared    string `json:"Cleared"`
				PostalCode string `json:"PostalCode"`
				Fias       struct {
					RegionID    string `json:"regionId"`
					Destination string `json:"destination"`
				} `json:"fias"`
				Label any `json:"label"`
			} `json:"address"`
			Geo struct {
				Lon string `json:"lon"`
				Lat string `json:"lat"`
			} `json:"geo"`
			Directions struct {
				Metro           string `json:"metro"`
				TripDescription string `json:"tripDescription"`
			} `json:"directions"`
			ServiceSchemes    []string `json:"serviceSchemes"`
			Caption           any      `json:"caption"`
			MerchantLegalInfo any      `json:"merchantLegalInfo"`
			Schedule          any      `json:"schedule"`
			SLASecMin         int      `json:"slaSecMin"`
			SLASecMax         int      `json:"slaSecMax"`
			SLAText           string   `json:"slaText"`
			GoodsInCartInfo   []any    `json:"goodsInCartInfo"`
		} `json:"location"`
		LocationID       string `json:"locationId"`
		MerchantName     string `json:"merchantName"`
		MerchantFullName string `json:"merchantFullName"`
		ShopIndex        struct {
			Type         string `json:"type"`
			Slug         string `json:"slug"`
			CollectionID string `json:"collectionId"`
		} `json:"shopIndex"`
		ServiceSchemes    []any  `json:"serviceSchemes"`
		Enabled           bool   `json:"enabled"`
		MerchantShortName string `json:"merchantShortName"`
	} `json:"locations"`
}
