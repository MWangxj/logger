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
	appName = "[ default appName ] "
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
	doOutput(isProd,s,logger)
}

func Println(v ...interface{}) {
	s := formatValue(v)
	doOutput(isProd,s,logger)
}

func Printf(f string, v ...interface{}) {
	s := fmt.Sprintf(f, v...)
	doOutput(isProd,s,logger)
}

func Fatal(v ...interface{}) {
	s := formatValue(v)
	doOutput(isProd,s,logger)
}

func Fatalln(v ...interface{}) {
	s := formatValue(v)
	doOutput(isProd,s,logger)
}

func Fatalf(f string, v ...interface{}) {
	s := fmt.Sprintf(f, v...)
	doOutput(isProd,s,logger)
}

func Debug(f string, v ...interface{}) {
	if level > 0 {
		return
	}
	s := fmt.Sprintf(f, v...)
	doOutputWithPrefix(isProd,s,"debug",logger)
}

func Infof(f string, v ...interface{}) {
	if level > 1 {
		return
	}
	s := fmt.Sprintf(f, v...)
	doOutputWithPrefix(isProd,s,"info",logger)
}

func Info(v ...interface{}) {
	if level > 1 {
		return
	}
	s := formatValue(v)
	doOutputWithPrefix(isProd,s,"info",logger)
}

func Warnf(f string, v ...interface{}) {
	if level > 2 {
		return
	}
	s := fmt.Sprintf(f, v...)
	doOutputWithPrefix(isProd,s,"warn",logger)
}

func Warn(v ...interface{}) {
	if level > 2 {
		return
	}
	s := formatValue(v)
	doOutputWithPrefix(isProd,s,"warn",logger)
}

func Errorf(f string, v ...interface{}) {
	if level > 3 {
		return
	}
	s := fmt.Sprintf(f, v...)
	doOutputWithPrefix(isProd,s,"error",logger)
}

func Error(v ...interface{}) {
	if level > 3 {
		return
	}
	s := formatValue(v)
	doOutputWithPrefix(isProd,s,"error",logger)
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
	fn := runtime.FuncForPC(pc).Name()
	return f + " " + strconv.Itoa(l) + " " + fn + " "
}

func doOutput(isProd bool,s string,l *log.Logger)  {
	if isProd {
		write2File("./logs/"+getYearMonthDay()+".log", s, l)
	} else {
		l.Output(2, s+"\n")
	}
}

func doOutputWithPrefix(isProd bool,s,p string,l *log.Logger)  {
	if isProd {
		write2File("./logs/"+getYearMonthDay()+".log", "[ "+p+" ] "+s, l)
	} else {
		l.Output(3, "[ "+p+" ] "+s+"\n")
	}
}