// Package types of both request and response for API
package types

// Status structure
type Status struct {
	Timestamp    string  `json:"timestamp"`
	ErrorCode    int     `json:"error_code"`
	ErrorMessage *string `json:"error_message"`
	Elapsed      int     `json:"elapsed"`
	CreditCount  int     `json:"credit_count"`
}
