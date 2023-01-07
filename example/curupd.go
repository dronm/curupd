package main

//prints out to stdout currency structure

import(
	"encoding/json"
	"fmt"	
	
	"github.com/dronm/curupd"
)

func main() {
	rates, err := curupd.GetCurrencyRates()
	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}
	
	b, err := json.Marshal(rates)
	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}
	fmt.Println(string(b))
}

