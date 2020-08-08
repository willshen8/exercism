package poker

import (
	"errors"
	"strings"
)

var (
	CardNumbers    = []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}
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

type hand struct {
	hand string
	rank ranking
}

func BestHand(hands []string) ([]string, error) {
	if len(hands) == 1 {
		return hands, nil
	}
	var results []string
	return results, nil
}

// ValidateHand parses and validate the cards being dealt
func ValidateHand(hand string) error {
	parsedCards := strings.Split(hand, " ")
	if len(parsedCards) != 5 {
		return ErrInvalidHand
	}
	cardsMap := make(map[string][]string, len(parsedCards))
	for _, cards := range parsedCards {
		card := strings.Split(cards, "")
		var number, suit string
		if len(card) == 2 {
			number, suit = card[0], card[1]
		} else {
			number = card[0] + card[1]
			suit = card[2]
		}
		if !Contains(number, CardNumbers) || !Contains(suit, CardSuit) {
			return ErrInvalidHand
		}
		if _, ok := cardsMap[number]; !ok {
			newNumber := []string{suit}
			cardsMap[number] = newNumber
		} else {
			cardsMap[number] = append(cardsMap[number], suit)
		}

	}
	return nil
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

func IsStraight(hand string) bool {
	return true
}

func IsFlush(hand map[string][]string) bool {
	return false
}

func IsFullHouse(hand string) bool {
	return true
}
