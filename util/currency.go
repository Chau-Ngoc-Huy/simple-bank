package util

//const (
//	USD = "USD"
//	VND = "VND"
//	EUR = "EUR"
//)

var SupportCurrencies = []string{
	"USD",
	"VND",
	"EUR",
}

func IsSupportCurrency(currency string) bool {
	for i := range SupportCurrencies {
		if currency == SupportCurrencies[i] {
			return true
		}
	}
	return false
}
