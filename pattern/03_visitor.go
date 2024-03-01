package pattern

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

/*
	Реализовать паттерн «посетитель».
	Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

/*
	Применимость:
		Необходимость выполнения операции над всеми элементами сложной структуры
		Нужно выполнить не связанный с классом опреции, но нет желания засорять ими сами классы
		Новое поведение имеет смысл для некоторых из них
	Плюсы:
		Упрощает добавление операций, работающих со сложной структурой объектов
		Объединение родственных операций в один класс
		Можно накапливать состояние
	Минусы:
		Невозможно использовать, если иерархия часто меняется
		Может быть нарушение инкапсуляции
*/

/*
	Предположим задача:
		Есть несколько разных фигур типов Circle, Rectangle и необходимо сериализовать их в один JSON
	Для выполнения вводится интерфейс Visitor и реализует его JsonVisitor
	Каждой структуре фигур добавлен метод Accept, вызывающий соответсвующую ему функцию в Vistor'е (double dispatch)
	Без доступа к классам пришлось бы самим определять тип Shape'а и вызывать нужную функцию
*/

// Сами фигуры
type Shape interface {
	Accept(v Visitor) error
}

type Circle struct {
	X, Y, R int
}

type Rectangle struct {
	X, Y, Widhth, Height int
}

func (c *Circle) Accept(v Visitor) error {
	return v.ForCircle(c)
}

func (r *Rectangle) Accept(v Visitor) error {
	return v.ForRectangle(r)
}

// Визитор
type Visitor interface {
	ForCircle(*Circle) error
	ForRectangle(*Rectangle) error
	AddShape(Shape)
	GetResult() ([]byte, error)
}

// Его реализация
type JsonVisitor struct {
	shapes []Shape
	res    *bytes.Buffer
}

func (j *JsonVisitor) AddShape(s Shape) {
	j.shapes = append(j.shapes, s)
}

func (j *JsonVisitor) ForCircle(c *Circle) error {
	_, err := fmt.Fprintf(j.res, `{"type":"circle","x":%d,"y":%d,"r":%d}`, c.X, c.Y, c.R)
	return err
}

func (j *JsonVisitor) ForRectangle(r *Rectangle) error {
	_, err := fmt.Fprintf(j.res, `{"type":"rectangle","x":%d,"y":%d,"width":%d,"height":%d}`, r.X, r.Y, r.Widhth, r.Height)
	return err
}

func (j *JsonVisitor) GetResult() ([]byte, error) {
	if _, err := j.res.WriteRune('['); err != nil {
		return nil, err
	}

	for i := len(j.shapes) - 1; i >= 0; i-- {
		if err := j.shapes[i].Accept(j); err != nil {
			return nil, err
		}
		if i != 0 {
			if _, err := j.res.WriteRune(','); err != nil {
				return nil, err
			}
		}
	}

	if _, err := j.res.WriteRune(']'); err != nil {
		return nil, err
	}
	return j.res.Bytes(), nil
}

func NewJsonVisitor() Visitor {
	return &JsonVisitor{
		make([]Shape, 0),
		&bytes.Buffer{},
	}
}

func ExampleVisitorPattern() {
	visitor := NewJsonVisitor()
	visitor.AddShape(&Circle{1, 2, 3})
	visitor.AddShape(&Rectangle{1, 2, 3, 4})
	visitor.AddShape(&Circle{3, 4, 6})
	bytes, err := visitor.GetResult()
	if err != nil {
		log.Fatalln(err)
	}
	file, err := os.Create("visitor.json")
	if err != nil {
		log.Fatalln(err)
	}
	if _, err = file.Write(bytes); err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Done")
}
