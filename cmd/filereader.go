package main

import (
	"bytes"
	"fmt"
	"io"
	"log"

	"github.com/Tihmmm/http-server/internal/request"
	"github.com/Tihmmm/http-server/pkg"
)

func readFile(f io.ReadCloser) <-chan string {
	ch := make(chan string, 1)

	go func() {
		defer f.Close()
		defer close(ch)

		var str string
		for {
			data := make([]byte, 8)
			n, err := f.Read(data)
			if err != nil {
				break
			}

			data = data[:n]
			if i := bytes.IndexByte(data, '\n'); i != -1 {
				str += string(data[:i])
				data = data[i+1:]
				ch <- str
				str = ""
			}

			str += string(data)
		}

		if len(str) > 0 {
			log.Printf("read: %s\n", str)
		}
	}()

	return ch
}

func main() {
	fmt.Println(pkg.RemoveFirstNLines("aaa\nbbb\nccc\nddd", "\n", -1))
	fmt.Println(request.ParseRequestLine("GET /coffee HTTP/1.1\r\nHost: localhost:42069\r\nUser-Agent: curl/7.81.0\r\nAccept: */*\r\n\r\n"))
}
