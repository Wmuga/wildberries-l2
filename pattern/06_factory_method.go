package pattern

import (
	"errors"
	"fmt"
	"log"
)

/*
	Реализовать паттерн «фабричный метод».
	Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

/*
	Применимость:
		Заранее неизвестно объекты каких типов надо создать
		Система должна быть независима от процесса создания новых объектов
		Нуждно делегировать создание новых объектов наследникам
	Плюсы:
		Мзбавляет класс от привязки к конкретным классам
		Выделяет производство в отдельные классы, упрощая поддержку кода
		Упрощает добавление новых продуктов
		Принцип открытости / закрытости
	Минусы:
		Большие параллельные иерархии классов
*/

/*
	Предположим задача:
		Приложение может работать как и в Web, так на Desktop. Необходимо чтобы в зависимости от платформы создавалась своя кнопка
	Класс App будет выбирать в зависимости от конфигурации нужный Creator, который создает конкретный Product(Button)
*/

var (
	ErrUnknownPlatform = errors.New("unknown platform")
)

type IButton interface {
	Click()
}

type WebButton struct{}

func (*WebButton) Click() {
	fmt.Println("Web button")
}

type DesktopButton struct{}

func (*DesktopButton) Click() {
	fmt.Println("Desktop button")
}

type IDialog interface {
	CreateButton() IButton
}

type WebDialog struct{}

func (w *WebDialog) CreateButton() IButton {
	return &WebButton{}
}

type DesktopDialog struct{}

func (w *DesktopDialog) CreateButton() IButton {
	return &DesktopButton{}
}

type App struct {
	dialog IDialog
}

func (a *App) CreateButton() IButton {
	return a.dialog.CreateButton()
}

func NewApp(platform string) (*App, error) {
	switch platform {
	case "desktop":
		return &App{&DesktopDialog{}}, nil
	case "web":
		return &App{&WebDialog{}}, nil
	}
	return nil, ErrUnknownPlatform
}

func ExampleFactoryMethodPattern() {
	app, err := NewApp("desktop")
	if err != nil {
		log.Fatalln(err)
	}
	button := app.CreateButton()
	button.Click()
}
