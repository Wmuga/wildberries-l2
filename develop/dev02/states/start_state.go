package states

import "unicode"

type Start struct {
	Ctx *SharedCTX
}

func (s *Start) Do() (StateNum, error) {
	if len(s.Ctx.Input) == 0 {
		return FinishedState, nil
	}
	// Нельзя начинать с числа
	if unicode.IsDigit(s.Ctx.Input[s.Ctx.Index]) {
		return FinishedState, ErrWrongStr
	}

	if s.Ctx.Input[s.Ctx.Index] == '\\' {
		s.Ctx.Index++
		return EscapedState, nil
	}

	s.Ctx.ToWrite = s.Ctx.Input[s.Ctx.Index]
	s.Ctx.WritesCount = 1
	s.Ctx.Index++
	return CharState, nil
}
