package pkg

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"strings"
)

func IsAplhabetic(str string) bool {
	buf := bytes.NewBufferString(str)
	for {
		c, err := buf.ReadByte()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			return false
		}

		if !(c >= 'A' && c <= 'Z') {
			return false
		}
	}

	return true
}

func RemoveFirstNLines(str, sep string, n int) (string, error) {
	if n <= 0 {
		return "", fmt.Errorf("n must be greater than zero")
	}

	split := strings.SplitN(str, sep, n+1)
	if len(split) <= n {
		return "", fmt.Errorf("n larger than number of lines")
	}

	return split[n], nil
}
