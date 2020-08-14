package poker

import (
	"errors"
	"reflect"
	"sort"
	"strings"
)

var (
	CardNumbers = map[string]int{"2": 2, "3": 3, "4": 4,
		"5": 5, "6": 6, "7": 7, "8": 8, "9": 9, "10": 10,
		"J": 11, "Q": 12, "K": 13, "A": 14}
	CardSuit       = []string{"♢", "♡", "♧", "♤"}
	ErrInvalidHand = errors.New("Error, invalid cards dealt")
)

type ranking int

const (
	FiveOfAKind ranking = iota + 1
	StraightFlush
	FourOfAKind
	FullHouse
	Flush
	Straight
	ThreeOfAKind
	TwoPair
	OnePair
	HighHand
)

type rankedHand struct {
	hand    string
	numbers []int
	rank    ranking
}

func BestHand(hands []string) ([]string, error) {
	var result []string
	var rankedHands []rankedHand

	if len(hands) == 1 {
		_, err := DetermineHandRanking(hands[0])
		if err != nil {
			return result, err
		}
		return hands, nil
	}

	for _, v := range hands {
		rankedHand, err := DetermineHandRanking(v)
		if err != nil {
			return result, err
		} else {
			rankedHands = append(rankedHands, rankedHand)
		}
	}

	highestRank := rankedHands[0].rank
	for _, v := range rankedHands {
		if v.rank < highestRank {
			highestRank = v.rank
		}
	}

	var highestRankedHands []rankedHand
	for _, v := range rankedHands {
		if v.rank == highestRank {
			highestRankedHands = append(highestRankedHands, v)
		}
	}

	if len(highestRankedHands) > 1 {
		// Process full house
		if highestRankedHands[0].rank == FullHouse {
			var highestTriplet int
			var highestPair int
			for _, hand := range highestRankedHands {
				var numFrequencyMap = make(map[int]int, 2)
				for _, v := range hand.numbers {
					if _, ok := numFrequencyMap[v]; !ok {
						numFrequencyMap[v] = 1
					} else {
						numFrequencyMap[v]++
					}
				}
				for k, v := range numFrequencyMap {
					if v == 3 && k > highestTriplet {
						highestTriplet = k
					} else if v == 2 && k > highestPair {
						highestPair = k
					}
				}
			}

			var deleteFullHouseIndices []int
			for index, hand := range highestRankedHands {
				var numFrequencyMap = make(map[int]int, 2)
				for _, v := range hand.numbers {
					if _, ok := numFrequencyMap[v]; !ok {
						numFrequencyMap[v] = 1
					} else {
						numFrequencyMap[v]++
					}
				}
				for k, v := range numFrequencyMap {
					if v == 3 && k < highestTriplet {
						deleteFullHouseIndices = append(deleteFullHouseIndices, index)
					}
				}
			}

			for len(deleteFullHouseIndices) > 0 {
				highestRankedHands = append(highestRankedHands[:deleteFullHouseIndices[0]], highestRankedHands[deleteFullHouseIndices[0]+1:]...)
				deleteFullHouseIndices = append(deleteFullHouseIndices[:0], deleteFullHouseIndices[1:]...)
				for i := range deleteFullHouseIndices {
					deleteFullHouseIndices[i]--
				}
			}

			if len(highestRankedHands) > 1 {
				var deletePairIndices []int
				for index, hand := range highestRankedHands {
					var numFrequencyMap = make(map[int]int, 2)
					for _, v := range hand.numbers {
						if _, ok := numFrequencyMap[v]; !ok {
							numFrequencyMap[v] = 1
						} else {
							numFrequencyMap[v]++
						}
					}
					for k, v := range numFrequencyMap {
						if v == 2 && k < highestPair {
							deletePairIndices = append(deletePairIndices, index)
						}
					}
				}

				for len(deletePairIndices) > 0 {
					highestRankedHands = append(highestRankedHands[:deletePairIndices[0]], highestRankedHands[deletePairIndices[0]+1:]...)
					deletePairIndices = append(deletePairIndices[:0], deletePairIndices[1:]...)
					for i := range deletePairIndices {
						deletePairIndices[i]--
					}
				}
			}
			for _, v := range highestRankedHands {
				result = append(result, v.hand)
			}
			return result, nil
		} else if highestRankedHands[0].rank == FourOfAKind {
			var highestQuadruplet int
			for _, hand := range highestRankedHands {
				var numFrequencyMap = make(map[int]int, 1)
				for _, v := range hand.numbers {
					if _, ok := numFrequencyMap[v]; !ok {
						numFrequencyMap[v] = 1
					} else {
						numFrequencyMap[v]++
					}
				}
				for k, v := range numFrequencyMap {
					if v == 4 && k > highestQuadruplet {
						highestQuadruplet = k
					}
				}
			}

			var deleteFourOfAKindIndices []int
			for index, hand := range highestRankedHands {
				var numFrequencyMap = make(map[int]int, 2)
				for _, v := range hand.numbers {
					if _, ok := numFrequencyMap[v]; !ok {
						numFrequencyMap[v] = 1
					} else {
						numFrequencyMap[v]++
					}
				}
				for k, v := range numFrequencyMap {
					if v == 4 && k < highestQuadruplet {
						deleteFourOfAKindIndices = append(deleteFourOfAKindIndices, index)
					}
				}
			}

			for len(deleteFourOfAKindIndices) > 0 {
				highestRankedHands = append(highestRankedHands[:deleteFourOfAKindIndices[0]], highestRankedHands[deleteFourOfAKindIndices[0]+1:]...)
				deleteFourOfAKindIndices = append(deleteFourOfAKindIndices[:0], deleteFourOfAKindIndices[1:]...)
				for i := range deleteFourOfAKindIndices {
					deleteFourOfAKindIndices[i]--
				}
			}
			if len(highestRankedHands) > 1 {
				var highestKicker int

				for _, hand := range highestRankedHands {
					var numFrequencyMap = make(map[int]int, 2)
					for _, v := range hand.numbers {
						if _, ok := numFrequencyMap[v]; !ok {
							numFrequencyMap[v] = 1
						} else {
							numFrequencyMap[v]++
						}
					}
					for k, v := range numFrequencyMap {
						if v == 1 && k > highestKicker {
							highestKicker = k
						}
					}

				}
				var deleteKickerIndices []int
				for index, hand := range highestRankedHands {
					var numFrequencyMap = make(map[int]int, 2)
					for _, v := range hand.numbers {
						if _, ok := numFrequencyMap[v]; !ok {
							numFrequencyMap[v] = 1
						} else {
							numFrequencyMap[v]++
						}
					}
					for k, v := range numFrequencyMap {
						if v == 1 && k < highestKicker {
							deleteKickerIndices = append(deleteKickerIndices, index)
						}
					}
				}

				for len(deleteKickerIndices) > 0 {
					highestRankedHands = append(highestRankedHands[:deleteKickerIndices[0]], highestRankedHands[deleteKickerIndices[0]+1:]...)
					deleteKickerIndices = append(deleteKickerIndices[:0], deleteKickerIndices[1:]...)
					for i := range deleteKickerIndices {
						deleteKickerIndices[i]--
					}
				}
			}

			for _, v := range highestRankedHands {
				result = append(result, v.hand)
			}
			return result, nil
		}

		for i := 0; i < 5; i++ {
			if len(highestRankedHands) == 1 {
				break
			}
			highestCardNum := highestRankedHands[0].numbers[i]
			for j := 0; j < len(highestRankedHands); j++ {
				if highestRankedHands[j].numbers[i] > highestCardNum {
					highestCardNum = highestRankedHands[j].numbers[i]
				}
			}

			var deleteIndices []int
			for index, card := range highestRankedHands {
				if card.numbers[i] != highestCardNum {
					deleteIndices = append(deleteIndices, index)
				}
			}

			for len(deleteIndices) > 0 {
				highestRankedHands = append(highestRankedHands[:deleteIndices[0]], highestRankedHands[deleteIndices[0]+1:]...)
				deleteIndices = append(deleteIndices[:0], deleteIndices[1:]...)
				for i := range deleteIndices {
					deleteIndices[i]--
				}
			}
		}
	}

	for _, v := range highestRankedHands {
		result = append(result, v.hand)
	}

	return result, nil
}

