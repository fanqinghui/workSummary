package utils

import (
	"fmt"
	"strings"
)

func CheckErr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func Substr(str string, start, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}
	return string(rs[start:end])
}

func DoJsonData(body []byte) string {
	bodyStr := string(body)
	bodyStr = Substr(bodyStr, 1, strings.LastIndex(bodyStr, "'")-1)
	bodyStr = strings.Replace(bodyStr, "{", "{\"", 1)
	fmt.Println("body:" + bodyStr)
	bodyStr = strings.Replace(bodyStr, ":", "\":\"", 1)
	bodyStr = strings.Replace(bodyStr, "}", "\"}", 1)
	return bodyStr
}

func Byte2Str(data []byte) string {
	return string(data)
}

func GetMonthByStr(month string) int {
	intMonth := 0
	switch month {
	case "January Month":
		intMonth = 1
	case "February":
		intMonth = 2
	case "March":
		intMonth = 3
	case "April":
		intMonth = 4
	case "May":
		intMonth = 5
	case "June":
		intMonth = 6
	case "July":
		intMonth = 7
	case "August":
		intMonth = 8
	case "September":
		intMonth = 9
	case "October":
		intMonth = 10
	case "November":
		intMonth = 11
	case "December":
		intMonth = 12
	default:
		intMonth = 0
	}

	return intMonth
}
