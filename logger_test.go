package logger

import (
	`os`
	`testing`
)

func TestLogInfo(t *testing.T)  {

}

func TestLogWarnning(t *testing.T)  {

}

func TestLogError(t *testing.T)  {

}

func TestWrite2File(t *testing.T)  {
	err := write2File("./test.log","wangxianjin12312")
	if err != nil {
		t.Fatal("fail")
		t.Fatalf("%s",err.Error())
	}
}

func TestYearMonthDay(t *testing.T)  {
	s := getYearMonthDay()
	t.Log(s)
}

func TestLogger2File(t *testing.T)  {
	logger.Output(2,"asdfsaf")
	var f *os.File
	var err1 error;
	if checkFileIsExist("./test.log") {
		f, err1 = os.OpenFile("./test.log", os.O_RDWR|os.O_APPEND, 0666)
	} else {
		f, err1 = os.Create("./test.log")
	}
	if err1 != nil {
		return
	}
	logger.SetOutput(f)
	logger.Output(2,"sdfafs")
}
