package logger

import (
	"io"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func checkFileIsExist(filename string) (bool) {
	var exist = true;
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false;
	}
	return exist;
}

func write2File(filePath, s string) error {
	var f *os.File
	var err1 error;
	if checkFileIsExist(filePath) {
		f, err1 = os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 0666)
	} else {
		f, err1 = os.Create(filePath)
	}
	if err1 != nil {
		return err1
	}
	if _, err1 = io.WriteString(f, s+"\n"); err1 != nil {
		return err1
	}

	return nil
}