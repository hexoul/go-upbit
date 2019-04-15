// Package upbit is an API Client for Upbit
package upbit

import (
	"os"
	"strings"
	"sync"
)

// Interface for APIs
type Interface interface {
	// Accounts(options *types.Options) (*types.CryptoInfoMap, error)
}

// Client for CoinMarketCap API
type Client struct {
	accessKey string
	secretKey string
}

var (
	instance  *Client
	once      sync.Once
	accessKey string
	secretKey string
)

const (
	baseURL = "https://api.upbit.com/v1"
)

func init() {
	for _, val := range os.Args {
		arg := strings.Split(val, "=")
		if len(arg) < 2 {
			continue
		} else if arg[0] == "-upbit:accesskey" {
			accessKey = arg[1]
		} else if arg[0] == "-upbit:secretkey" {
			secretKey = arg[1]
		}
	}
}

// GetInstance returns singleton
func GetInstance() *Client {
	once.Do(func() {
		if accessKey == "" || secretKey == "" {
			panic("KEYS FOR BOTH ACCESS AND SECRET REQUIRED")
		}
		instance = &Client{
			accessKey: accessKey,
			secretKey: secretKey,
		}
	})
	return instance
}

// GetInstanceWithKey returns singleton
func GetInstanceWithKey(accessKey, secretKey string) *Client {
	once.Do(func() {
		if accessKey == "" || secretKey == "" {
			panic("KEYS FOR BOTH ACCESS AND SECRET REQUIRED")
		}
		instance = &Client{
			accessKey: accessKey,
			secretKey: secretKey,
		}
	})
	return instance
}

/*
func (s *Client) getResponse(url string) (*types.Response, []byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}
	req.Header.Add("X-CMC_PRO_API_KEY", s.proAPIKey)
	body, err := util.DoReq(req)
	if err != nil {
		return nil, nil, err
	}
	resp := new(types.Response)
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, nil, err
	}
	if resp.Status.ErrorCode != 0 {
		return nil, nil, fmt.Errorf("[%d] %s", resp.Status.ErrorCode, *resp.Status.ErrorMessage)
	}
	return resp, body, nil
}
*/
