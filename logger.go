package logger

import (
	`fmt`
	json "github.com/json-iterator/go"
	`log`
	`os`
	`runtime`
	`strconv`
	`time`
	`unsafe`
)

var logger *log.Logger
var appName string
var level int
var isProd bool

func init() {
	logger = log.New(os.Stdout, appName, log.Lshortfile|log.Ldate|log.Ltime) // log.Lshortfile|log.Ldate|log.Ltime
}

func SetIsProd(prod bool) {
	isProd = prod
}

func SetAppName(name string) {
	name = "[ " + name + " ] "
	appName = name
	logger.SetPrefix(name)
}

func SetLogLevel(l int) {
	level = l
}

func Print(v ...interface{}) {
	s := formatValue(v)
	if isProd {
		write2File("./logs/"+getYearMonthDay()+".log", s, logger)
	} else {
		logger.Output(2,s)
	}
}

func Println(v ...interface{}) {
	s := formatValue(v)
	if isProd {
		write2File("./logs/"+getYearMonthDay()+".log", s, logger)
	} else {
		logger.Output(2,s+"\n")
	}
}

func Printf(f string, v ...interface{}) {
	s := fmt.Sprintf(f, v...)
	if isProd {
		write2File("./logs/"+getYearMonthDay()+".log", s, logger)
	} else {
		logger.Output(2,s+"\n")
	}
}

func Fatal(v ...interface{}) {
	s := formatValue(v)
	if isProd {
		write2File("./logs/"+getYearMonthDay()+".log", s, logger)
	} else {
		logger.Output(2,s+"\n")
	}
}

func Fatalln(v ...interface{}) {
	s := formatValue(v)
	if isProd {
		write2File("./logs/"+getYearMonthDay()+".log", s, logger)
	} else {
		logger.Output(2,s+"\n")
	}
}

func Fatalf(f string, v ...interface{}) {
	s := fmt.Sprintf(f, v...)
	if isProd {
		write2File("./logs/"+getYearMonthDay()+".log", s, logger)
	} else {
		logger.Output(2,s+"\n")
	}
}

func Debug(f string, v ...interface{}) {
	if level > 0 {
		return
	}
	s := fmt.Sprintf(f, v...)
	if isProd {
		write2File("./logs/"+getYearMonthDay()+".log", s, logger)
	} else {
		logger.Output(2,"[ debug ] "+s+"\n")
	}
}

func Infof(f string, v ...interface{}) {
	if level > 1 {
		return
	}
	s := fmt.Sprintf(f, v...)
	if isProd {
		write2File("./logs/"+getYearMonthDay()+".log", s, logger)
	} else {
		logger.Output(2,"[ info ] "+s+"\n")
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
		logger.Output(2,"[ info ] "+s+"\n")
	}
}

func Warnf(f string, v ...interface{}) {
	if level > 2 {
		return
	}
	s := fmt.Sprintf(f, v...)
	if isProd {
		write2File("./logs/"+getYearMonthDay()+".log", s, logger)
	} else {
		logger.Output(2,"[ warn ] "+s+"\n")
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
		logger.Output(2,"[ warn ] "+s+"\n")
	}
}

func Errorf(f string, v ...interface{}) {
	if level > 3 {
		return
	}
	s := fmt.Sprintf(f, v...)
	if isProd {
		write2File("./logs/"+getYearMonthDay()+".log", s, logger)
	} else {
		logger.Output(2,"[ error ] "+s+"\n")
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
		logger.Output(2,"[ error ] "+s+"\n")
	}
}

func getYearMonthDay() string {
	n := time.Now()
	y := n.Year()
	m := n.Month()
	d := n.Day()
	return appName + strconv.Itoa(y) + month2string(m) + strconv.Itoa(d)
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

func formatValue(v interface{}) string {
	data, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return bytes2string(data)
}

func getCallerInfo() string {
	pc, f, l, _ := runtime.Caller(2)
	fn :=runtime.FuncForPC(pc).Name()
	return f + " " + strconv.Itoa(l) + " "+fn+" "
}