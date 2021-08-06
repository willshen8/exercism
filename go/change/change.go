package change

import (
	"errors"
	"sort"
)

func Change(coins []int, target int, resultsCache map[int][]int) ([]int, error) {
	if val, ok := resultsCache[target]; ok {
		return val, nil
	}
	if target == 0 {
		return make([]int, 0), nil
	}
	if target < 0 {
		return nil, errors.New("Target can't be less than zero")
	}

	var leastNumOfCoinChangeCombinations []int
	for _, coin := range coins {
		remainder := target - coin
		remainderCombination, _ := Change(coins, remainder, resultsCache)

		if remainderCombination != nil {
			combination := append(remainderCombination, coin)
			if leastNumOfCoinChangeCombinations == nil || len(combination) < len(leastNumOfCoinChangeCombinations) {
				leastNumOfCoinChangeCombinations = combination
			}
		}
	}
	if leastNumOfCoinChangeCombinations == nil {
		return nil, errors.New("Can't find changes from coin combinations")
	}
	sort.Ints(leastNumOfCoinChangeCombinations)
	resultsCache[target] = leastNumOfCoinChangeCombinations
	return leastNumOfCoinChangeCombinations, nil
}
