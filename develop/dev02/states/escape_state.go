package states

type Escape struct {
	Ctx *SharedCTX
}

func (s *Escape) Do() (StateNum, error) {
	// Если конец строки - неверный вход
	if s.Ctx.Index == len(s.Ctx.Input) {
		return FinishedState, ErrWrongStr
	}
	// Текущий символ только на дозапись
	s.Ctx.ToWrite = s.Ctx.Input[s.Ctx.Index]
	s.Ctx.WritesCount = 1
	s.Ctx.Index++

	return CharState, nil
}
