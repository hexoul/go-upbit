// Package util supports specific parsing
package util

import (
	"io/ioutil"
	"net/http"
	"time"
)

// TimeStamp makes UNIX timestamp
func TimeStamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// DoReq HTTP client
func DoReq(req *http.Request) (body []byte, err error) {
	requestTimeout := time.Duration(5 * time.Second)
	client := &http.Client{
		Timeout: requestTimeout,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if body, err = ioutil.ReadAll(resp.Body); err != nil {
		return nil, err
	}
	return
}
