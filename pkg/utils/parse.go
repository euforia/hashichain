package utils

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

// ParseKeyValuePairs parses a set of key value pairs one per line
func ParseKeyValuePairs(rd io.Reader) (map[string]string, error) {
	buf := bufio.NewReader(rd)

	out := make(map[string]string)

	for {

		line, _, err := buf.ReadLine()
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			return out, err
		}

		line = bytes.TrimSpace(line)
		if len(line) == 0 {
			continue
		}

		i := bytes.IndexRune(line, '=')
		if i < 1 {
			return nil, fmt.Errorf("not a key-value pair: '%s'", line)
		}

		key := strings.TrimSpace(string(line[:i]))
		val := strings.TrimSpace(string(line[i+1:]))
		out[key] = val

	}
}
