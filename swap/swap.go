package swap

import (
	"strconv"

	uniswap "github.com/hirokimoto/uniswap-api"
)

func SwapName(swap uniswap.Swap) (name string) {
	if swap.Pair.Token0.Symbol == "WETH" {
		name = swap.Pair.Token1.Name
	} else {
		name = swap.Pair.Token0.Name
	}
	return name
}

func SwapPrice(swap uniswap.Swap) (price float64, target string) {
	amountUSD, _ := strconv.ParseFloat(swap.AmountUSD, 32)
	amountToken, _ := strconv.ParseFloat(swap.Amount0Out, 32)

	if swap.Pair.Token0.Symbol == "WETH" {
		if swap.Amount0In != "0" && swap.Amount1Out != "0" {
			amountToken, _ = strconv.ParseFloat(swap.Amount1Out, 32)
			target = "BUY"
		} else if swap.Amount0Out != "0" && swap.Amount1In != "0" {
			amountToken, _ = strconv.ParseFloat(swap.Amount1In, 32)
			target = "SELL"
		}
	} else {
		if swap.Amount0Out != "0" && swap.Amount1In != "0" {
			amountToken, _ = strconv.ParseFloat(swap.Amount0Out, 32)
			target = "BUY"
		} else if swap.Amount0In != "0" && swap.Amount1Out != "0" {
			amountToken, _ = strconv.ParseFloat(swap.Amount0In, 32)
			target = "SELL"
		}
	}

	price = amountUSD / amountToken
	return price, target
}
