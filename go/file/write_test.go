package file_test

import (
	"fmt"
	"testing"

	"github.com/qinhan-shu/study/go/file"
)

func TestRead(t *testing.T) {
	path := "/Users/qinhan/go/src/github.com/qinhan-shu/study/go/file/check.go"
	fileBytes, err := file.Read(path)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("%s", fileBytes)
}

func TestWrite(t *testing.T) {
	path := "/Users/qinhan/go/src/github.com/qinhan-shu/study/go/file/2.txt"
	writeBytes := []byte("qinhan")
	if err := file.Write(path, writeBytes, 0666); err != nil {
		t.Error(err)
		return
	}
}
