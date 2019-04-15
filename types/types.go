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
