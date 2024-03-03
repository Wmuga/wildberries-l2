package states

import "unicode"

type Digit struct {
	Ctx *SharedCTX
}

func (s *Digit) Do() (StateNum, error) {
	if s.Ctx.Index == len(s.Ctx.Input) {
		return FinishedState, nil
	}
	// Интересует только подсчет значений
	if unicode.IsDigit(s.Ctx.Input[s.Ctx.Index]) {
		s.Ctx.WritesCount = s.Ctx.WritesCount*10 + int(s.Ctx.Input[s.Ctx.Index]-'0')
		s.Ctx.Index++
		return DigitState, nil
	}

	return CharState, nil
}
