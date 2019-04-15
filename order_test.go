package upbit

import (
	"testing"

	"github.com/hexoul/go-upbit/types"
)

func TestGetOrders(t *testing.T) {
	if ret, err := GetInstance().GetOrders(&types.Options{
		State: types.StateOptions.Done,
	}); err != nil {
		t.Fatal(err)
	} else {
		for _, v := range ret {
			t.Log(v.UUID, v.Market, v.State, v.ExecutedVolume)
		}
	}
}
