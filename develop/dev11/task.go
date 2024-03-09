package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Wmuga/wildberries-l2/develop/dev11/config"
	"github.com/Wmuga/wildberries-l2/develop/dev11/controllers"
	"github.com/Wmuga/wildberries-l2/develop/dev11/middleware"
	"github.com/Wmuga/wildberries-l2/develop/dev11/repo"
	"github.com/Wmuga/wildberries-l2/develop/dev11/service"
)

/*
=== HTTP server ===

Реализовать HTTP сервер для работы с календарем. В рамках задания необходимо работать строго со стандартной HTTP библиотекой.
В рамках задания необходимо:
	1. Реализовать вспомогательные функции для сериализации объектов доменной области в JSON.
	2. Реализовать вспомогательные функции для парсинга и валидации параметров методов /create_event и /update_event.
	3. Реализовать HTTP обработчики для каждого из методов API, используя вспомогательные функции и объекты доменной области.
	4. Реализовать middleware для логирования запросов
Методы API: POST /create_event POST /update_event POST /delete_event GET /events_for_day GET /events_for_week GET /events_for_month
Параметры передаются в виде www-url-form-encoded (т.е. обычные user_id=3&date=2019-09-09).
В GET методах параметры передаются через queryString, в POST через тело запроса.
В результате каждого запроса должен возвращаться JSON документ содержащий либо {"result": "..."} в случае успешного выполнения метода,
либо {"error": "..."} в случае ошибки бизнес-логики.

В рамках задачи необходимо:
	1. Реализовать все методы.
	2. Бизнес логика НЕ должна зависеть от кода HTTP сервера.
	3. В случае ошибки бизнес-логики сервер должен возвращать HTTP 503. В случае ошибки входных данных (невалидный int например) сервер должен возвращать HTTP 400. В случае остальных ошибок сервер должен возвращать HTTP 500. Web-сервер должен запускаться на порту указанном в конфиге и выводить в лог каждый обработанный запрос.
	4. Код должен проходить проверки go vet и golint.
*/

func main() {
	cfg, err := config.ReadConfig("config.json")
	if err != nil {
		log.Fatalln(err)
	}

	eventRepo := repo.NewMemoryRepo()
	eventService := service.NewEventService(eventRepo)
	controller := controllers.NewEventController(
		eventService,
		log.New(os.Stderr, "[ERR]", log.LUTC))

	logMiddleware := middleware.GetRequestLogger(
		log.New(os.Stdout, "[REQ]", log.LUTC),
	)

	mux := http.NewServeMux()
	// POST методы
	mux.Handle(
		"/create_event",
		middleware.SetMethodFunc("POST", controller.CreateEvent),
	)
	mux.Handle(
		"/update_event",
		middleware.SetMethodFunc("POST", controller.UpdateEvent),
	)
	mux.Handle(
		"/delete_event",
		middleware.SetMethodFunc("POST", controller.DeleteEvent),
	)
	// GET методы
	mux.Handle(
		"/events_for_day",
		middleware.SetMethodFunc("GET", controller.EventsForDay),
	)
	mux.Handle(
		"/events_for_week",
		middleware.SetMethodFunc("GET", controller.EventsForWeek),
	)
	mux.Handle(
		"/events_for_month",
		middleware.SetMethodFunc("GET", controller.EventsForMonth),
	)

	server := http.Server{
		Addr:         cfg.Address,
		Handler:      logMiddleware(mux),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	fmt.Println("Server working on", cfg.Address)
	fmt.Println(server.ListenAndServe())
}
