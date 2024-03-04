package handlers

import (
	"math"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/Wmuga/wildberries-l2/develop/dev03/flags"
)

type sorter struct {
	*baseNextHandler
}

// Handle implements Handler.
func (s *sorter) Handle(flags flags.Flags, inputs []string) ([]string, error) {
	splitWords := make([][]string, len(inputs))
	for i := range inputs {
		splitWords[i] = strings.Split(inputs[i], " ")
		// Убрать пустые строки
		splitWords[i] = slices.DeleteFunc(splitWords[i], func(s string) bool { return s == "" })
	}
	// Убрать пустые слайсы
	splitWords = slices.DeleteFunc(splitWords, func(s []string) bool { return len(s) == 0 })

	sorter := getCompByStr(splitWords, flags.Colon)
	if flags.ByInt {
		sorter = getCompByInt(splitWords, flags.Colon)
	}

	sort.SliceStable(splitWords, sorter)

	out := make([]string, len(splitWords))
	for i := range out {
		out[i] = strings.Join(splitWords[i], " ")
	}

	return out, nil
}

// SetNext implements Handler.
// Subtle: this method shadows the method (*baseNextHandler).SetNext of sorter.baseNextHandler.
// Этот Handler должен быть последним. Поэтому установка следующего ничего не делает
func (*sorter) SetNext(Handler) {}

func NewSorterHandler() Handler {
	return &sorter{&baseNextHandler{}}
}

// Конвертирует строку в число. Если не удается - возвращает -1
func convNoError(inp string) int64 {
	data, err := strconv.ParseInt(inp, 10, 64)
	if err != nil {
		return math.MaxInt64
	}
	return data
}

// Функция сортировки по колонке пытаясь преобразовать строку в число
func getCompByInt(inp [][]string, colon int) func(i, j int) bool {
	return func(i, j int) bool {
		iColon := min(colon, len(inp[i])-1)
		jColon := min(colon, len(inp[j])-1)

		iNum := convNoError(inp[i][iColon])
		jNum := convNoError(inp[j][jColon])

		if iColon == -1 {
			return false
		}

		if jColon == -1 {
			return true
		}

		return iNum < jNum
	}
}

// Функция сортировки по колонке
func getCompByStr(inp [][]string, colon int) func(i, j int) bool {
	return func(i, j int) bool {
		iColon := min(colon, len(inp[i])-1)
		jColon := min(colon, len(inp[j])-1)
		if iColon == -1 {
			return false
		}

		if jColon == -1 {
			return true
		}
		return inp[i][iColon] < inp[j][jColon]
	}
}
