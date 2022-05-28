package coinchange

import (
	"testing"
)

func TestCoinChange(t *testing.T) {

	var coins = []int{186, 419, 83, 408}
	var amount = 6249
	var want = 13
	got := coinChange(coins, amount)
	if got != want {
		t.Fatalf(`coinChange func failed: want %d, got %d`, want, got)
	}

}
