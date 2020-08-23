package exchangerates

type HistoryResponse struct {
	Rates   map[string]Rates `json:"rates"`
	Base    string           `json:"base"`
	StartAt string           `json:"start_at"`
	EndAt   string           `json:"end_at"`
}

type LatestResponse struct {
	Rates Rates  `json:"rates"`
	Base  string `json:"base"`
	Date  string `json:"date"`
}

type Rates struct {
	CAD float64 `json:"CAD"`
	HKD float64 `json:"HKD"`
	ISK float64 `json:"ISK"`
	PHP float64 `json:"PHP"`
	DKK float64 `json:"DKK"`
	HUF float64 `json:"HUF"`
	CZK float64 `json:"CZK"`
	AUD float64 `json:"AUD"`
	RON float64 `json:"RON"`
	SEK float64 `json:"SEK"`
	IDR float64 `json:"IDR"`
	INR float64 `json:"INR"`
	BRL float64 `json:"BRL"`
	RUB float64 `json:"RUB"`
	HRK float64 `json:"HRK"`
	JPY float64 `json:"JPY"`
	THB float64 `json:"THB"`
	CHF float64 `json:"CHF"`
	SGD float64 `json:"SGD"`
	PLN float64 `json:"PLN"`
	BGN float64 `json:"BGN"`
	TRY float64 `json:"TRY"`
	CNY float64 `json:"CNY"`
	NOK float64 `json:"NOK"`
	NZD float64 `json:"NZD"`
	ZAR float64 `json:"ZAR"`
	USD float64 `json:"USD"`
	MXN float64 `json:"MXN"`
	ILS float64 `json:"ILS"`
	GBP float64 `json:"GBP"`
	KRW float64 `json:"KRW"`
	MYR float64 `json:"MYR"`
}
