package handlers

import (
	"github.com/Wmuga/wildberries-l2/develop/dev05/flags"
	"strconv"
)

// countHandler подсчитывает все строки, подходящие по запросу
type countHandler struct {
	*baseNextHandler
}

func (c *countHandler) Handle(pattern string, flags flags.Flags, strings []string) ([]HandlerResult, error) {
	res, err := c.invokeNext(pattern, flags, strings)
	if err != nil {
		return nil, err
	}
	// Проверка, что модуль нужен
	if !flags.Count {
		return res, err
	}

	var count int64
	for i := range res {
		if res[i].Include {
			count++
		}
	}
	return []HandlerResult{{Line: strconv.FormatInt(count, 10), Include: true}}, nil
}

func NewCountHandler() Handler {
	return &countHandler{&baseNextHandler{}}
}
