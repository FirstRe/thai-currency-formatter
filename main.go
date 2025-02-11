package main

import (
	"fmt"

	currency "github.com/FirstRe/thai-currency-formatter/currency"
	"github.com/shopspring/decimal"
)

func main() {
	inputs := []decimal.Decimal{
		decimal.NewFromFloat(1234),
		decimal.NewFromFloat(33333.75),
	}

	for _, input := range inputs {
		fmt.Println(currency.FormatAmount(input))
		fmt.Println("-------------")
	}

}
