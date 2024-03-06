package main

import (
	"bufio"
	"fmt"
	"os"
	"os/user"
	"strings"

	"github.com/Wmuga/wildberries-l2/develop/dev08/commands"
)

/*
Необходимо реализовать свой собственный UNIX-шелл-утилиту с поддержкой ряда простейших команд:
- cd <args> - смена директории (в качестве аргумента могут быть то-то и то)
- pwd - показать путь до текущего каталога
- echo <args> - вывод аргумента в STDOUT
- kill <args> - "убить" процесс, переданный в качесте аргумента (пример: такой-то пример)
- ps - выводит общую информацию по запущенным процессам в формате *такой-то формат*

Так же требуется поддерживать функционал fork/exec-команд

Дополнительно необходимо поддерживать конвейер на пайпах (linux pipes, пример cmd1 | cmd2 | .... | cmdN).
*/

func GetCommands() map[string]commands.Command {
	return map[string]commands.Command{
		"cd":   commands.CD,
		"pwd":  commands.PWD,
		"kill": commands.Kill,
		"echo": commands.Echo,
		"ps":   commands.PS,
	}
}

func runCommmand(cmd commands.Command, args string, pidInfo bool) {
	pid := os.Getpid()
	if pidInfo {
		fmt.Println("[1]", pid)
		fmt.Println("[1]+ Done", pid)
	}

	out, err := cmd(args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	// Принудительное добавление \n, если его нет
	if strings.HasSuffix(out, "\n") {
		fmt.Print(out)
		return
	}
	fmt.Println(out)
}

func parseInput(input string, cmds map[string]commands.Command) (commands.Command, string, bool) {
	// подтереть от пробелов и лишних символов
	inputSplit := strings.Split(strings.Trim(input, " \r\n"), " ")
	// выковыривание команды и ее аргументов
	cmd := inputSplit[0]
	args := strings.Join(inputSplit[1:], " ")
	forkExec := false
	if strings.HasSuffix(args, "&") {
		forkExec = true
		args = args[:len(args)-1]
	}
	// Выбор команды
	if cmd == "exit" {
		os.Exit(0)
	}
	return cmds[cmd], args, forkExec
}

func main() {
	cmds := GetCommands()
	reader := bufio.NewReader(os.Stdin)
	// Получить данные для промптера
	hostName, err := os.Hostname()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	user, err := user.Current()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	userName := user.Name
	for {
		// Рабочая директория
		wd, err := os.Getwd()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		wdSplit := strings.Split(wd, "/")
		wd = wdSplit[len(wdSplit)-1]
		// Вывод промптера
		fmt.Printf("%s@%s:[%s]$ ", userName, hostName, wd)
		// Ввод команд
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		// Получение команды, аргументов
		handler, args, forkExec := parseInput(input, cmds)
		if handler == nil {
			fmt.Fprintln(os.Stderr, "Can't find command")
			continue
		}
		// Вызов команды
		if forkExec {
			go runCommmand(handler, args, true)
			continue
		}
		runCommmand(handler, args, false)
	}
}
