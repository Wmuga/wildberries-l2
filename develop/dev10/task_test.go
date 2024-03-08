package main

import (
	"bytes"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/Wmuga/wildberries-l2/develop/dev10/telnet"
)

var (
	address = "localhost"
	port    = "1324"
	timeout = 1
	testMsg = "testmsg\n"
)

func TestTimeout(t *testing.T) {
	start := time.Now()
	err := telnet.Connect(address, port, timeout, os.Stdin, os.Stdout)
	dur := time.Since(start)
	if err == nil {
		t.Error("shouldn't connect")
		return
	}

	if dur < time.Second*time.Duration(timeout) || dur > time.Second*time.Duration(timeout+1) {
		t.Error("wrong timeout")
	}
}

func TestConnection(t *testing.T) {
	go SetupLoopback(port)
	reader := bytes.Buffer{}
	writer := bytes.Buffer{}
	writer.WriteString(testMsg)
	go telnet.Connect(address, port, timeout, &writer, &reader)
	// дождаться обратного сообщения
	time.Sleep(time.Millisecond * 100)
	msg := reader.String()
	if !strings.Contains(msg, testMsg) {
		t.Error("didnt recieve same message Got", msg, "Expected", testMsg)
	}
}
