package logger

import (
	`fmt`
	json "github.com/json-iterator/go"
	`log`
	`os`
	`strconv`
	`time`
	`unsafe`
)

var logger *log.Logger
var prefix string
var level int
var isProd bool

func init() {
	logger = log.New(os.Stdout, prefix, log.Lshortfile|log.Ldate|log.Ltime)
}

func SetIsProd(prod bool) {
	isProd = prod
}

func SetAppName(name string) {
	logger.SetPrefix("[ " + name + " ] ")
	prefix = name
}

func SetLogLevel(l int) {
	level = l
}

func Print(v ...interface{}) {
	s := formatValue(v)
	if isProd {
		write2File("./logs/"+getYearMonthDay()+".log", s, logger)
	} else {
		logger.Print(s)
	}
}

func Println(v ...interface{}) {
	s := formatValue(v)
	if isProd {
		write2File("./logs/"+getYearMonthDay()+".log", s, logger)
	} else {
		logger.Println(s)
	}
}

func Printf(f string, v ...interface{}) {
	if isProd {
		s := fmt.Sprintf(f, v)
		write2File("./logs/"+getYearMonthDay()+".log", s, logger)
	} else {
		logger.Printf(f, v)
	}
}

func Fatal(v ...interface{}) {
	s := formatValue(v)
	if isProd {
		write2File("./logs/"+getYearMonthDay()+".log", s, logger)
	} else {
		logger.Fatal(s)
	}
}

func Fatalln(v ...interface{}) {
	s := formatValue(v)
	if isProd {
		write2File("./logs/"+getYearMonthDay()+".log", s, logger)
	} else {
		logger.Fatalln(s)
	}
}

func Fatalf(f string, v ...interface{}) {
	if isProd {
		s := fmt.Sprintf(f, v)
		write2File("./logs/"+getYearMonthDay()+".log", s, logger)
	} else {
		logger.Fatalf(f, v)
	}
}

func Debug(f string, v ...interface{}) {
	if level > 0 {
		return
	}
	if isProd {
		s := fmt.Sprintf(f, v)
		write2File("./logs/"+getYearMonthDay()+".log", s, logger)
	} else {
		logger.Printf("[ debug ] "+f, v)
	}
}

func Infof(f string, v ...interface{}) {
	if level > 1 {
		return
	}
	if isProd {
		s := fmt.Sprintf(f, v)
		write2File("./logs/"+getYearMonthDay()+".log", s, logger)
	} else {
		logger.Printf("[ info ] "+f, v)
	}
}

func Info(v ...interface{}) {
	if level > 1 {
		return
	}
	s := formatValue(v)
	if isProd {
		write2File("./logs/"+getYearMonthDay()+".log", s, logger)
	} else {
		logger.Printf("[ info ] " + s)
	}
}

func Warnf(f string, v ...interface{}) {
	if level > 2 {
		return
	}
	if isProd {
		s := fmt.Sprintf(f, v)
		write2File("./logs/"+getYearMonthDay()+".log", s, logger)
	} else {
		logger.Printf("[ warn ] "+f, v)
	}
}

func Warn(v ...interface{}) {
	if level > 2 {
		return
	}
	s := formatValue(v)
	if isProd {
		write2File("./logs/"+getYearMonthDay()+".log", s, logger)
	} else {
		logger.Printf("[ warn ] " + s)
	}
}

func Errorf(f string, v ...interface{}) {
	if level > 3 {
		return
	}
	if isProd {
		s := fmt.Sprintf(f, v)
		write2File("./logs/"+getYearMonthDay()+".log", s, logger)
	} else {
		logger.Printf("[ error ] "+f, v)
	}
}

func Error(v ...interface{}) {
	if level > 3 {
		return
	}
	s := formatValue(v)
	if isProd {
		write2File("./logs/"+getYearMonthDay()+".log", s, logger)
	} else {
		logger.Printf("[ error ] " + s)
	}
}

func getYearMonthDay() string {
	n := time.Now()
	y := n.Year()
	m := n.Month()
	d := n.Day()
	return prefix + strconv.Itoa(y) + month2string(m) + strconv.Itoa(d)
}

func month2string(m time.Month) string {
	d := (int)(m)
	if d < 10 {
		return "0" + strconv.Itoa(d)
	}
	return strconv.Itoa(d)
}

// string2bytes String to bytes
func string2bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

// bytes2string Bytes to string
func bytes2string(b []byte) string {
	h := (*[3]uintptr)(unsafe.Pointer(&b))
	x := [2]uintptr{h[0], h[1]}
	return *(*string)(unsafe.Pointer(&x))
}

func formatValue(v ...interface{}) string {
	data, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return bytes2string(data)
}
