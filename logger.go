package logger

import (
	`encoding/json`
	`fmt`
	_ `fmt`
	`log`
	`os`
	`strconv`
	`time`
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
	if isProd {
		s := fmt.Sprintf("%s", v)
		write2File("./logs/"+getYearMonthDay()+".log", s, logger)
	} else {
		logger.Print(v)
	}
}

func Println(v ...interface{}) {
	if isProd {
		s := fmt.Sprintf("%s", v)
		write2File("./logs/"+getYearMonthDay()+".log", s, logger)
	} else {
		logger.Println(v)
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
	if isProd {
		s := fmt.Sprintf("%s", v)
		write2File("./logs/"+getYearMonthDay()+".log", s, logger)
	} else {
		logger.Fatal(v)
	}
}

func Fatalln(v ...interface{}) {
	if isProd {
		s := fmt.Sprintf("%s", v)
		write2File("./logs/"+getYearMonthDay()+".log", s, logger)
	} else {
		logger.Fatalln(v)
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

func Info(v...interface{})  {
	if level > 1 {
		return
	}
	data,_ := json.Marshal(v)
	s := string(data)
	if isProd {
		write2File("./logs/"+getYearMonthDay()+".log", s, logger)
	} else {
		logger.Printf("[ info ] "+ s)
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

func Warn(v...interface{})  {
	if level > 2 {
		return
	}
	data,_ := json.Marshal(v)
	s := string(data)
	if isProd {
		write2File("./logs/"+getYearMonthDay()+".log", s, logger)
	} else {
		logger.Printf("[ warn ] "+s)
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

func Error(v...interface{})  {
	if level > 3 {
		return
	}
	data,_ := json.Marshal(v)
	s := string(data)
	if isProd {
		write2File("./logs/"+getYearMonthDay()+".log", s, logger)
	} else {
		logger.Printf("[ error ] "+s)
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
