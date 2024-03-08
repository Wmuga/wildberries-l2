package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/Wmuga/wildberries-l2/develop/dev10/flags"
	"github.com/Wmuga/wildberries-l2/develop/dev10/telnet"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

// Для теста поднимается сервак, отсылающий все обратно
func SetupLoopback(port string) {
	l, err := net.Listen("tcp", "localhost:"+port)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	con, err := l.Accept()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	reader := bufio.NewReader(con)
	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		// Завершение по exit
		if strings.HasPrefix(str, "exit") {
			con.Close()
			l.Close()
			return
		}

		fmt.Fprint(con, str)
	}
}

func main() {
	flags, args := flags.ParseArgs()
	// go SetupLoopback(args[1])
	err := telnet.Connect(args[0], args[1], flags.Timeout, os.Stdin, os.Stdout)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
