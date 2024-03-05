package handlers

import (
	"fmt"
	"github.com/Wmuga/wildberries-l2/develop/dev05/flags"
)

// Добавляет номер строки к строкам
type lineNumHandler struct {
	*baseNextHandler
}

func (l *lineNumHandler) Handle(pattern string, flags flags.Flags, strings []string) ([]HandlerResult, error) {
	res, err := l.invokeNext(pattern, flags, strings)
	if err != nil {
		return nil, err
	}

	if !flags.LineNum {
		return res, nil
	}

	for i := range res {
		c := '-'
		if res[i].Include {
			c = ':'
		}
		res[i].Line = fmt.Sprintf("%d%c%s", i+1, c, res[i].Line)
	}

	return res, nil
}

func NewLineNumHandler() Handler {
	return &lineNumHandler{&baseNextHandler{}}
}
