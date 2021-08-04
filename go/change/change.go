package change

import (
	"errors"
	"sort"
)

func Change(coins []int, target int) ([]int, error) {
	if target == 0 {
		return make([]int, 0), nil
	}
	if target < 0 {
		return nil, errors.New("Target can't be less than zero")
	}

	var leastNumOfCoinChangeCombinations []int
	for _, coin := range coins {
		remainder := target - coin
		remainderCombination, _ := Change(coins, remainder)
		if remainderCombination != nil {
			combination := append(remainderCombination, coin)
			if leastNumOfCoinChangeCombinations == nil || len(combination) < len(leastNumOfCoinChangeCombinations) {
				leastNumOfCoinChangeCombinations = combination
			}
		}
	}
	sort.Ints(leastNumOfCoinChangeCombinations)
	return leastNumOfCoinChangeCombinations, nil
}
