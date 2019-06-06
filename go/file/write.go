package file

import (
	"io/ioutil"
	"os"
)

// 对于文件的读写都可以通过"io/ioutil"来实现

// Read : get file content
func Read(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}

// Write : write file
func Write(path string, writeBytes []byte, mode os.FileMode) error {
	return ioutil.WriteFile(path, writeBytes, mode)
}