// DetermineHandRanking parses and validate the cards and return the ranking and error
// if exists
func DetermineHandRanking(hand string) (rankedHand, error) {
	parsedCards := strings.Split(hand, " ")
	result := rankedHand{hand: hand}
	if len(parsedCards) != 5 {
		return result, ErrInvalidHand
	}
	var numbers []int
	var suits []string

	for _, cards := range parsedCards {
		card := strings.Split(cards, "")
		var number, suit string
		if len(card) == 2 {
			number, suit = card[0], card[1]
		} else if len(card) == 3 {
			number = card[0] + card[1]
			suit = card[2]
		} else {
			return result, ErrInvalidHand
		}

		if _, ok := CardNumbers[number]; !ok {
			return result, ErrInvalidHand
		}
		if !Contains(suit, CardSuit) {
			return result, ErrInvalidHand
		}

		numbers = append(numbers, CardNumbers[number])
		suits = append(suits, suit)

	}
	var numbFrequencyMap = make(map[int]int, 5)
	for _, v := range numbers {
		if _, ok := numbFrequencyMap[v]; !ok {
			numbFrequencyMap[v] = 1
		} else {
			numbFrequencyMap[v]++
		}
	}
	// get the values of the numbersMap
	var cardFrequency []int
	for _, v := range numbFrequencyMap {
		cardFrequency = append(cardFrequency, v)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(cardFrequency)))
	sort.Sort(sort.Reverse(sort.IntSlice(numbers)))
	isStraightFlag := IsStraight(numbers)
	isFlushFlag := IsFlush(suits)
	// if A exists in a low straight, then treat is as a lower card
	if isStraightFlag && reflect.DeepEqual(numbers, []int{14, 5, 4, 3, 2}) {
		numbers = append(numbers[:0], []int{5, 4, 3, 2, 1}...)
	}

	result.numbers = numbers

	if isFlushFlag && isStraightFlag {
		result.rank = StraightFlush
		return result, nil
	} else if isFlushFlag {
		result.rank = Flush
		return result, nil
	} else if isStraightFlag {
		result.rank = Straight
		return result, nil
	}

	var threeOfAKind bool = false
	var twoPairs int
	for _, v := range cardFrequency {
		if v == 4 {
			result.rank = FourOfAKind
			return result, nil
		} else if v == 3 {
			threeOfAKind = true
		} else if v == 2 {
			twoPairs++
		}
	}
	if threeOfAKind && twoPairs == 1 {
		result.rank = FullHouse
		return result, nil
	} else if threeOfAKind {
		result.rank = ThreeOfAKind
		return result, nil
	} else if twoPairs == 2 {
		result.rank = TwoPair
		return result, nil
	} else if twoPairs == 1 {
		result.rank = OnePair
		return result, nil
	}

	result.rank = HighHand
	return result, nil
}

// Contains return a bool if s exists in collection slice
func Contains(s string, collection []string) bool {
	for _, v := range collection {
		if v == s {
			return true
		}
	}
	return false
}

func IsStraight(hand []int) bool {
	if hand[0] == 14 && hand[1] == 5 && hand[2] == 4 && hand[3] == 3 && hand[4] == 2 {
		return true
	}

	for i := range hand {
		if i != 0 && hand[i]+1 != hand[i-1] {
			return false
		}
	}
	return true
}

func IsFlush(hand []string) bool {
	for i := range hand {
		if i != 0 && hand[i] != hand[i-1] {
			return false
		}
	}
	return true
}
