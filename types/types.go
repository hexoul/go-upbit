// Package types of both request and response for API
package types

// Error structure
type Error struct {
	Message string `json:"message"`
	Name    string `json:"name"`
}

// Response structure
type Response struct {
	Err *Error `json:"error"`
}

// Balance structure
type Balance struct {
	Currency       string `json:"currency"`
	Balance        string `json:"balance"`
	Locked         string `json:"locked"`
	AvgKrwBuyPrice string `json:"avg_krw_buy_price"`
	Modified       bool   `json:"modified"`
}

// Order structure
type Order struct {
	UUID            string `json:"uuid"`
	Side            string `json:"side"`
	OrdType         string `json:"ord_type"`
	Price           string `json:"price"`
	AvgPrice        string `json:"avg_price"`
	State           string `json:"state"`
	Market          string `json:"market"`
	CreatedAt       string `json:"created_at"`
	Volume          string `json:"volume"`
	RemainingVolume string `json:"remainingVolume"`
	ReservedFee     string `json:"reserved_fee"`
	RemainingFee    string `json:"remaining_fee"`
	PaidFee         string `json:"paid_fee"`
	Locked          string `json:"locked"`
	ExecutedVolume  string `json:"executed_volume"`
	TradeCount      string `json:"trade_count"`
}
