package validator

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Validator struct {
	errors []error
}

func New() *Validator {
	return &Validator{}
}

func (v *Validator) LogErrors() {
	if len(v.errors) != 0 {
		for _, err := range v.errors {
			log.Print(err)
		}
		os.Exit(0)
	}
}

func (v *Validator) IsFileHasTxtExtension(filename string) bool {
	if len(filename) <= 4 {
		return false
	}
	return filename[len(filename)-4:] == ".txt"
}

func (v *Validator) IsArrayContainsString(strs []string, s string) bool {
	for _, str := range strs {
		if str == s {
			return true
		}
	}
	return false
}

func (v *Validator) IsMd5HashEqual(pathToFile string, hash string) bool {
	fileContents, err := ioutil.ReadFile(pathToFile)
	if err != nil {
		v.errors = append(v.errors, err)
		return false
	}

	fileHash := hashFile(fileContents)

	return fileHash == hash
}

func (v *Validator) IsTextAscii(text string) bool {
	for _, ch := range text {
		if (int(ch) < 32 || int(ch) > 126) && int(ch) != 10 && int(ch) != 9 {
			return false
		}
	}
	return true
}

func hashFile(data []byte) string {
	h := md5.Sum(data)
	return fmt.Sprintf("%x", h)
}
