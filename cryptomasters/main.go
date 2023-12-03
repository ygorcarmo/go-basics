package main

import (
	"example/crypto/api"
	"fmt"
	"sync"
)

func main() {
	currencies := []string{"btc", "eth", "bch"}

	var wg sync.WaitGroup

	for _, currency := range currencies {
		wg.Add(1)
		go func(currencyCode string) {
			getCurrencyData(currencyCode)
			wg.Done()
		}(currency)
	}
	wg.Wait()
}

func getCurrencyData(currency string) {
	rate, err := api.GetRate(currency)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	fmt.Printf("The rate for %v is %.2f \n", rate.Currency, rate.Price)

}
