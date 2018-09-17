# logger

## 日志记录工具类

>USE go get -u -v github.com/MWangxj/logger

        package logger
            import "github.com/MWangxj/logger"


        FUNCTIONS

        func Debug(f string, v ...interface{})
            Debug debug信息输出

        func Error(v ...interface{})
            Error 错误

        func Errorf(f string, v ...interface{})
            Errorf 错误信息格式化输出

        func Fatal(v ...interface{})
            Fatal 失败

        func Fatalf(f string, v ...interface{})
            Fatalf 格式化输出失败

        func Fatalln(v ...interface{})
            Fatalln 失败

        func Info(v ...interface{})
            Info info信息

        func Infof(f string, v ...interface{})
            Infof info信息输出

        func Print(v ...interface{})
            Print 打印结构

        func Printf(f string, v ...interface{})
            Printf 格式化打印

        func Println(v ...interface{})
            Println 打印结构

        func SetAppName(name string)
            SetAppName 设置app名称 日志打印的第一个

        func SetIsProd(prod bool)
            SetIsProd 设置是否生产环境

        func SetLogLevel(l int)
            SetLogLevel 设置日志等级 1 debug信息不输出 2 info信息不输出 3 警告信息不输出 4 全都不输出

        func Warn(v ...interface{})
            Warn 警告

        func Warnf(f string, v ...interface{})
            Warnf 警告信息
