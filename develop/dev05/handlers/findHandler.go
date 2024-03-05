package handlers

import (
	"github.com/Wmuga/wildberries-l2/develop/dev05/flags"
	"strings"
)

// findHandler ищет заданную строку в строках файла
type findHandler struct {
	*baseNextHandler
}

func (f findHandler) Handle(pattern string, flags flags.Flags, inp []string) ([]HandlerResult, error) {
	fun := stringsContains
	if flags.IgnoreCase {
		fun = stringsContainsInsensitive
	}

	// скан строк регуляркой
	res := make([]HandlerResult, len(inp))
	for i := range res {
		res[i].Line = inp[i]
		res[i].Include = fun(inp[i], pattern)
	}

	return res, nil
}

func NewFindHandler() Handler {
	return &findHandler{&baseNextHandler{}}
}

// stringsContains функция сравнения строк с учетом регистра
func stringsContains(main, pattern string) bool {
	return strings.Contains(main, pattern)
}

// stringsContainsInsensitive функция сравнения строк без учета регистра
func stringsContainsInsensitive(main, pattern string) bool {
	return strings.Contains(
		strings.ToLower(main),
		strings.ToLower(pattern))

}
