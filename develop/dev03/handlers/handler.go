package handlers

import (
	"errors"

	"github.com/Wmuga/wildberries-l2/develop/dev03/flags"
)

// Интерфейс для реализации паттерна "цепочка вызовов" или "цепочка обязанностей"
type Handler interface {
	SetNext(Handler)
	invokeNext(flags.Flags, []string) ([]string, error)
	Handle(flags.Flags, []string) ([]string, error)
}

type baseNextHandler struct {
	next Handler
}

func (b *baseNextHandler) SetNext(next Handler) {
	b.next = next
}

func (b *baseNextHandler) invokeNext(f flags.Flags, s []string) ([]string, error) {
	if b.next == nil {
		return nil, ErrNextNotSet
	}
	return b.next.Handle(f, s)
}

var (
	ErrNextNotSet = errors.New("next handler is not set")
)
