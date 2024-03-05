package handlers

import (
	"github.com/Wmuga/wildberries-l2/develop/dev05/flags"
)

// contextHandler помечает строки, которые должны войти как контекст для результата
type contextHandler struct {
	*baseNextHandler
}

func (c *contextHandler) Handle(pattern string, flags flags.Flags, strings []string) ([]HandlerResult, error) {
	res, err := c.invokeNext(pattern, flags, strings)
	if err != nil {
		return nil, err
	}
	// Добавляем контекст, если другие флаги были 0
	if flags.After == 0 {
		flags.After = flags.Context
	}
	if flags.Before == 0 {
		flags.Before = flags.Context
	}
	// Если котекст не нужен - скип
	if flags.Before == 0 && flags.After == 0 {
		return res, nil
	}
	// Разметка, что строки нужны как контекст
	for i := range res {
		if !res[i].Include {
			continue
		}
		beforeLast := min(max(0, i-flags.Before), len(res)-1)
		for j := i - 1; j >= beforeLast; j-- {
			res[j].IncludeContext = true
		}
		afterLast := min(len(res)-1, i+flags.After)
		for j := i + 1; j <= afterLast; j++ {
			res[j].IncludeContext = true
		}
	}

	return res, nil
}

func NewContextHandler() Handler {
	return &contextHandler{&baseNextHandler{}}
}
