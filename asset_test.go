// NOTE: before testing, apply your Upbit API keys to init() of `upbit_test.go`
package upbit

import (
	"testing"
)

func TestAccounts(t *testing.T) {
	if ret, err := GetInstance().Accounts(); err != nil {
		t.Fatal(err)
	} else {
		t.Logf("%v", ret)
	}
}
