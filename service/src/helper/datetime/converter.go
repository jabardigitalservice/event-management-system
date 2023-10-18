package datetime

import (
	"fmt"
	"time"
)

var IndonesiaDate = map[string]string{
	"Sunday":    "Minggu",
	"Monday":    "Senin",
	"Tuesday":   "Selasa",
	"Wednesday": "Rabu",
	"Thursday":  "Kamis",
	"Friday":    "Jum'at",
	"Saturday":  "Sabtu",
}

var IndonesiaMonth = map[string]string{
	"January":   "Januari",
	"February":  "Februari",
	"March":     "Maret",
	"April":     "April",
	"May":       "Mei",
	"June":      "Juni",
	"July":      "Juli",
	"August":    "Agustus",
	"September": "September",
	"October":   "Oktober",
	"November":  "November",
	"December":  "Desember",
}

func ConvertToIndonesiaDateFormat(time *time.Time) string {
	day := time.Day()
	month := time.Month()
	year := time.Year()

	return fmt.Sprintf("%d %s %d", day, IndonesiaMonth[month.String()], year)
}

func ConvertToWIBHourTime(timpestamp *time.Time) string {
	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		fmt.Println("Failed to load location:", err)
	}

	wibTime := timpestamp.In(location)

	timeFormat := wibTime.Format("15:04")

	result := fmt.Sprintf("%s WIB", timeFormat)

	return result
}

func GetCompleteFormatted(time *time.Time) string {
	date := DateFormaterWithSeparator(time, " ")
	hour := ConvertToWIBHourTime(time)

	result := fmt.Sprintf("%s - %s", date, hour)

	return result
}

func ConvertToWeekdayAndDateFormat(time *time.Time) string {
	weekday := IndonesiaDate[time.Weekday().String()]
	day := time.Day()
	month := IndonesiaMonth[time.Month().String()][:3]
	year := time.Year()

	result := fmt.Sprintf("%s, %d-%s-%d", weekday, day, month, year)

	return result
}

func DateFormaterWithSeparator(time *time.Time, separator string) string {
	weekday := IndonesiaDate[time.Weekday().String()][:3]
	day := time.Day()
	month := IndonesiaMonth[time.Month().String()][:3]
	year := time.Year()

	result := fmt.Sprintf("%s, %d%s%s%s%d", weekday, day, separator, month, separator, year)

	return result
}
