package upbit

import (
	"testing"

	"github.com/hexoul/go-upbit/types"
)

func TestGetOrders(t *testing.T) {
	if ret, err := GetInstance().GetOrders(&types.Options{
		State:   types.StateOptions.Done,
		OrderBy: "desc",
	}); err != nil {
		t.Fatal(err)
	} else {
		for _, v := range ret {
			t.Log(v.UUID, v.Market, v.Side, v.CreatedAt, v.Price, v.ExecutedVolume, v.PaidFee)
		}
	}
}
