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

	numbers, err := strconv.Atoi(num)
	if err != nil {
		return 0, err
	}
	return numbers, nil
}

func ReadPrice(reader *bufio.Reader) (float64, error) {
	num, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	num = strings.TrimSuffix(num, "\r\n")

	price, err := strconv.ParseFloat(num, 64)
	if err != nil {
		return 0, err
	}
	return price, nil
}

func ReadLine(reader *bufio.Reader) (string, error) {
	line, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	line = strings.TrimSuffix(line, "\r\n")

	return line, nil
}