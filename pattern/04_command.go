package pattern

import "fmt"

/*
	Реализовать паттерн «комманда».
	Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

/*
	Применимость:
		Необходимость параметризации объекта выполняемым действием
		Нужно выполнить операции по очереди/расписанию
		Нужна операция отмены
	Плюсы:
		Убирает зависимость объеакта от выполняемой операции
		Реализация простой отмены / повтора операции
		Реализация отложенного запуска операции
		Сбор сложной команды из простых
		Принцип открытости / закрытости
	Минусы:
		Усложнение кода программы дополнительными классами
*/

/*
	Предположим задача:
		Тестовый редактор. Нужно при нажатии на кнопку "Paste String" добавлять в строку "pasta that". Действие должно быть отменяемым

		Для выполнения вводится интерфейс комманды и ее реализация - PasteCommand.
	Она привязывается к кнопне по интерфейсу, что убирает зависимость объекта кнопки от операции
*/

type Command interface {
	Execute()
	Undo()
}

type PasteCommand struct {
	CurStr     *string // Указатель, т.к. меняет строку за своими пределами
	StrToPaste string
	prevStr    string
}

func (pc *PasteCommand) Execute() {
	pc.prevStr = *pc.CurStr
	*pc.CurStr += pc.StrToPaste
}

func (pc *PasteCommand) Undo() {
	*pc.CurStr = pc.prevStr
}

type Button struct {
	Click Command
}

func ExampleCommandPattern() {
	editorString := "this is editor string || "
	pc := &PasteCommand{
		CurStr:     &editorString,
		StrToPaste: "pasta that",
	}
	pasteButton := Button{pc}
	fmt.Println("Original", editorString)
	pasteButton.Click.Execute()
	fmt.Println("Changed", editorString)
	pasteButton.Click.Undo()
	fmt.Println("Undo", editorString)
}
