package poker

import (
	"go-poker/ranks"
	"log"
	"strings"
)

var cardRanks = ranks.CardVals
var rankWeights = map[string]int{"four-of-a-kind": 6, "full-house": 5, "triples": 4, "two-pairs": 3, "one-pair": 2, "high-card": 1}

//HandCompare accepts two hands - checks if valid and compares ranks to get winner hand.
func HandCompare(hand1 string, hand2 string) (bool, bool) {

	winHand1, windHand2 := false, false

	if IsValid(hand1) && IsValid(hand2) {

		rank1, cardsOut1 := GetRank(hand1)
		rank2, cardsOut2 := GetRank(hand2)

		if rankWeights[rank1] > rankWeights[rank2] {
			winHand1, windHand2 = true, false
		} else if rankWeights[rank1] < rankWeights[rank2] {
			winHand1, windHand2 = false, true
		} else {
			for k, i := range cardsOut1 {
				if cardRanks[i] > cardRanks[cardsOut2[k]] {
					winHand1, windHand2 = true, false
					return winHand1, windHand2
				} else if cardRanks[i] < cardRanks[cardsOut2[k]] {
					winHand1, windHand2 = false, true
					return winHand1, windHand2
				}
			}
			winHand1, windHand2 = true, true
		}
	}
	return winHand1, windHand2
}

//IsValid checks if hand is valid - contains acceptable letters and of allowed length.
func IsValid(hand string) bool {

	handChars := strings.Split(hand, "")
	if len(handChars) != 5 {
		return false
	}

	for _, s := range handChars {
		if _, found := cardRanks[s]; !found {
			return false
		}
	}

	return true
}

//GetRank executes ranking functions in descending order of weight and returns the matching rank and cards order.
func GetRank(hand string) (string, []string) {

	var rank string
	var cardsOut []string
	handMap := make(map[string]int)
	handChars := strings.Split(hand, "")

	for _, s := range handChars {
		if _, found := handMap[s]; found {
			handMap[s] = handMap[s] + 1
		} else {
			handMap[s] = 1
		}
	}

	rankFuncs := []func(map[string]int) (string, []string){ranks.FourOfaKind, ranks.FullHouse, ranks.Triple,
		ranks.TwoPairs, ranks.OnePair, ranks.HighCard}

	for _, f := range rankFuncs {
		rank, cardsOut = f(handMap)
		if len(rank) > 0 {
			break
		}
	}

	log.Printf("Rank for hand %s is : %s and cards ordered as : %q \n", hand, rank, cardsOut)
	return rank, cardsOut
}
