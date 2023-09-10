package files

import (
	"io/ioutil"
	"strings"
)

func Write(filename string, text string) error {
	err := ioutil.WriteFile(filename, []byte(text), 0644)
	return err
}

func Read(filename string) (string, error) {
	data, err := ioutil.ReadFile(filename)
	str := string(data)
	str = strings.ReplaceAll(str, "\r", "")
	return str, err
}

func GetWidth() int {
	const maxInt = 9223372036854775807
	return maxInt
}
