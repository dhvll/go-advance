package util

const (
	USD = "USD"
	EUR = "EUR"
	INR = "INR"
	BTC = "BTC"
)

// IsSupportedCurrency returns true if the currency is supported
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, INR, BTC:
		return true
	}
	return false
}
