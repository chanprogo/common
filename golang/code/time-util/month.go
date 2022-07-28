package timeutil

import (
	"strconv"
	"time"
)

// GetMonthStartAndEnd
func GetMonthStartAndEnd(myYear string, myMonth string) map[string]string {

	if len(myMonth) == 1 {
		myMonth = "0" + myMonth
	}
	yInt, _ := strconv.Atoi(myYear)

	loc, _ := time.LoadLocation("Local")

	theTime, _ := time.ParseInLocation("2006-01-02 15:04:05", myYear+"-"+myMonth+"-01 00:00:00", loc)
	newMonth := theTime.Month()

	t1 := time.Date(yInt, newMonth, 1, 0, 0, 0, 0, time.Local).Format("2006-01-02 15:04:05")
	t2 := time.Date(yInt, newMonth+1, 0, 0, 0, 0, 0, time.Local).Format("2006-01-02 15:04:05")
	result := map[string]string{"start": t1, "end": t2}
	return result
}

// 获取月份开始、结束
func GetMonthStartEnd(dd time.Time) (start time.Time, end time.Time) {
	year, month, _ := dd.Date()
	loc := dd.Location()

	startOfMonth := time.Date(year, month, 1, 0, 0, 0, 0, loc)
	endOfMonth := startOfMonth.AddDate(0, 1, -1)
	return startOfMonth, endOfMonth
}
