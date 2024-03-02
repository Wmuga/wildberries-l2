package main

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

/*
	=== Базовая задача ===

	Создать программу печатающую точное время с использованием NTP библиотеки.Инициализировать как go module.
	Использовать библиотеку https://github.com/beevik/ntp.
	Написать программу печатающую текущее время / точное время с использованием этой библиотеки.

	Программа должна быть оформлена с использованием как go module.
	Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и возвращать ненулевой код выхода в OS.
	Программа должна проходить проверки go vet и golint.
*/

func main() {
	time := time.Now()
	response, err := ntp.Query("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	err = response.Validate()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println("System clock time:", time)
	fmt.Println("NTP clock time:", response.Time)
}
