package flags

import (
	"flag"
	"fmt"
	"os"
)

// Стурктура флагов из командной строки
type Flags struct {
	Colon   int
	ByInt   bool
	Reverse bool
	Unique  bool
}

func ParseArgs() (Flags, []string) {
	f := Flags{}
	flag.IntVar(&f.Colon, "k", 0, "Индекс колонки для сортировки")
	flag.BoolVar(&f.ByInt, "n", false, "Сортировать по числовому значению")
	flag.BoolVar(&f.Reverse, "r", false, "Сортировать в обратном порядке")
	flag.BoolVar(&f.Unique, "u", false, "Вывести уникальные строки")
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "sort [options] file1 file2 ...")
		flag.PrintDefaults()
		os.Exit(0)
	}
	flag.Parse()
	return f, flag.Args()
}
