package states

import (
	"errors"
	"strings"
)

type State interface {
	Do() (StateNum, error)
}

// Разделяемый контекст между состояниями
type SharedCTX struct {
	ToWrite     rune
	WritesCount int
	Index       int
	Input       []rune
	Escape      bool
	Builder     strings.Builder
}

type StateNum int

const (
	StartState    StateNum = iota // смотрим кто пришел
	CharState                     // выводим предыдущее пока встречается число
	DigitState                    // записываем числа пока не символ
	EscapedState                  // используем следующий символ только как записываемое
	FinishedState                 // ничего не делаем
)

var (
	ErrWrongStr = errors.New("wrong input string")
)

func (ctx *SharedCTX) Write() error {
	if ctx.ToWrite == '\000' {
		return ErrWrongStr
	}

	for i := 0; i < ctx.WritesCount; i++ {
		_, err := ctx.Builder.WriteRune(ctx.ToWrite)
		if err != nil {
			return err
		}
	}

	ctx.WritesCount = 0
	ctx.ToWrite = '\000'
	return nil
}
