Получение списка валют с сайта [www.cbr.ru](www.cbr.ru "Центральный банк России")
============================================================

Получает актуальные курсы валют с сайта [www.cbr.ru](www.cbr.ru "Центральный банк России"), возвращает заполненную структуру.
```go
	import (
		"fmt"
		"github.com/dronm/curupd"
	)

	rates, err := curupd.GetCurrencyRates()
	if err != nil {
		panic(fmt.Sprintf("GetCurrencyRates() failed: %v",err))
	}
	
	fmt.Println(rates)

```
