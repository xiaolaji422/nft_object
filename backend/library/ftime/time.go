package ftime

import "time"

const default_layout = "2006-01-02 15:04:05"

func StrToTime(str string, layout ...string) (time.Time, error) {
	layout_str := default_layout
	if len(layout) > 0 {
		layout_str = layout[0]
	}
	return time.Parse(layout_str, str)
}
