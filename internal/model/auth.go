package model

type AuthBody struct {
	Auth AuthFiled `json:"auth"`
}

type AuthFiled struct {
	LocationID  string `json:"locationId"`
	AppPlatform string `json:"appPlatform"`
	AppVersion  int    `json:"appVersion"`
	Experiments struct {
		Num8     string `json:"8"`
		Num55    string `json:"55"`
		Num58    string `json:"58"`
		Num62    string `json:"62"`
		Num68    string `json:"68"`
		Num69    string `json:"69"`
		Num79    string `json:"79"`
		Num84    string `json:"84"`
		Num96    string `json:"96"`
		Num98    string `json:"98"`
		Num99    string `json:"99"`
		Num107   string `json:"107"`
		Num109   string `json:"109"`
		Num119   string `json:"119"`
		Num120   string `json:"120"`
		Num121   string `json:"121"`
		Num122   string `json:"122"`
		Num128   string `json:"128"`
		Num130   string `json:"130"`
		Num132   string `json:"132"`
		Num144   string `json:"144"`
		Num147   string `json:"147"`
		Num154   string `json:"154"`
		Num163   string `json:"163"`
		Num173   string `json:"173"`
		Num178   string `json:"178"`
		Num184   string `json:"184"`
		Num186   string `json:"186"`
		Num190   string `json:"190"`
		Num192   string `json:"192"`
		Num194   string `json:"194"`
		Num200   string `json:"200"`
		Num205   string `json:"205"`
		Num209   string `json:"209"`
		Num213   string `json:"213"`
		Num218   string `json:"218"`
		Num228   string `json:"228"`
		Num229   string `json:"229"`
		Num235   string `json:"235"`
		Num237   string `json:"237"`
		Num238   string `json:"238"`
		Num243   string `json:"243"`
		Num255   string `json:"255"`
		Num644   string `json:"644"`
		Num645   string `json:"645"`
		Num723   string `json:"723"`
		Num5779  string `json:"5779"`
		Num20121 string `json:"20121"`
		Num43568 string `json:"43568"`
		Num70070 string `json:"70070"`
		Num80283 string `json:"80283"`
		Num85160 string `json:"85160"`
	} `json:"experiments"`
	Os string `json:"os"`
}
