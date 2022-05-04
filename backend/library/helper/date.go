package helper

import (
	"strings"
	"time"

	"github.com/gogf/gf/os/gtime"
)

// GetBetweenDates 根据开始日期和结束日期计算出时间段内所有日期
// 参数为日期字符串格式，如：2020-01-01
func GetBetweenDates(sdate, edate string) []string {
	d := []string{}
	timeFormatTpl := "2006-01-02 15:04:05"
	date, err := time.Parse(timeFormatTpl[0:len(sdate)], sdate)
	if err != nil {
		// 时间解析，异常
		return d
	}
	date2, err := time.Parse(timeFormatTpl[0:len(edate)], edate)
	if err != nil {
		// 时间解析，异常
		return d
	}
	if date2.Before(date) {
		// 如果结束时间小于开始时间，异常
		return d
	}
	// 输出日期格式固定
	timeFormatTpl = "2006-01-02"
	date2Str := date2.Format(timeFormatTpl)
	date1Str := date.Format(timeFormatTpl)
	d = append(d, date1Str)
	if date1Str == date2Str {
		return d
	}
	for {
		date = date.AddDate(0, 0, 1)
		dateStr := date.Format(timeFormatTpl)
		d = append(d, dateStr)
		if dateStr == date2Str {
			break
		}
	}
	return d
}

// GetDiffDay 计算两个日期差，默认返回天数
// 参数为日期字符串格式，如：2020-01-01
func GetDiffDay(sdate, edate string) int {
	timeFormatTpl := "2006-01-02 15:04:05"
	date, err := time.Parse(timeFormatTpl[0:len(sdate)], sdate)
	if err != nil {
		// 时间解析，异常
		return 0
	}
	date2, err := time.Parse(timeFormatTpl[0:len(edate)], edate)
	if err != nil {
		// 时间解析，异常
		return 0
	}
	return int(date2.Sub(date).Hours() / 24)
}

// GetWeekDate 根据指定时间和1~7之间的数字（周几），获取对应周几的日期
// date 参数为日期字符串格式，如：2020-01-01
// weekday 参数为数字，如：1表示周一
// format 参数为字符串，如：Y-m-d H:i:s
func GetWeekDate(date string, weekday int, format ...string) string {
	if len(format) == 0 {
		format = []string{"Y-m-d"}
	}
	timeFormatTpl := "2006-01-02 15:04:05"
	d, err := time.Parse(timeFormatTpl[0:len(date)], date)
	if err != nil {
		// 时间解析，异常
		return ""
	}
	dateNew := d.AddDate(0, 0, weekday-int(d.Weekday()))
	return gtime.New(dateNew).Format(format[0])
}

// GetDaysBetweenDate 获取两个时间之间年份、月份，日数据
// 参数为日期字符串格式，如：2020-01-01  2020-03-01
// format 格式：Ymd、Ym、Y
// prefix 格式:dwd_admin_
func GetDaysBetweenDate(sdate, edate, format string, prefix ...string) []string {
	if sdate > edate {
		sdateTmp := sdate
		sdate = edate
		edate = sdateTmp
	}
	startRet := ""
	endRet := ""
	data := []string{}
	plusYears := 0
	plusMonths := 0
	plusDays := 0
	prefixStr := ""
	lastRet := ""
	if len(prefix) > 0 {
		prefixStr = prefix[0]
	}
	if strings.Contains(format, "d") {
		plusDays = 1
	} else if strings.Contains(format, "m") {
		plusMonths = 1
	} else if strings.Contains(format, "Y") {
		plusYears = 1
	}
	if sdate != "" {
		startRet = gtime.NewFromStr(sdate).Format(format)
		data = append(data, prefixStr+startRet)
	}

	if edate != "" {
		endRet = gtime.NewFromStr(edate).Format(format)
		if startRet != endRet {
			lastRet = prefixStr + endRet
		}
	}

	if startRet != "" && endRet != "" {
		if startRet != endRet {
			for {
				if plusYears == 0 && plusMonths == 0 && plusDays == 0 {
					break
				}
				st := gtime.NewFromStr(sdate).AddDate(plusYears, plusMonths, plusDays)
				startRet = st.Format(format)
				if startRet == endRet {
					break
				}
				sdate = st.String()
				data = append(data, prefixStr+startRet)
			}
		}
	} else {
		//如果时间参数为空，则默认当前时间
		data = append(data, prefixStr+gtime.Now().Format(format))
	}
	if lastRet != "" {
		data = append(data, lastRet)
	}
	return data
}

// 计算年份差距  (生日)
func GetYearDiffer(start_time, end_time string) (age int, err error) {
	var pslTime string
	if strings.Contains(start_time, ".") {
		pslTime = "2006.01.02"
	} else if strings.Contains(start_time, "-") {
		pslTime = "2006-01-02"
	} else {
		pslTime = "2006/01/02"
	}
	// 开始时间
	t1, err := time.ParseInLocation(pslTime, start_time, time.Local)
	if err != nil {
		return
	}
	// 结束时间
	t2, err := time.ParseInLocation(pslTime, end_time, time.Local)
	if err != nil {
		return
	}
	if t1.Before(t2) {
		diff := t2.Unix() - t1.Unix()
		age = int(diff / (3600 * 365 * 24))

	}
	return
}
