package helpers

import (
	"log"
	"os"
	"testing"
)

func TestContents(t *testing.T) {
	tt := []struct {
		description string
		content     string
		path        string
		expecterr   error
		readPath    string
	}{
		{"Empty", "", "testFile.txt", nil, "testFile.txt"},
		{"specSymbols", "!@#$%^&*()_+~`'/\\", "testFile.txt", nil, "testFile.txt"},
		{"enters", "\n\n\n\n", "testFile.txt", nil, "testFile.txt"},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			f, err := os.Create(tc.path)
			if err != nil {
				log.Fatal(err)
			}
			f.Write([]byte(tc.content))
			content, err := Contents(tc.readPath)
			if content != tc.content || err != tc.expecterr {
				t.Errorf("got %s with error %v, but expected %s with error %v", content, err, tc.content, tc.expecterr)
			}
			err = os.Remove("testFile.txt")
			if err != nil {
				log.Fatal(err)
			}
		})
	}
}
