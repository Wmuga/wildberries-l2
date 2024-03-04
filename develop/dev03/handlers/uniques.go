package handlers

import "github.com/Wmuga/wildberries-l2/develop/dev03/flags"

type uniques struct {
	*baseNextHandler
}

// Handle implements Handler.
func (u *uniques) Handle(flags flags.Flags, input []string) ([]string, error) {
	resNext, err := u.invokeNext(flags, input)
	if err != nil {
		return nil, err
	}

	if !flags.Unique {
		return resNext, nil
	}

	// копируем уникальные строки
	res := make([]string, 0)
	for i := 0; i < len(resNext); i++ {
		if i != 0 && resNext[i] == resNext[i-1] {
			continue
		}
		res = append(res, resNext[i])
	}
	return res, nil
}

func NewUniquesHandler() Handler {
	return &uniques{&baseNextHandler{}}
}
