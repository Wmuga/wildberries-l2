package states

import "unicode"

type Char struct {
	Ctx *SharedCTX
}

func (s *Char) Do() (StateNum, error) {
	if s.Ctx.Index == len(s.Ctx.Input) {
		return FinishedState, nil
	}
	// Если число - будет пересчитываться количество записей
	if unicode.IsDigit(s.Ctx.Input[s.Ctx.Index]) {
		s.Ctx.WritesCount = 0
		return DigitState, nil
	}
	// Пробуем записать, что есть
	err := s.Ctx.Write()
	if err != nil {
		return FinishedState, err
	}
	// Если сейчас \ - идем к Escape состоянию
	if s.Ctx.Input[s.Ctx.Index] == '\\' {
		s.Ctx.Index++
		return EscapedState, nil
	}
	// Иначе записываем что есть
	s.Ctx.ToWrite = s.Ctx.Input[s.Ctx.Index]
	s.Ctx.WritesCount = 1
	s.Ctx.Index++
	return CharState, nil
}
