package until

import (
	"strconv"
	"strings"
)

func ParseFileName(fileName string) (minMonth, maxMonth int64, err error) {
	length := len(fileName)
	if length > 10 {
		// 多月表
		stringArr := strings.Split(fileName, "_")
		minMonth, err = strconv.ParseInt(stringArr[2], 10, 64)
		if err != nil {
			return 0, 0, err
		}
		maxMonth, err = strconv.ParseInt(stringArr[4], 10, 64)
		if err != nil {
			return 0, 0, err
		}
		return
	} else {
		// 单月表
		stringMonth := strings.Split(fileName, "_")
		minMonth = 0
		maxMonth, err = strconv.ParseInt(stringMonth[2], 10, 64)
		if err != nil {
			return 0, 0, err
		}
		return
	}
}
