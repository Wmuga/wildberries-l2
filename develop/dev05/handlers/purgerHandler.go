package handlers

import (
	"github.com/Wmuga/wildberries-l2/develop/dev05/flags"
	"slices"
)

// purgerHandler очищает невошедшие в результат строки
type purgerHandler struct {
	*baseNextHandler
}

func (p *purgerHandler) Handle(pattern string, flags flags.Flags, strings []string) ([]HandlerResult, error) {
	res, err := p.invokeNext(pattern, flags, strings)
	if err != nil {
		return nil, err
	}

	res = slices.DeleteFunc(res, func(result HandlerResult) bool { return !result.Include && !result.IncludeContext })

	return res, nil
}

func NewPurgerHandler() Handler {
	return &purgerHandler{&baseNextHandler{}}
}
