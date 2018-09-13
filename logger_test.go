package logger

import (
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
