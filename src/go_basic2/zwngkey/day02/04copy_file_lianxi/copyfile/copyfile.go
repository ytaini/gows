package copyfile

import (
	"io"
	"os"
)

func CopyFile(srcFile, destFile string) (bool, error) {
	file, err := os.OpenFile(srcFile, os.O_RDONLY, 0644)
	if err != nil {
		return false, err
	}
	defer file.Close()

	file1, err1 := os.OpenFile(destFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err1 != nil {
		return false, err
	}
	defer file1.Close()

	_, err2 := io.Copy(file1, file)
	if err2 != nil {
		return false, err
	}
	return true, nil
}
