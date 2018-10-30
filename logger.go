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
var appName, logPath string
var level int
var isProd bool

func init() {
	logger = log.New(os.Stdout, appName, log.Lshortfile|log.Ldate|log.Ltime) // log.Lshortfile|log.Ldate|log.Ltime
	appName = "[ default appName ] "
}

// SetIsProd 设置是否生产环境
func SetIsProd(prod bool) {
	isProd = prod
}

func SetLogFilePath(path string) {
	logPath = path
}

// SetAppName 设置app名称 日志打印的第一个
func SetAppName(name string) {
	name = "[ " + name + " ] "
	appName = name
	logger.SetPrefix(name)
}

// SetLogLevel 设置日志等级 1 debug信息不输出 2 info信息不输出 3 警告信息不输出 4 全都不输出
func SetLogLevel(l int) {
	level = l
}

// Print 打印结构
func Print(v ...interface{}) {
	s := formatValue(v)
	doOutput(isProd, s, logger)
}

// Println 打印结构
func Println(v ...interface{}) {
	s := formatValue(v)
	doOutput(isProd, s, logger)
}

// Printf 格式化打印
func Printf(f string, v ...interface{}) {
	s := fmt.Sprintf(f, v...)
	doOutput(isProd, s, logger)
}

// Fatal 失败
func Fatal(v ...interface{}) {
	s := formatValue(v)
	doOutput(isProd, s, logger)
}

// Fatalln 失败
func Fatalln(v ...interface{}) {
	s := formatValue(v)
	doOutput(isProd, s, logger)
}

// Fatalf 格式化输出失败
func Fatalf(f string, v ...interface{}) {
	s := fmt.Sprintf(f, v...)
	doOutput(isProd, s, logger)
}

// Debug debug信息输出
func Debug(f string, v ...interface{}) {
	if level > 0 {
		return
	}
	s := fmt.Sprintf(f, v...)
	doOutputWithPrefix(isProd, s, "debug", logger)
}

// Infof info信息输出
func Infof(f string, v ...interface{}) {
	if level > 1 {
		return
	}
	s := fmt.Sprintf(f, v...)
	doOutputWithPrefix(isProd, s, "info", logger)
}

// Info info信息
func Info(v ...interface{}) {
	if level > 1 {
		return
	}
	s := formatValue(v)
	doOutputWithPrefix(isProd, s, "info", logger)
}

// Warnf 警告信息
func Warnf(f string, v ...interface{}) {
	if level > 2 {
		return
	}
	s := fmt.Sprintf(f, v...)
	doOutputWithPrefix(isProd, s, "warn", logger)
}

// Warn 警告
func Warn(v ...interface{}) {
	if level > 2 {
		return
	}
	s := formatValue(v)
	doOutputWithPrefix(isProd, s, "warn", logger)
}

// Errorf 错误信息格式化输出
func Errorf(f string, v ...interface{}) {
	if level > 3 {
		return
	}
	s := fmt.Sprintf(f, v...)
	doOutputWithPrefix(isProd, s, "error", logger)
}

// Error 错误
func Error(v ...interface{}) {
	if level > 3 {
		return
	}
	s := formatValue(v)
	doOutputWithPrefix(isProd, s, "error", logger)
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

func doOutput(isProd bool, s string, l *log.Logger) {
	if isProd {
		write2File("./logs/"+getYearMonthDay()+".log", s, l)
	} else {
		l.Output(3, s+"\n")
	}
}

func doOutputWithPrefix(isProd bool, s, p string, l *log.Logger) {
	if isProd {
		if logPath != "" {
			write2File(logPath+"/"+getYearMonthDay()+"-"+p+".log", "[ "+p+" ] "+s, l)
		} else {
			write2File("./logs/"+getYearMonthDay()+"-"+p+".log", "[ "+p+" ] "+s, l)
		}
	} else {
		l.Output(3, "[ "+p+" ] "+s+"\n")
	}
}
