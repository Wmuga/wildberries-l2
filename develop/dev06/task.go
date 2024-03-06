package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type Flags struct {
	FieldNums []int
	Delimiter string
	Separated bool
}

func GetFlags() Flags {
	var flags Flags
	var fieldsStr string

	flag.StringVar(&fieldsStr, "f", "", "выбрать поля (колонки). Разделение по запятой")
	flag.StringVar(&flags.Delimiter, "d", "\t", "использовать другой разделитель")
	flag.BoolVar(&flags.Separated, "s", false, "только строки с разделителем")
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "cut [OPTION...]")
		flag.PrintDefaults()
		os.Exit(0)
	}
	flag.Parse()

	splitData := strings.Split(fieldsStr, flags.Delimiter)
	flags.FieldNums = make([]int, 0, len(splitData))
	for i := range splitData {
		val, err := strconv.Atoi(splitData[i])
		if err != nil || val < 0 {
			continue
		}
		flags.FieldNums = append(flags.FieldNums, val)
	}

	if len(flags.FieldNums) == 0 {
		flag.Usage()
	}

	return flags
}

func Cut(flags Flags, inp string) []string {
	// Разделение
	strs := strings.Split(inp, flags.Delimiter)
	// Если нет разделитиля и поднят флаг -s - ничего не выводим
	if flags.Separated && len(strs) == 1 {
		return nil
	}

	res := make([]string, 0, len(flags.FieldNums))
	// Копирование нужных полей
	for _, v := range flags.FieldNums {
		if v > len(strs) || v <= 0 {
			continue
		}
		res = append(res, strs[v-1])
	}

	return res
}

func main() {
	flags := GetFlags()
	for {
		var inp string
		// Чтение всего stdin
		n, err := fmt.Scanln(&inp)
		if err != nil && !errors.Is(err, io.EOF) {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		if n <= 0 {
			break
		}

		fmt.Println(strings.Join(Cut(flags, inp), flags.Delimiter))
	}
}
