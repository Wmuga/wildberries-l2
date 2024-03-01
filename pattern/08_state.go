package pattern

import "fmt"

/*
	Реализовать паттерн «состояние».
	Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

/*
	Применимость:
		Наличие объекта, поведение которого меняется в зависимости от состояния
		Код класса содержит много больших схожих операторов, выбирающих поведение в зависимости от значения полей
		Сознательное использования fsm
	Плюсы:
		Убирает множество условных операторов fsm
		Концентрация кода, связанного с состоянием
		Упрощает код контекста
	Минусы:
		Может неоправданно усложнить код при малом количестве состояний
*/

/*
	Предположим задача:
		Необходимо подсчитать количество слов в предложении. Слово - набор символов, разделенными пробелами
	Данный пример будет является таким, для которого паттерн является overkill'ом. Много повторяющихся действий
	Выделяется четыре состояния - начало, слово, пробелы, конец
*/

type State int

const (
	Start State = iota
	Word
	Delim
	End
)

type IState interface {
	Do(rune) State
}

type StartState struct{}

func (s *StartState) Do(r rune) State {
	if r == ' ' {
		return Delim
	}
	return Word
}

type WordState struct {
	wordCounter *int // куда записываем число слов
}

func (s *WordState) Do(r rune) State {
	if r == ' ' {
		(*s.wordCounter)++
		return Delim
	}
	return Word
}

type DelimState struct {
}

func (s *DelimState) Do(r rune) State {
	if r == ' ' {
		return Delim
	}
	return Word
}

type FSM struct {
	wordCount int
	curState  State
	curString []rune
	curIndex  int
	states    map[State]IState
}

func (fsm *FSM) StartWith(str string) {
	fsm.curState = Start
	fsm.wordCount = 0
	fsm.curIndex = 0
	fsm.curString = []rune(str)
}

func (fsm *FSM) Step() bool {
	// Ничего не делаем при окончании
	if fsm.curState == End {
		return false
	}
	// Перешли за границу - закончили
	if fsm.curIndex == len(fsm.curString) {
		fsm.curState = End
		if fsm.curState == Word {
			fsm.wordCount++
		}
		return true
	}
	// Шагаем дальше
	c := fsm.curString[fsm.curIndex]
	fsm.curState = fsm.states[fsm.curState].Do(c)
	fsm.curIndex++
	return true
}

func NewFSM() *FSM {
	fsm := &FSM{}
	fsm.states = map[State]IState{
		Start: &StartState{},
		Word:  &WordState{&fsm.wordCount},
		Delim: &DelimState{},
	}
	return fsm
}

func (fsm *FSM) GetCount() int {
	if fsm.curState != End {
		return 0
	}

	return fsm.wordCount
}

func ExampleStatePattern() {
	words := "a   lot of words      with different    spaces count     "
	fmt.Println("String:", words)

	fsm := NewFSM()
	fsm.StartWith(words)
	for fsm.Step() {
	}
	fmt.Println(fsm.GetCount())
}
