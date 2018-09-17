package logger

import (
	`log`
	"os"
	`strings`
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

func write2File(filePath, s string, l *log.Logger) error {
	var f *os.File
	var err1 error;
	if checkFileIsExist(filePath) {
		f, err1 = os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 0666)
	} else {
		if err1 = os.MkdirAll(getPathDir(filePath), 0700); err1 != nil {
			return err1
		}
		f, err1 = os.Create(filePath)
	}
	if err1 != nil {
		return err1
	}
	l.SetOutput(f)

	return l.Output(4, s+"\n")
}

func getPathDir(filePath string) string {
	s := strings.Split(filePath, "/")
	return strings.Join(s[:len(s)-1], "/")
}
