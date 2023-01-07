package curupd

import(
	"testing"
	"fmt"
)

func TestGetRates(t *testing.T) {
	rates, err := GetCurrencyRates()
	if err != nil {
		panic(fmt.Sprintf("GetCurrencyRates() failed: %v",err))
	}
	
	fmt.Println(rates)
}

