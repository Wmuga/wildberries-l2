package flags

import (
	"flag"
	"fmt"
	"os"
)

type Flags struct {
	Timeout int
}

func ParseArgs() (Flags, []string) {
	f := Flags{}
	flag.IntVar(&f.Timeout, "timeout", 10, "время (в секундах) попыток подключения к хосту")
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "go-telnet [OPTION...] address port")
		flag.PrintDefaults()
		os.Exit(0)
	}
	flag.Parse()

	if len(flag.Args()) != 2 {
		flag.Usage()
	}

	if f.Timeout < 0 {
		f.Timeout = 0
	}

	return f, flag.Args()
}
