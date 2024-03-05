package handlers

import (
	"errors"

	"github.com/Wmuga/wildberries-l2/develop/dev05/flags"
)

type HandlerResult struct {
	Line           string
	Include        bool
	IncludeContext bool
}

// Handler - Интерфейс для реализации паттерна "цепочка вызовов" или "цепочка обязанностей"
type Handler interface {
	SetNext(Handler)
	invokeNext(string, flags.Flags, []string) ([]HandlerResult, error)
	Handle(string, flags.Flags, []string) ([]HandlerResult, error)
}

type baseNextHandler struct {
	next Handler
}

func (b *baseNextHandler) SetNext(next Handler) {
	b.next = next
}

func (b *baseNextHandler) invokeNext(pattern string, f flags.Flags, s []string) ([]HandlerResult, error) {
	if b.next == nil {
		return nil, ErrNextNotSet
	}
	return b.next.Handle(pattern, f, s)
}

var (
	ErrNextNotSet = errors.New("next handler is not set")
)
