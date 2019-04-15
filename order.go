package upbit

import (
	"encoding/json"

	"github.com/hexoul/go-upbit/types"
	"github.com/hexoul/go-upbit/util"
)

// GetOrders returns order list
//   arg: market, state, page, order_by
//   src: https://api.upbit.com/v1/orders
//   doc: https://docs.upbit.com/reference#%EC%A3%BC%EB%AC%B8-%EB%A6%AC%EC%8A%A4%ED%8A%B8-%EC%A1%B0%ED%9A%8C
func (s *Client) GetOrders(options *types.Options) (orders []*types.Order, err error) {
	url := baseURL + "/orders"
	query := util.ParseOptions(options)
	sign, err := s.sign(query)
	if err != nil {
		return nil, err
	}

	body, err := s.getResponse(url, "GET", sign, query)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &orders)
	return
}
