package logger

import (
	`log`
	`os`
)

var logger *log.Logger
var prefix string

func init() {
	logger = log.New(os.Stdout, prefix, log.Lshortfile|log.Ldate|log.Ltime)
}

func SetPrefix(p string)  {
	log.SetPrefix(p)
}

func Print(v...interface{})  {
	logger.Print(v)
}

func Println(v...interface{})  {
	logger.Println(v)
}

func Printf(f string,v...interface{})  {
	logger.Printf(f,v)
}

func Fatal(v...interface{})  {
	logger.Fatal(v)
}

func Fatalln(v...interface{})  {
	logger.Fatalln(v)
}

func Fatalf(f string,v...interface{})  {
	logger.Fatalf(f,v)
}