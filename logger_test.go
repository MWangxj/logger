package logger

import (
	`log`
	`testing`
)

func TestLogInfo(t *testing.T) {

}

func TestLogWarnning(t *testing.T) {

}

func TestLogError(t *testing.T) {

}

func TestWrite2File(t *testing.T) {
	SetAppName("test")
	err := write2File("./test.log", "wangxianjin12312", logger)
	if err != nil {
		t.Fatal("fail")
		t.Fatalf("%s", err.Error())
	}
}

func TestYearMonthDay(t *testing.T) {
	s := getYearMonthDay()
	t.Log(s)
}

func TestLogger2File(t *testing.T) {
	write2File("./test/test.log", "123", logger)
}

func TestSplitPath(t *testing.T) {
	path := "/user/local/log.test"
	dir := getPathDir(path)
	t.Log("----->" + dir)
}

type test struct {
	Prepoty0 string `json:"prepoty_0"`
	Prepoty1 string `json:"prepoty_1"`
}

func TestSerial(t *testing.T) {
	ts := &test{}
	ts.Prepoty0 = "0"
	ts.Prepoty1 = "1"
	SetIsProd(true)
	Print(ts)
	log.Print(ts)
}
