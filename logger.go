package logger

import (
	`log`
	`os`
)

var logger *log.Logger
var prefix string
var level int

func init() {
	logger = log.New(os.Stdout, prefix, log.Lshortfile|log.Ldate|log.Ltime)
}

func SetAppName(name string) {
	logger.SetPrefix("[ " + name + " ] ")
}

func SetLogLevel(l int)  {
	level = l
}

func Print(v ...interface{}) {
	logger.Print(v)
}

func Println(v ...interface{}) {
	logger.Println(v)
}

func Printf(f string, v ...interface{}) {
	logger.Printf(f, v)
}

func Fatal(v ...interface{}) {
	logger.Fatal(v)
}

func Fatalln(v ...interface{}) {
	logger.Fatalln(v)
}

func Fatalf(f string, v ...interface{}) {
	logger.Fatalf(f, v)
}

func Debug(f string, v ...interface{}) {
	if level > 0 {
		return
	}
	logger.Printf("[ debug ] "+f, v)
}

func Info(f string, v ...interface{}) {
	if level > 1 {
		return
	}
	logger.Printf("[ info ] "+f, v)
}

func Warn(f string, v ...interface{}) {
	if level > 2 {
		return
	}
	logger.Printf("[ warn ] "+f, v)
}

func Error(f string, v ...interface{}) {
	if level > 3 {
		return
	}
	logger.Printf("[ error ] "+f, v)
}
