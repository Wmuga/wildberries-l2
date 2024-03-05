package main

import (
	"fmt"
	"github.com/Wmuga/wildberries-l2/develop/dev05/flags"
	"github.com/Wmuga/wildberries-l2/develop/dev05/handlers"
	"os"
	"strings"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// GetGrepper устанавливает цепочку count - purger - line_num -  context - inverse - regex_find - find
func GetGrepper() handlers.Handler {
	findHandler := handlers.NewFindHandler()

	regexHandler := handlers.NewRegexHandler()
	regexHandler.SetNext(findHandler)

	inverseHandler := handlers.NewInverseHandler()
	inverseHandler.SetNext(regexHandler)

	contextHandler := handlers.NewContextHandler()
	contextHandler.SetNext(inverseHandler)

	lineNumHandler := handlers.NewLineNumHandler()
	lineNumHandler.SetNext(contextHandler)

	purgerHandler := handlers.NewPurgerHandler()
	purgerHandler.SetNext(lineNumHandler)

	countHandler := handlers.NewCountHandler()
	countHandler.SetNext(purgerHandler)

	return countHandler
}

func Grep(pattern string, flags flags.Flags, lines []string) ([]string, error) {
	grepper := GetGrepper()
	res, err := grepper.Handle(pattern, flags, lines)
	if err != nil {
		return nil, err
	}
	// выковыривыривание строк из результатов
	out := make([]string, len(res))
	for i := range res {
		out[i] = res[i].Line
	}

	return out, nil
}

func main() {
	flags, args := flags.ParseArgs()
	if len(args) < 2 {
		fmt.Fprintln(os.Stderr, "Не введены файлы для чтения и / или паттерн")
		os.Exit(1)
	}
	pattern := args[0]
	for i := 1; i < len(args); i++ {
		bytes, err := os.ReadFile(args[i])
		if err != nil {
			fmt.Fprintln(os.Stderr, "Не введены файлы для чтения и / или паттерн")
			os.Exit(1)
		}
		lines := strings.Split(string(bytes), "\n")

		res, err := Grep(pattern, flags, lines)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Не введены файлы для чтения и / или паттерн")
			os.Exit(1)
		}

		for j := range res {
			fmt.Println(res[j])
		}
	}
}
