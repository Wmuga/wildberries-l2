package pattern

import (
	"errors"
	"fmt"
)

/*
	Реализовать паттерн «цепочка вызовов».
	Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

/*
	Применимость:
		Обработка запроса несколькими способами, заранее не известно какие запросы придут, какие обработчики понадобятся
		Необходимость выполнения в строгом порядке
		Набор объектов, способных обработать запрос, должен задаваться динамически
	Плюсы:
		Убирает зависимость клиента и обработчика
		Принцип единственной обязанности
		Принцип открытости / закрытости
	Минусы:
		Запрос может остаться необработанным
*/

/*
	Предположим задача:
		Приходит запрос пользователя с логином, паролем.
		Необходимо проверить, что пользователь существует и вернуть ему профиль
	Для решения задачи разработ интерфейс Handler и его реализации AuthHandler, ProfileHandler
	При успешном выполнении AuthHandler будет вызываться ProfileHandler. Иначе - возвращаться ошибка
*/

var (
	ErrNoUser = errors.New("can't find user with auth data")
)

type Request struct {
	Login, Password string
}

type Response struct {
	Username string
}

type Handler interface {
	SetNext(Handler)
	Handle(*Request) (*Response, error)
}

type AuthHandler struct {
	next Handler
}

func (auth *AuthHandler) SetNext(h Handler) {
	auth.next = h
}
func (auth *AuthHandler) Handle(r *Request) (*Response, error) {
	if r.Login == "admin" && r.Password == "password" {
		return auth.next.Handle(r)
	}
	return nil, ErrNoUser
}

// Предполодим, что ProfileHandler в данном случае - конечная точка
type ProfileHandler struct {
}

func (p *ProfileHandler) SetNext(Handler) {}

func (p *ProfileHandler) Handle(r *Request) (*Response, error) {
	return &Response{
		Username: "Admin Админович",
	}, nil
}

func ExampleChainPattern() {
	ph := &ProfileHandler{}
	ah := &AuthHandler{}

	var start Handler = ah

	ah.SetNext(ph)
	r := &Request{"admin", "password"}
	fmt.Println(start.Handle(r))
}
