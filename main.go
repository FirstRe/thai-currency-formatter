package main

import (
	"fmt"
	"sync"

	currency "github.com/FirstRe/thai-currency-formatter/currency"
	"github.com/shopspring/decimal"
)

func main() {
	inputs := []decimal.Decimal{
		decimal.NewFromFloat(1234),
		decimal.NewFromFloat(33333.75),
	}
	var wg sync.WaitGroup

	ch := make(chan string, len(inputs))

	for _, input := range inputs {
		wg.Add(1)
		go func(input decimal.Decimal) {
			defer wg.Done()
			result := currency.FormatAmount(input)
			ch <- result
		}(input)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for output := range ch {
		fmt.Println(output)
		fmt.Println("-------------")
	}
}
