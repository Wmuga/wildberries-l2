package commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/mitchellh/go-ps"
)

type Command func(args string) (stdout string, stderr error)

func CD(args string) (string, error) {
	err := os.Chdir(args)
	if err != nil {
		return "", err
	}
	return "", nil
}

func PWD(_ string) (string, error) {
	return os.Getwd()
}

func Echo(args string) (string, error) {
	return args, nil
}

func Kill(args string) (string, error) {
	pc, err := os.FindProcess(0)
	if err != nil {
		return "", err
	}
	return "", pc.Kill()
}

func PS(_ string) (string, error) {
	pss, err := ps.Processes()
	if err != nil {
		return "", err
	}

	builder := strings.Builder{}
	builder.WriteString("  PID  PPID NAME\n")
	for _, ps := range pss {
		builder.WriteString(fmt.Sprintf("%5d %5d %s\n", ps.Pid(), ps.PPid(), ps.Executable()))
	}
	return builder.String(), nil
}
