package testing

import (
	"testing"

	currency "github.com/FirstRe/thai-currency-formatter/currency"
	"github.com/shopspring/decimal"
)

func TestNumberToThai(t *testing.T) {
	tests := []struct {
		input    int64
		expected string
	}{
		{1, "หนึ่ง"},
		{-1, ""},
		{9, "เก้า"},
		{10, "สิบ"},
		{11, "สิบเอ็ด"},
		{21, "ยี่สิบเอ็ด"},
		{99, "เก้าสิบเก้า"},
		{100, "หนึ่งร้อย"},
		{105, "หนึ่งร้อยห้า"},
		{999, "เก้าร้อยเก้าสิบเก้า"},
		{1000, "หนึ่งพัน"},
		{1000000, "หนึ่งล้าน"},
		{2000000, "สองล้าน"},
		{1000000000, "หนึ่งพันล้าน"},
		{10000000000, "หนึ่งหมื่นล้าน"},
		{100000000000, "หนึ่งแสนล้าน"},
		{1000001000001, "หนึ่งล้านหนึ่งล้านหนึ่ง"},
	}
	for _, test := range tests {
		result := currency.NumberToThai(test.input)
		if result != test.expected {
			t.Errorf("Expected currency.NumberToThai(%d) to be \"%s\", but got \"%s\"", test.input, test.expected, result)
		}
	}
}

func TestFormatAmount(t *testing.T) {
	tests := []struct {
		input    decimal.Decimal
		expected string
	}{
		{decimal.NewFromFloat(0), "ศูนย์บาทถ้วน"},
		{decimal.NewFromFloat(1), "หนึ่งบาทถ้วน"},
		{decimal.NewFromFloat(99.99), "เก้าสิบเก้าบาทเก้าสิบเก้าสตางค์"},
		{decimal.NewFromFloat(9.99), "เก้าบาทเก้าสิบเก้าสตางค์"},
		{decimal.NewFromFloat(11.25), "สิบเอ็ดบาทยี่สิบห้าสตางค์"},
		{decimal.NewFromFloat(21.50), "ยี่สิบเอ็ดบาทห้าสิบสตางค์"},
		{decimal.NewFromFloat(11), "สิบเอ็ดบาทถ้วน"},
		{decimal.NewFromFloat(1000000000.99), "หนึ่งพันล้านบาทเก้าสิบเก้าสตางค์"},
		{decimal.NewFromFloat(1000000.75), "หนึ่งล้านบาทเจ็ดสิบห้าสตางค์"},
		{decimal.NewFromFloat(2000000.00), "สองล้านบาทถ้วน"},
		{decimal.NewFromFloat(-5), "จำนวนควรมากกว่าหรือเท่ากับ 0"},
	}
	for _, test := range tests {
		result := currency.FormatAmount(test.input)
		if result != test.expected {
			t.Errorf("Expected currency.FormatAmount(%s) to be \"%s\", but got \"%s\"", test.input, test.expected, result)
		}
	}
}
