package coinchange

import (
	"testing"
)

func TestCoinChange(t *testing.T) {

	var coins = []int{3, 5, 1}
	var amount = 11
	var want = 3
	got := coinChange(coins, amount)
	if got != want {
		t.Fatalf(`coinChange func failed: want %d, got %d`, want, got)
	}

}
