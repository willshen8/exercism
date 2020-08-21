package poker

import (
	"errors"
	"sort"
	"strings"
)

var (
	ErrInvalidHand = errors.New("Error, invalid cards dealt")
	CardRank       = map[string]int{"2": 2, "3": 3, "4": 4,
		"5": 5, "6": 6, "7": 7, "8": 8, "9": 9, "10": 10,
		"J": 11, "Q": 12, "K": 13, "A": 14}
)

type Card struct {
	rank int
	suit string
}

type Hand struct {
	rank  int
	cards []Card
}

const (
	Heart   = '♡'
	Diamond = '♢'
	Spade   = '♧'
	Club    = '♤'
)
const (
	HighHand = iota + 1
	OnePair
	TwoPair
	ThreeOfAKind
	Straight
	Flush
	FullHouse
	FourOfAKind
	StraightFlush
)

func BestHand(hands []string) ([]string, error) {
	var results []string
	var bestHand Hand
	for _, hand := range hands {
		h, err := parseHand(hand)
		if err != nil {
			return results, err
		}
		if h.IsBetterHand(bestHand) {
			bestHand, results = h, []string{hand}
			continue
		}
		if bestHand.IsBetterHand(h) {
			continue
		}
		results = append(results, hand)
	}
	return results, nil
}

func (h *Hand) IsBetterHand(otherHand Hand) bool {
	if h.rank != otherHand.rank {
		return h.rank > otherHand.rank
	}
	if h.rank == StraightFlush {
		if h.cards[0].rank == 2 && h.cards[4].rank == 14 {
			return false
		}
	}

	for i, v := range h.cards {
		if v.rank != otherHand.cards[i].rank {
			return v.rank > otherHand.cards[i].rank
		}
	}
	return false
}

func parseHand(hand string) (Hand, error) {
	var parsedHand Hand
	parsedCards := strings.Split(hand, " ")
	if len(parsedCards) != 5 {
		return parsedHand, ErrInvalidHand
	}
	for _, cardStr := range parsedCards {
		processedCard, err := parseCard(cardStr)
		if err != nil {
			return parsedHand, err
		}
		parsedHand.cards = append(parsedHand.cards, processedCard)
	}
	parsedHand.determineRank()
	return parsedHand, nil
}

func parseCard(cardStr string) (Card, error) {
	var c Card
	card := strings.Split(cardStr, "")
	var rank string
	if len(card) == 2 {
		rank = card[0]
		c.suit = card[1]
	} else if len(card) == 3 {
		rank = card[0] + card[1]
		c.suit = card[2]
	} else {
		return c, ErrInvalidHand
	}
	if _, ok := CardRank[rank]; !ok {
		return c, ErrInvalidHand
	}
	switch c.suit {
	case string(Heart), string(Diamond), string(Spade), string(Club):
		break
	default:
		return c, ErrInvalidHand
	}
	c.rank = CardRank[rank]
	return c, nil
}

func (h *Hand) determineRank() {
	var rankFrequencyMap = make(map[int]int, 5)
	for _, card := range h.cards {
		if _, ok := rankFrequencyMap[card.rank]; !ok {
			rankFrequencyMap[card.rank] = 1
		} else {
			rankFrequencyMap[card.rank]++
		}
	}

	sort.Slice(h.cards, func(i, j int) bool {
		if rankFrequencyMap[h.cards[i].rank] > rankFrequencyMap[h.cards[j].rank] {
			return true
		}
		if rankFrequencyMap[h.cards[i].rank] == rankFrequencyMap[h.cards[j].rank] {
			return h.cards[i].rank > h.cards[j].rank
		}
		return false
	})

	var isStraightFlag, isFlushFlag, threeOfAKindFlag bool
	var twoPairs int
	isStraightFlag = IsStraight(*h)
	isFlushFlag = IsFlush(*h)

	if isFlushFlag && isStraightFlag {
		h.rank = StraightFlush
		return
	} else if isFlushFlag {
		h.rank = Flush
		return
	} else if isStraightFlag {
		h.rank = Straight
		return
	}

	for _, v := range rankFrequencyMap {
		if v == 4 {
			h.rank = FourOfAKind
			return
		} else if v == 3 {
			threeOfAKindFlag = true
		} else if v == 2 {
			twoPairs++
		}
	}

	if threeOfAKindFlag && twoPairs == 1 {
		h.rank = FullHouse
	} else if threeOfAKindFlag {
		h.rank = ThreeOfAKind
	} else if twoPairs == 2 {
		h.rank = TwoPair
	} else if twoPairs == 1 {
		h.rank = OnePair
	} else {
		h.rank = HighHand
	}
}

func IsStraight(h Hand) bool {
	// special case where A is treated as a 1
	if h.cards[0].rank == 14 && h.cards[1].rank == 5 &&
		h.cards[2].rank == 4 && h.cards[3].rank == 3 &&
		h.cards[4].rank == 2 {
		h.cards[0].rank = 1
		return true
	}

	for i := 1; i < len(h.cards); i++ {
		if h.cards[i-1].rank != h.cards[i].rank+1 {
			return false
		}
	}
	return true
}

func IsFlush(h Hand) bool {
	for i := range h.cards {
		if i != 0 && h.cards[i].suit != h.cards[i-1].suit {
			return false
		}
	}
	return true
}
