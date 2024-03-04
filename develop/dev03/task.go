package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Wmuga/wildberries-l2/develop/dev03/flags"
	"github.com/Wmuga/wildberries-l2/develop/dev03/handlers"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// Паттерн цепочка вызовов или цепочка обязанностей
func GetSorter() handlers.Handler {
	sorter := handlers.NewSorterHandler()

	uniq := handlers.NewUniquesHandler()
	uniq.SetNext(sorter)

	rev := handlers.NewReverseHandler()
	rev.SetNext(uniq)
	return rev
}

func Sort(flags flags.Flags, input []string) ([]string, error) {
	sorter := GetSorter()
	return sorter.Handle(flags, input)
}

func main() {
	// Получить входные данные
	flags, args := flags.ParseArgs()
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "Не введены файлы для чтения")
		os.Exit(1)
	}
	// По каждому файлу
	for i := range args {
		// Пытаемся прочитать
		bytes, err := os.ReadFile(args[i])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		// Разбиваем по строкам
		fileStrs := strings.Split(string(bytes), "\n")
		fileStrs, err = Sort(flags, fileStrs)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		// Выводим
		for j := range fileStrs {
			fmt.Println(fileStrs[j])
		}
	}
}
