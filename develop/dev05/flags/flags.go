package flags

import (
	"flag"
	"fmt"
	"os"
)

// Стурктура флагов из командной строки
type Flags struct {
	After      int
	Before     int
	Context    int
	Count      bool
	IgnoreCase bool
	Invert     bool
	Fixed      bool
	LineNum    bool
}

func ParseArgs() (Flags, []string) {
	f := Flags{}
	flag.IntVar(&f.After, "A", 0, "печатать +N строк после совпадения")
	flag.IntVar(&f.Before, "B", 0, "печатать +N строк до совпадения")
	flag.IntVar(&f.Context, "C", 0, "печатать ±N строк вокруг совпадения")
	flag.BoolVar(&f.Count, "c", false, "количество строк")
	flag.BoolVar(&f.IgnoreCase, "i", false, "игнорировать регистр")
	flag.BoolVar(&f.Invert, "v", false, "вместо совпадения, исключать")
	flag.BoolVar(&f.Fixed, "F", false, "точное совпадение со строкой, не паттерн")
	flag.BoolVar(&f.LineNum, "n", false, "печатать номер строки")
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "grep [OPTION...] PATTERN [FILE...]")
		flag.PrintDefaults()
		os.Exit(0)
	}
	flag.Parse()
	return f, flag.Args()
}
