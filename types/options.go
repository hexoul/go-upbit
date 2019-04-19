package types

// Options for request
type Options struct {
	State   string `json:"state,omitempty"`
	Market  string `json:"market,omitempty"`
	Page    string `json:"page,omitempty"`
	OrderBy string `json:"order_by,omitempty"`
}

type stateOptions struct {
	Wait   string
	Done   string
	Cancel string
}

// StateOptions for interval
var StateOptions stateOptions

type orderOptions struct {
	Asec string
	Desc string
}

// OrderOptions for interval
var OrderOptions orderOptions

func init() {
	StateOptions = stateOptions{
		Wait:   "wait",
		Done:   "done",
		Cancel: "cancel",
	}
	OrderOptions = orderOptions{
		Asec: "asec",
		Desc: "desc",
	}
}
