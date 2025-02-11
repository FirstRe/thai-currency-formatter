package currency

import (
	"strings"

	"github.com/shopspring/decimal"
)

var textThai = map[int64]string{
	1:       "หนึ่ง",
	2:       "สอง",
	3:       "สาม",
	4:       "สี่",
	5:       "ห้า",
	6:       "หก",
	7:       "เจ็ด",
	8:       "แปด",
	9:       "เก้า",
	10:      "สิบ",
	100:     "ร้อย",
	1000:    "พัน",
	10000:   "หมื่น",
	100000:  "แสน",
	1000000: "ล้าน",
}

func NumberToThai(num int64) string {
	keys := []int64{100000, 10000, 1000, 100, 10, 1}
	if num <= 0 {
		return ""
	}

	if num < 10 {
		return textThai[num]
	}

	if num >= 1000000 {
		mil := num / 1000000
		rest := num % 1000000

		return NumberToThai(mil) + "ล้าน" + NumberToThai(rest)
	}

	result := ""

	for _, key := range keys {
		if num < key {
			continue
		}

		div := num / key

		if num == 0 {
			continue
		}
		if key == 10 {
			if div == 2 {
				result += "ยี่สิบ"
			} else if div == 1 {
				result += "สิบ"
			} else {
				result += textThai[div] + textThai[key]
			}
		} else if key == 1 && div == 1 && strings.HasSuffix(result, "สิบ") {
			result += "เอ็ด"
		} else if key < 10 {
			result += textThai[num]
		} else {
			result += textThai[div] + textThai[key]
		}
		num %= key
	}
	return result
}

func FormatAmount(input decimal.Decimal) string {

	if input.IsNegative() {
		return "จำนวนควรมากกว่าหรือเท่ากับ 0"
	}

	intValue := input.IntPart()

	decimalValue := input.Sub(decimal.NewFromInt(intValue)).Mul(decimal.NewFromInt(100)).IntPart()
	result := ""

	if intValue == 0 {
		result += "ศูนย์บาท"
	} else {
		result = NumberToThai(intValue) + "บาท"
	}

	if decimalValue > 0 {
		result += NumberToThai(decimalValue) + "สตางค์"
	} else {
		result += "ถ้วน"
	}

	return result
}
