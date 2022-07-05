package utils

import (
	"regexp"
	"strings"
	"time"
)

func ExtractString(str string) string {
	start, end := 0, len(str)
	if strings.HasPrefix(str, `"`) {
		start += 1
	}
	if strings.HasSuffix(str, `"`) {
		end -= 1
	}
	return str[start:end]
}

// date -d "2022-06-30T08:44:58Z" +"%Y-%m-%d %H:%M:%S"
func TransTime(tzTime string) (string, error) {
	parTime, err := time.Parse("2006-01-02T15:04:05Z", tzTime)
	if err != nil {
		// log.Warnf("Time parse failed with %s", err)
		return tzTime, err
	}
	zone, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		// log.Warnf("Time load location failed with %s", err)
		return tzTime, err
	}
	ttTime := parTime.In(zone).Format("2006-01-02 15:04:05 MST")
	return ttTime, nil
}

func CompressSpace(str string) string {
	if str == "" {
		return ""
	}
	reg := regexp.MustCompile("\\s+")
	return reg.ReplaceAllString(str, " ")
}
