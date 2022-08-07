package utilities

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func NewReader() *bufio.Reader {
	reader := bufio.NewReader(os.Stdin)
	return reader
}

func ReadNumber(reader *bufio.Reader) (int, error) {
	num, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	num = strings.TrimSuffix(num, "\r\n")

	option, err := strconv.Atoi(num)
	if err != nil {
		return 0, err
	}
	return option, nil
}

func ReadLine(reader *bufio.Reader) (string, error) {
	line, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	line = strings.TrimSuffix(line, "\r\n")

	return line, nil
}