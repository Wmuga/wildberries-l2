package pattern

import "strings"

/*
	Реализовать паттерн «строитель».
	Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

/*
	Применимость:
		Избавление от "телскопического конструктора", т.е. в котором все и вся
		Нужно создавать разные представления объекта
	Плюсы:
		Пошаговое создание объектов
		Один код на разные продукты
		Изоляция сложной сборки от бизнес-логики
	Минусы:
		Усложнение кода из-за доп классов
		Привязка к конкретным строителям
*/

/*
	Предположим задачу: Сделать собиратель бургеров.

	Директор предлагает чизбургер, обыкновенный бургер
		Билдер Добавляет каждый ингридиент
			Класс бургер в конструкторе имеет ингридиенты
*/

// Сам страшный класс бургер с непонятным конструктором
type Burger struct {
	patty    bool
	ketchup  bool
	cheese   bool
	bun      bool
	addPatty bool
	mayo     bool
}

func NewBurger(patty, ketchup, cheese, bun, addPatty, mayo bool) *Burger {
	return &Burger{patty, ketchup, cheese, bun, addPatty, mayo}
}

func (b *Burger) Consume() (string, error) {
	builder := strings.Builder{}
	if _, err := builder.WriteString("You ate a burger with: "); err != nil {
		return "", err
	}
	insides := []bool{b.patty, b.ketchup, b.cheese, b.bun, b.addPatty, b.mayo}
	insidesStr := []string{"patty ", "ketchup ", "cheese ", "bun ", "additional patty ", "mayo "}
	for i := range insides {
		if !insides[i] {
			continue
		}
		if _, err := builder.WriteString(insidesStr[i]); err != nil {
			return "", err
		}
	}
	return builder.String(), nil
}

// Строитель бургера
type BurgerBuilder struct {
	patty, ketchup, cheese, bun, addPatty, mayo bool
}

func (bb *BurgerBuilder) AddPatty() {
	bb.patty = true
}

func (bb *BurgerBuilder) AddKetchup() {
	bb.ketchup = true
}

func (bb *BurgerBuilder) AddCheese() {
	bb.cheese = true
}

func (bb *BurgerBuilder) AddBun() {
	bb.bun = true
}

func (bb *BurgerBuilder) AddSecondPatty() {
	bb.addPatty = true
}

func (bb *BurgerBuilder) AddMayo() {
	bb.mayo = true
}

func (bb *BurgerBuilder) Build() *Burger {
	burger := NewBurger(bb.patty, bb.ketchup, bb.cheese, bb.bun, bb.addPatty, bb.mayo)
	bb.patty = false
	bb.ketchup = false
	bb.cheese = false
	bb.bun = false
	bb.addPatty = false
	bb.mayo = false
	return burger
}

type BurgerDirector struct {
	builder *BurgerBuilder
}

func NewBurgerDirector() *BurgerDirector {
	return &BurgerDirector{&BurgerBuilder{}}
}

func (bd *BurgerDirector) GetHamBurger() *Burger {
	bd.builder.AddPatty()
	bd.builder.AddKetchup()
	bd.builder.AddMayo()
	return bd.builder.Build()
}
func (bd *BurgerDirector) GetCheeseBurger() *Burger {
	bd.builder.AddPatty()
	bd.builder.AddKetchup()
	bd.builder.AddCheese()
	bd.builder.AddBun()
	bd.builder.AddSecondPatty()
	return bd.builder.Build()
}

func ExampleBuilderPattern() {
	// Получается для получения нужного "подтипа" необходимо только попросить
	director := NewBurgerDirector()
	_ = director.GetHamBurger()
	_ = director.GetCheeseBurger()
}
