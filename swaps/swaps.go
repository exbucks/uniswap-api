package swaps

import (
	uniswap "github.com/hirokimoto/uniswap-api"
	"github.com/hirokimoto/uniswap-api/swap"
)

// WholePriceChanges returns how much has been changed in the whole swaps.
func WholePriceChanges(swaps uniswap.Swaps) (price float64, change float64) {
	if swaps.Data.Swaps != nil && len(swaps.Data.Swaps) > 0 {
		first, _ := swap.Price(swaps.Data.Swaps[0])
		last, _ := swap.Price(swaps.Data.Swaps[len(swaps.Data.Swaps)-1])
		change = first - last
		return first, change
	}
	return 0.0, 0.0
}

// LastPriceChanges returns the amount of price changes of the last 2 swaps.
func LastPriceChanges(swaps uniswap.Swaps) (float64, float64) {
	if swaps.Data.Swaps != nil && len(swaps.Data.Swaps) > 0 {
		first, _ := swap.Price(swaps.Data.Swaps[0])
		second, _ := swap.Price(swaps.Data.Swaps[len(swaps.Data.Swaps)-1])
		change := first - second
		return first, change
	}
	return 0.0, 0.0
}
