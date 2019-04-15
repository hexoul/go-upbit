package upbit

import (
	"encoding/json"

	"github.com/hexoul/go-upbit/types"
)

// Accounts returns asset list
//   src: https://api.upbit.com/v1/accounts
//   doc: https://docs.upbit.com/reference#%EC%9E%90%EC%82%B0-%EC%A0%84%EC%B2%B4-%EC%A1%B0%ED%9A%8C
func (s *Client) Accounts() (balances []*types.Balance, err error) {
	url := baseURL + "/accounts"
	sign, err := s.sign(nil)
	if err != nil {
		return nil, err
	}

	body, err := s.getResponse(url, "GET", sign, nil)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &balances)
	return
}
