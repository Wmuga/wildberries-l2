package telnet

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

var (
	ErrConnectTimeout = errors.New("timed out connection")
)

func tryConnect(address, port string, timeout int) (net.Conn, error) {
	// timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	var con net.Conn
	var err error
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				con, err = net.Dial("tcp", address+":"+port)
				if err == nil {
					cancel()
					return
				}
			}
		}
	}()
	<-ctx.Done()
	// Проверки что удачно подключились
	if err != nil {
		return nil, err
	}

	if con == nil {
		return nil, ErrConnectTimeout
	}

	return con, nil
}

func Connect(address, port string, timeout int, inp io.Reader, outp io.Writer) error {
	con, err := tryConnect(address, port, timeout)
	if err != nil {
		return err
	}
	fmt.Println("Connected to", address, port)
	done := make(chan struct{})
	// Читатель
	go func() {
		conReader := bufio.NewReader(con)
		for {
			select {
			case <-done:
				return
			default:
				str, err := conReader.ReadString('\n')
				if err != nil {
					close(done)
					return
				}
				fmt.Fprintln(outp, str)
			}
		}
	}()
	go func() {
		// Отправлялка STDIN в соединение
		reader := bufio.NewReader(inp)
		for {
			str, err := reader.ReadString('\n')
			// завершение по CTRL+D
			if errors.Is(err, io.EOF) {
				// предотвратить закрытие закрытого канала
				select {
				case <-done:
				default:
					close(done)
				}
				return
			}
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue
			}
			// Отправить в соединение
			_, err = fmt.Fprint(con, str)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				close(done)
				continue
			}
		}
	}()
	<-done
	return nil
}
