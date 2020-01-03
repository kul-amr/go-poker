package ranks

import (
	"sort"
)

//CardVals stores value corresponding to the card.
var CardVals = map[string]int{"2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9, "T": 10, "J": 11, "Q": 12, "K": 13, "A": 14}

//FourOfaKind function checks if four cards of same value exists.
func FourOfaKind(handMap map[string]int) (string, []string) {
	var rank string
	var cardsOut []string

	if len(handMap) == 2 {
		for k, val := range handMap {
			if val == 4 {
				rank = "four-of-a-kind"
				cardsOut = append([]string{k, k, k, k}, cardsOut...)
			} else {
				cardsOut = append(cardsOut, k)
			}
		}
	}
	return rank, cardsOut
}

//FullHouse function checks if three cards of same value and the other two of similar value exists.
func FullHouse(handMap map[string]int) (string, []string) {

	var cardTriple []string
	var cardDouble []string
	var rank string
	var cardsOut []string

	if len(handMap) == 2 {
		for k, val := range handMap {
			if val == 3 {
				cardTriple = []string{k, k, k}
			} else if val == 2 {
				cardDouble = []string{k, k}
			}
		}

		if len(cardTriple) > 0 && len(cardDouble) > 0 {
			rank = "full-house"
			cardsOut = append(cardTriple, cardDouble...)
		}
	}
	return rank, cardsOut
}

//Triple function checks if three cards of same value and other two different values exists.
func Triple(handMap map[string]int) (string, []string) {

	var cardTriple bool
	var rank string
	var cardsOut []string
	weightMap := make(map[int]string)
	var weightArr []int

	for k, val := range handMap {
		if val == 3 {
			cardTriple = true
			cardsOut = append([]string{k, k, k}, cardsOut...)
		} else {
			weightMap[CardVals[k]] = k
			weightArr = append(weightArr, CardVals[k])
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(weightArr)))
	for _, i := range weightArr {
		cardsOut = append(cardsOut, weightMap[i])
	}

	if cardTriple {
		rank = "triples"
	}

	return rank, cardsOut
}

//TwoPairs function checks if two pairs of same value cards exists.
func TwoPairs(handMap map[string]int) (string, []string) {

	var numPairs int
	var rank string
	var cardsOut []string
	var tempSingleCard string

	if len(handMap) == 3 {
		for k, val := range handMap {
			if val == 2 {
				numPairs++
				if len(cardsOut) > 0 && CardVals[k] > CardVals[cardsOut[0]] {
					cardsOut = append([]string{k, k}, cardsOut...)
				} else {
					cardsOut = append(cardsOut, []string{k, k}...)
				}
			} else {
				tempSingleCard = k
			}
		}

		cardsOut = append(cardsOut, tempSingleCard)

		if numPairs == 2 {
			rank = "two-pairs"
		}
	}
	return rank, cardsOut
}

//OnePair function checks if one pair of same value card exists.
func OnePair(handMap map[string]int) (string, []string) {

	var numPairs int
	var rank string
	var cardsOut []string

	weightMap := make(map[int]string)
	var weightArr []int

	for k, val := range handMap {
		if val == 2 {
			numPairs++
			cardsOut = append([]string{k, k}, cardsOut...)
		} else {
			weightMap[CardVals[k]] = k
			weightArr = append(weightArr, CardVals[k])
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(weightArr)))
	for _, i := range weightArr {
		cardsOut = append(cardsOut, weightMap[i])
	}

	if numPairs == 1 {
		rank = "one-pair"
	}
	return rank, cardsOut
}

//HighCard function compares different value cards.
func HighCard(handMap map[string]int) (string, []string) {

	var cardsOut []string
	rank := "high-card"

	weightMap := make(map[int]string)
	var weightArr []int

	for k := range handMap {
		weightMap[CardVals[k]] = k
		weightArr = append(weightArr, CardVals[k])
	}

	sort.Sort(sort.Reverse(sort.IntSlice(weightArr)))
	for _, i := range weightArr {
		cardsOut = append(cardsOut, weightMap[i])
	}

	return rank, cardsOut
}
