package util

import "time"

var (
	//时间转换的模板，golang里面只能是 "2006-01-02 15:04:05" （go的诞生时间）
	timeTemplate1 = "2006-01-02 15:04:05" //常规类型
	// timeTemplate2 = "2006/01/02 15:04:05" //其他类型
	// timeTemplate3 = "2006-01-02"          //其他类型
	// timeTemplate4 = "15:04:05"            //其他类型
)

func TimeStringToInt64(timeString string) int64 {
	time, _ := time.ParseInLocation(timeTemplate1, timeString, time.Local)
	return time.Unix()
}
