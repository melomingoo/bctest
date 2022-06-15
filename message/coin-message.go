package message

import "time"

type CoinList struct {
	Status Status `json:"status"`
	Data   []Data `json:"data"`
}

type CoinListRequest struct {
	Currency       string   `json:"currency"`
	TargetCurrency []string `json:"target_currency""`
}

type CoinCurrentListRequest struct {
	Currency string `json:"currency"`
	Target   string `json:"target""`
}

type CoinCurrentListResponse struct {
	CoinCurrentData []CoinCurrentData `json:"coin_current_data"`
}

type CoinCurrentData struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price"`
	Time   string  `json:"time""`
}

type CoinResponse struct {
	LastUpdate string   `json:"last_update"`
	BTC        Currency `json:"BTC"`
	ETH        Currency `json:"ETH"`
}

type Status struct {
	Timestamp    string `json:"timestamp"`
	ErrorCode    int32  `json:"error_code"`
	ErrorMessage string `json:"error_message"`
	Elapsed      int32  `json:"elapsed"`
	CreditCount  int32  `json:"credit_count"`
	Notice       string `json:"notice"`
	TotalCount   int32  `json:"total_count"`
}

type Data struct {
	Id                            int32    `json:"id"`
	Name                          string   `json:"name"`
	Symbol                        string   `json:"symbol"`
	Slug                          string   `json:"slug"`
	NumMarketPairs                float64  `json:"num_market_pairs"`
	DateAdded                     string   `json:"date_added"`
	Tags                          []string `json:"tags"`
	MaxSupply                     float64  `json:"max_supply"`
	CirculatingSupply             float64  `json:"circulating_supply"`
	TotalSupply                   float64  `json:"total_supply"`
	Platform                      Platform `json:"platform"`
	CmcRank                       float64  `json:"cmc_rank"`
	SelfReportedCirculatingSupply float64  `json:"self_reported_circulating_supply"`
	SelfReportedMarketCap         float64  `json:"self_reported_market_cap"`
	LastUpdated                   string   `json:"last_updated"`
	Quote                         Quote    `json:"quote"`
}

type Platform struct {
	Id           int64  `json:"id"`
	Name         string `json:"name"`
	Symbol       string `json:"symbol"`
	Slug         string `json:"slug"`
	TokenAddress string `json:"TokenAddress"`
}

type Quote struct {
	USD  Currency `json:"USD"`
	KRW  Currency `json:"KRW"`
	ATOM ATOM     `json:"ATOM"`
	XRP  ATOM     `json:"XRP"`
}

type Currency struct {
	Price                 float64 `json:"price"`
	Volume24h             float64 `json:"volume_24h"`
	VolumeChange24h       float64 `json:"volume_change_24h"`
	PercentChange1h       float64 `json:"percent_change_1h"`
	PercentChange24h      float64 `json:"percent_change_24h"`
	PercentChange7d       float64 `json:"percent_change_7d"`
	PercentChange30d      float64 `json:"percent_change_30d"`
	PercentChange60d      float64 `json:"percent_change_60d"`
	PercentChange90d      float64 `json:"percent_change_90d"`
	MarketCap             float64 `json:"market_cap"`
	MarketCapDominance    float64 `json:"market_cap_dominance"`
	FullyDilutedMarketCap float64 `json:"fully_diluted_market_cap"`
	LastUpdated           string  `json:"last_updated"`
}

