package request

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

type Request struct {
	RequestLine RequestLine
	Headers     map[string]string
	Body        []byte
}

type RequestLine struct {
	Method        string
	RequestTarget string
	HttpVersion   string
}

var ErrorBadRequestLine = fmt.Errorf("malformed request line")

const (
	SEPARATOR = "\r\n"
)

func ParseRequestLine(s string) (*RequestLine, string, error) {
	idx := strings.Index(s, SEPARATOR)
	if idx == -1 {
		return nil, s, nil
	}

	startLine := s[:idx]
	restOfMsg := s[idx+len(SEPARATOR):]

	parts := strings.Split(startLine, " ")
	if len(parts) != 3 {
		return nil, "", ErrorBadRequestLine
	}

	versionParts := strings.Split(parts[2], "/")
	if len(versionParts) != 2 || versionParts[0] != "HTTP" || versionParts[1] != "1.1" {
		return nil, "", ErrorBadRequestLine
	}
	httpVersion := versionParts[1]

	rl := &RequestLine{
		Method:        parts[0],
		RequestTarget: parts[1],
		HttpVersion:   httpVersion,
	}

	return rl, restOfMsg, nil
}

func RequestFromReader(r io.Reader) (*Request, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, errors.Join(fmt.Errorf("error reading request"), err)
	}

	str := string(data)
	rl, _, err := ParseRequestLine(str)
	if err != nil {
		return nil, err
	}

	return &Request{
		RequestLine: *rl,
	}, nil
}
