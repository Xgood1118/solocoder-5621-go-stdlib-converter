package internal

type CurrencyConfig struct {
	Rates map[string]float64
}

var currencyRates = map[string]float64{
	"USD": 1,
	"CNY": 7.25,
	"EUR": 0.92,
	"JPY": 157.5,
	"GBP": 0.79,
}

var currencyNames = map[string][]string{
	"CNY": {"人民币", "CNY", "RMB", "yuan"},
	"USD": {"美元", "USD", "dollar"},
	"EUR": {"欧元", "EUR", "euro"},
	"JPY": {"日元", "JPY", "yen"},
	"GBP": {"英镑", "GBP", "pound"},
}

func init() {
	for code, rate := range currencyRates {
		RegisterUnit(&Unit{
			Key:      code,
			Category: CatCurrency,
			Names:    currencyNames[code],
			Factor:   NewFactor(rate),
		})
	}
}

func LoadCurrencyRates(rates map[string]float64) {
	for code, rate := range rates {
		currencyRates[code] = rate
	}
	for code, rate := range currencyRates {
		RegisterUnit(&Unit{
			Key:      code,
			Category: CatCurrency,
			Names:    currencyNames[code],
			Factor:   NewFactor(rate),
		})
	}
}

func SetCurrencyRate(code string, rate float64) {
	currencyRates[code] = rate
	RegisterUnit(&Unit{
		Key:      code,
		Category: CatCurrency,
		Names:    currencyNames[code],
		Factor:   NewFactor(rate),
	})
}

func GetCurrencyRate(code string) (float64, bool) {
	rate, ok := currencyRates[code]
	return rate, ok
}

func IsCurrencyUnit(name string) bool {
	u, ok := LookupUnit(name)
	if !ok {
		return false
	}
	return u.Category == CatCurrency
}
