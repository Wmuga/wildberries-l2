package flags

import (
	"flag"
	"fmt"
	"os"
)

type Flags struct {
	Wait    int
	Recurse int
	Prefix  string
}

func ParseArgs() (Flags, []string) {
	f := Flags{}
	flag.StringVar(&f.Prefix, "o", "./output", "директория куда сохранять файлы")
	flag.IntVar(&f.Wait, "t", 1, "сколько ждать (в секундах) между загрузкой файлов")
	flag.IntVar(&f.Recurse, "r", 1, "насколько глубоко скачивать сайт")
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "go-wget [OPTION...] url")
		flag.PrintDefaults()
		os.Exit(0)
	}
	flag.Parse()
	return f, flag.Args()
}
