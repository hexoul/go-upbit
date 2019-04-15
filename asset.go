package upbit

import (
	"encoding/json"

	"github.com/hexoul/go-upbit/types"
)

// Accounts returns asset list
func (s *Client) Accounts() (balances []*types.Balance, err error) {
	url := baseURL + "/accounts"
	sign, err := s.sign(nil)
	if err != nil {
		return nil, err
	}

	body, err := s.getResponse(url, "GET", sign)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &balances)
	return
}
