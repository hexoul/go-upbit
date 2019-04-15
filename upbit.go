// Package upbit is an API Client for Upbit
package upbit

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"

	jwt "github.com/dgrijalva/jwt-go"

	"github.com/hexoul/go-upbit/types"
	"github.com/hexoul/go-upbit/util"
)

// Interface for APIs
type Interface interface {
	Accounts() ([]*types.Balance, error)

	GetOrders() ([]*types.Order, error)
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

func (s *Client) getResponse(url, method, sign string, query map[string]string) ([]byte, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	if query != nil {
		q := req.URL.Query()
		for k, v := range query {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}
	req.Header.Add("Authorization", "Bearer "+sign)
	body, err := util.DoReq(req)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(body))
	resp := new(types.Response)
	if err := json.Unmarshal(body, &resp); err == nil && resp.Err != nil {
		return nil, fmt.Errorf("[%s] %s", resp.Err.Name, resp.Err.Message)
	}
	return body, nil
}

func (s *Client) sign(query map[string]string) (string, error) {
	claim := jwt.MapClaims{
		"access_key": s.accessKey,
		"nonce":      util.TimeStamp(),
	}

	if query != nil {
		url := new(url.URL)
		q := url.Query()
		for i, value := range query {
			q.Add(i, value)
		}
		rawQuery := q.Encode()
		claim["query"] = rawQuery
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString([]byte(s.secretKey[:]))
}
