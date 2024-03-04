package handlers

import (
	"slices"

	"github.com/Wmuga/wildberries-l2/develop/dev03/flags"
)

type reverse struct {
	*baseNextHandler
}

// Handle implements Handler.
func (r *reverse) Handle(flags flags.Flags, input []string) ([]string, error) {
	resNext, err := r.invokeNext(flags, input)
	if err != nil {
		return nil, err
	}

	if flags.Reverse {
		slices.Reverse(resNext)
	}

	return resNext, nil
}

func NewReverseHandler() Handler {
	return &reverse{&baseNextHandler{}}
}