type CoinHistoryRequest struct {
	Currency  string `json:"currency"`
	Target    string `json:"target"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

type CoinHistoryResponse struct {
	Id                    int32   `json:"id"`
	Name                  string  `json:"name"`
	Symbol                string  `json:"symbol"`
	Price                 float64 `json:"price"`
	Volume24h             float64 `json:"volume_24h"`
	VolumeChange24h       float64 `json:"volume_change_24h"`
	PercentChange1h       float64 `json:"percent_change_1h"`
	PercentChange24h      float64 `json:"percent_change_24h"`
	PercentChange7d       float64 `json:"percent_change_7d"`
	PercentChange30d      float64 `json:"percent_change_30d"`
	PercentChange60d      float64 `json:"percent_change_60d"`
	PercentChange90d      float64 `json:"percent_change_90d"`
	MarketCap             float64 `json:"market_cap"`
	MarketCapDominance    float64 `json:"market_cap_dominance"`
	FullyDilutedMarketCap float64 `json:"fully_diluted_market_cap"`
	LastUpdated           string  `json:"last_updated"`
}
type CoinHistoryListResponse struct {
	Items []*CoinHistoryResponse
}

type CoinChangeRequest struct {
	From   string `json:"from"`
	To     string `json:"to"`
	Amount string `json:"amount"`
}

type CoinChangeResponse struct {
	Price  float64 `json:"price"`
	Price2 float64 `json:"price2"`
	Quote  Quote   `json:"quote"`
}

type ChangeDataResponse struct {
	Status ChangeStatus `json:"status"`
	Data   []ChangeData `json:"data"`
}

type ChangeData struct {
	ID          int32   `json:"id"`
	Symbol      string  `json:"symbol"`
	Name        string  `json:"name"`
	Amount      float64 `json:"amount"`
	LastUpdated string  `json:"last_update"`
	Quote       Quote   `json:"quote"`
}

type ATOM struct {
	Price       float64 `json:"price"`
	LastUpdated string  `json:"last_update"`
}

type ChangeStatus struct {
	Timestamp    string `json:"timestamp"`
	ErrorCode    int32  `json:"error_code"`
	ErrorMessage string `json:"error_message"`
	Elapsed      int32  `json:"elapsed"`
	CreditCount  int32  `json:"credit_count"`
	Notice       string `json:"notice"`
}

type AutoGenerated struct {
	Status Status2 `json:"status"`
	Data   Data2   `json:"data"`
}
type Status2 struct {
	Timestamp    time.Time   `json:"timestamp"`
	ErrorCode    int         `json:"error_code"`
	ErrorMessage interface{} `json:"error_message"`
	Elapsed      int         `json:"elapsed"`
	CreditCount  int         `json:"credit_count"`
	Notice       interface{} `json:"notice"`
}
type Tags struct {
	Slug     string `json:"slug"`
	Name     string `json:"name"`
	Category string `json:"category"`
}
type CURE struct {
	Price                 float64   `json:"price"`
	Volume24H             float64   `json:"volume_24h"`
	VolumeChange24H       float64   `json:"volume_change_24h"`
	PercentChange1H       float64   `json:"percent_change_1h"`
	PercentChange24H      float64   `json:"percent_change_24h"`
	PercentChange7D       float64   `json:"percent_change_7d"`
	PercentChange30D      float64   `json:"percent_change_30d"`
	PercentChange60D      float64   `json:"percent_change_60d"`
	PercentChange90D      float64   `json:"percent_change_90d"`
	MarketCap             float64   `json:"market_cap"`
	MarketCapDominance    float64   `json:"market_cap_dominance"`
	FullyDilutedMarketCap float64   `json:"fully_diluted_market_cap"`
	LastUpdated           time.Time `json:"last_updated"`
}
type Quote2 struct {
	USD CURE `json:""`
	KRW CURE `json:""`
}
type Cripto struct {
	ID                            int         `json:"id"`
	Name                          string      `json:"name"`
	Symbol                        string      `json:"symbol"`
	Slug                          string      `json:"slug"`
	NumMarketPairs                int         `json:"num_market_pairs"`
	DateAdded                     time.Time   `json:"date_added"`
	Tags                          []Tags      `json:"tags"`
	MaxSupply                     int         `json:"max_supply"`
	CirculatingSupply             int         `json:"circulating_supply"`
	TotalSupply                   int         `json:"total_supply"`
	IsActive                      int         `json:"is_active"`
	Platform                      interface{} `json:"platform"`
	CmcRank                       int         `json:"cmc_rank"`
	IsFiat                        int         `json:"is_fiat"`
	SelfReportedCirculatingSupply interface{} `json:"self_reported_circulating_supply"`
	SelfReportedMarketCap         interface{} `json:"self_reported_market_cap"`
	LastUpdated                   time.Time   `json:"last_updated"`
	Quote                         Quote2      `json:"quote"`
}
type Data2 struct {
	ETH  []Cripto `json:""`
	BTC  []Cripto `json:""`
	USDT []Cripto `json:""`
}
