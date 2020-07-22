package time

import "time"

//获取相差时间
func GetHourDiffer(start_time, end_time string) int64 {
	var hour int64
	t1, err := time.ParseInLocation("2006-01-02 15:04:05", start_time, time.Local)
	t2, err := time.ParseInLocation("2006-01-02 15:04:05", end_time, time.Local)
	if err == nil && t1.Before(t2) {
		diff := t2.Unix() - t1.Unix() //
		hour = diff / 3600
		return hour
	} else {
		return hour
	}
}

// 获取格式化运行时间
func GetRunTime(hour int64) {
	if hour >= 24 && hour <= 24 * 365{
		// 如果大于24小时 并且小于一年

	}
}
