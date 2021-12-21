package ptime

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func Format(src time.Time, format string) string {

	timezone, offset := src.Zone()
	offsetStr := strconv.FormatInt(int64(offset), 10)
	if offsetStr == "" {
		offsetStr = "0"
	}
	isoYear, isoWeek := src.ISOWeek()

	n := ""
	if src.Weekday() == 0 {
		n = "7"
	} else {
		n = strconv.FormatInt(int64(src.Weekday()), 10)
	}

	replacer := strings.NewReplacer(
		"d", "02", // Day of the month, 2 digits with leading zeros
		"j", "2", // Day of the month without leading zeros
		"m", "01", // Numeric representation of a month, with leading zeros
		"n", "1", // Numeric representation of a month, without leading zeros
		"Y", "2006", // A full numeric representation of a year, 4 digits
		"y", "06", // A two digit representation of a year
		"g", "3", // 12-hour format of an hour without leading zeros
		"h", "03", // 12-hour format of an hour with leading zeros
		"H", "15", // 24-hour format of an hour with leading zeros
		"i", "04", // Minutes with leading zeros
		"s", "05", // Seconds with leading zeros
		"u", "000000", // Microseconds
		"v", "000", // Milliseconds
		"O", "-0700", // Difference to Greenwich time (GMT) without colon between hours and minutes
		"P", "-07:00", // Difference to Greenwich time (GMT) with colon between hours and minutes
	)

	formatted := src.Format(replacer.Replace(format))

	replacerWords := strings.NewReplacer(
		"a", src.Format("pm"), // Lowercase Ante meridiem and Post meridiem
		"A", src.Format("PM"), // Uppercase Ante meridiem and Post meridiem
		"D", src.Format("Mon"), // A textual representation of a day, three letters
		"l", src.Format("Monday"), // A full textual representation of the day of the week
		"F", src.Format("January"), // A full textual representation of a month, such as January or March
		"M", src.Format("Jan"), // A short textual representation of a month, three letters
		"T", src.Format("MST"), // Timezone abbreviation, if known; otherwise the GMT offset.

		// The day of the year (starting from 0)
		"z", strconv.Itoa(src.YearDay()-1),
		// Numeric representation of the day of the week
		"w", strconv.Itoa(int(src.Weekday())),

		// ISO 8601 numeric representation of the day of the week
		"N", n,

		// English ordinal suffix for the day of the month, 2 characters st, nd, rd, th
		"S", ordinalSuffix(src.Day()),

		// ISO 8601 week number of year, weeks starting on Monday
		"W", fmt.Sprintf("%02d", isoWeek),

		// 24-hour format of an hour without leading zeros
		"G", strconv.Itoa(src.Hour()),

		// ISO 8601 week-numbering year. This has the same value as Y, except that if the ISO week number (W) belongs to the previous or next year, that year is used instead.
		"o", strconv.Itoa(isoYear),

		// Number of days in the given month
		"t", strconv.Itoa(daysInMonth(src)),

		// Whether it's a leap year
		"L", leapYear(src),

		// Seconds since the Unix Epoch (January 1 1970 00:00:00 GMT)
		"U", strconv.FormatInt(src.Unix(), 10),

		// Timezone identifier
		"e", timezone,
		// Timezone offset in seconds.
		"z", offsetStr,
	)
	return replacerWords.Replace(formatted)
}

func leapYear(src time.Time) string {
	if src.Year()%4 == 0 && src.Year()%100 != 0 || src.Year()%400 == 0 {
		return "1"
	}
	return "0"
}

func daysInMonth(src time.Time) int {
	return time.Date(src.Year(), src.Month()+1, 1, 0, 0, 0, -1, src.Location()).Day()
}

func ordinalSuffix(day int) string {
	switch day {
	case 1, 21, 31:
		return "st"
	case 2, 22:
		return "nd"
	case 3, 23:
		return "rd"
	default:
		return "th"
	}
}
