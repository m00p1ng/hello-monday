package main

import (
	"fmt"
	"time"
)

var thaiWeekDay = [7]string{
	"อาทิตย์",
	"จันทร์",
	"อังคาร",
	"พุธ",
	"พฤหัสบดี",
	"ศุกร์",
	"เสาร์",
}

var thaiMonth = [12]string{
	"มกราคม",
	"กุมภาพันธ์",
	"มีนาคม",
	"เมษายน",
	"พฤษภาคม",
	"มิถุนายน",
	"กรกฎาคม",
	"สิงหาคม",
	"กันยายน",
	"ตุลาคม",
	"พฤศจิกายน",
	"ธันวาคม",
}

func getCurrentThaiDate() string {
	now := time.Now()
	weekday := int(now.Weekday())
	day := now.Day()
	month := now.Month()
	year := now.Year() + 543
	hour := now.Hour()
	minute := now.Minute()

	return fmt.Sprintf(
		"วัน%sที่ %d เดือน %s พ.ศ. %d เวลา %02d:%02d",
		thaiWeekDay[weekday],
		day,
		thaiMonth[month-1],
		year,
		hour,
		minute,
	)
}
