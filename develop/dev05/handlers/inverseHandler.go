package handlers

import "github.com/Wmuga/wildberries-l2/develop/dev05/flags"

// inverseHandler инвертирует результаты grep
type inverseHandler struct {
	*baseNextHandler
}

func (i *inverseHandler) Handle(pattern string, flags flags.Flags, strings []string) ([]HandlerResult, error) {
	res, err := i.invokeNext(pattern, flags, strings)
	if err != nil {
		return nil, err
	}

	if !flags.Invert {
		return res, nil
	}

	for i := range res {
		res[i].Include = !res[i].Include
	}

	return res, nil
}

func NewInverseHandler() Handler {
	return &inverseHandler{&baseNextHandler{}}
}
