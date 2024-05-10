package model

type GetProfileResponse struct {
	Profile   CustomerProfile `json:"profile"`
	Addresses []struct {
		KladrID      string `json:"kladrId"`
		ID           string `json:"id"`
		Full         string `json:"full"`
		Apartment    string `json:"apartment"`
		Building     string `json:"building"`
		BuildingType string `json:"buildingType"`
		City         string `json:"city"`
		CityType     string `json:"cityType"`
		Entrance     string `json:"entrance"`
		Flat         string `json:"flat"`
		FlatType     string `json:"flatType"`
		Floor        string `json:"floor"`
		House        string `json:"house"`
		HouseType    string `json:"houseType"`
		Intercom     string `json:"intercom"`
		PostalCode   string `json:"postalCode"`
		Region       string `json:"region"`
		RegionType   string `json:"regionType"`
		Street       string `json:"street"`
		StreetType   string `json:"streetType"`
		Name         string `json:"name"`
		Type         string `json:"type"`
		Geo          struct {
			Lon float64 `json:"lon"`
			Lat float64 `json:"lat"`
		} `json:"geo"`
		IsDefault bool   `json:"isDefault"`
		RegionID  string `json:"regionId"`
	} `json:"addresses"`
	Privileges           []any    `json:"privileges"`
	Groups               []string `json:"groups"`
	LoyaltyEmergencyMode bool     `json:"loyaltyEmergencyMode"`
}

type Main struct {
	Profile              Profile       `json:"profile"`
	Addresses            []interface{} `json:"addresses"`
	Privileges           []interface{} `json:"privileges"`
	Groups               []string      `json:"groups"`
	LoyaltyEmergencyMode bool          `json:"loyaltyEmergencyMode"`
}

type Profile struct {
	ID                     string           `json:"id"`
	Email                  string           `json:"email"`
	Phone                  string           `json:"phone"`
	FirstName              string           `json:"firstName"`
	LastName               string           `json:"lastName"`
	SecondName             string           `json:"secondName"`
	BirthDate              string           `json:"birthDate"`
	CheckReceipt           string           `json:"checkReceipt"`
	Gender                 string           `json:"gender"`
	BonusBalance           int64            `json:"bonusBalance"`
	AvatarLink             string           `json:"avatarLink"`
	Misc                   Misc             `json:"misc"`
	IsAuthenticated        bool             `json:"isAuthenticated"`
	IsBlackList            bool             `json:"isBlackList"`
	IsRequiredFieldsFilled bool             `json:"isRequiredFieldsFilled"`
	AgreedToTermsOfUseAt   string           `json:"agreedToTermsOfUseAt"`
	LastPaymentWays        []LastPaymentWay `json:"lastPaymentWays"`
}

type LastPaymentWay struct {
	PaymentType string      `json:"paymentType"`
	Params      interface{} `json:"params"`
}

type Misc struct {
	ReviewsCount      int64  `json:"reviewsCount"`
	IosAppVersion     string `json:"iosAppVersion"`
	AndroidAppVersion string `json:"androidAppVersion"`
}
