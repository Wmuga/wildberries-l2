package main

import "github.com/Wmuga/wildberries-l2/develop/dev02/states"

// Реализация с использованием паттерна "состояние"

type fsm struct {
	ctx    *states.SharedCTX
	states map[states.StateNum]states.State
}

func (f *fsm) Unpack(inp string) (string, error) {
	// Инициализация контекста
	f.ctx.Index = 0
	f.ctx.Input = []rune(inp)
	f.ctx.ToWrite = '\000'
	f.ctx.WritesCount = 0
	state := states.StartState

	var err error
	for {
		// Мечемся между состояниями
		state, err = f.states[state].Do()
		if err != nil {
			return "", err
		}
		// Закончена распаковка,
		if state == states.FinishedState {
			break
		}
	}
	// Если есть что дописать - дописывается
	if f.ctx.ToWrite != '\000' {
		err := f.ctx.Write()
		if err != nil {
			return "", err
		}
	}
	return f.ctx.Builder.String(), nil
}

func newFSM() *fsm {
	ctx := &states.SharedCTX{ToWrite: '\000'}
	return &fsm{
		ctx,
		map[states.StateNum]states.State{
			states.StartState:   &states.Start{Ctx: ctx},
			states.CharState:    &states.Char{Ctx: ctx},
			states.DigitState:   &states.Digit{Ctx: ctx},
			states.EscapedState: &states.Escape{Ctx: ctx},
		},
	}
}

func UnpackStringState(input string) (string, error) {
	fsm := newFSM()
	return fsm.Unpack(input)
}
