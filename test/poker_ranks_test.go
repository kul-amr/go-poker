package test

import (
	"go-poker/ranks"
	"reflect"
	"testing"
)

func TestFourOfaKind(t *testing.T) {

	expectedRank := "four-of-a-kind"
	expectedCardsOut := []string{"K", "K", "K", "K", "Q"}
	handMap := map[string]int{"K": 4, "Q": 1}
	rank, cardsOut := ranks.FourOfaKind(handMap)

	if rank != expectedRank || !reflect.DeepEqual(cardsOut, expectedCardsOut) {
		t.Errorf("Expected %s, %q, got %s, %q", expectedRank, expectedCardsOut, rank, cardsOut)
	}

	handMap = map[string]int{"K": 3, "Q": 1, "J": 1}
	rank, cardsOut = ranks.FourOfaKind(handMap)

	if len(rank) != 0 || len(cardsOut) != 0 {
		t.Errorf("Expected %s, %q, got %s, %q", expectedRank, expectedCardsOut, rank, cardsOut)
	}
}

func TestFullHouse(t *testing.T) {

	expectedRank := "full-house"
	expectedCardsOut := []string{"K", "K", "K", "Q", "Q"}
	handMap := map[string]int{"K": 3, "Q": 2}
	rank, cardsOut := ranks.FullHouse(handMap)

	if rank != expectedRank || !reflect.DeepEqual(cardsOut, expectedCardsOut) {
		t.Errorf("Expected %s, %q, got %s, %q", expectedRank, expectedCardsOut, rank, cardsOut)
	}

	handMap = map[string]int{"K": 3, "Q": 1, "J": 1}
	rank, cardsOut = ranks.FullHouse(handMap)

	if len(rank) != 0 || len(cardsOut) != 0 {
		t.Errorf("Expected %s, %q, got %s, %q", expectedRank, expectedCardsOut, rank, cardsOut)
	}

}

func TestTriple(t *testing.T) {

	expectedRank := "triples"
	expectedCardsOut := []string{"K", "K", "K", "Q", "J"}
	handMap := map[string]int{"K": 3, "Q": 1, "J": 1}
	rank, cardsOut := ranks.Triple(handMap)

	if rank != expectedRank || !reflect.DeepEqual(cardsOut, expectedCardsOut) {
		t.Errorf("Expected %s, %q, got %s, %q", expectedRank, expectedCardsOut, rank, cardsOut)
	}
}

func TestTwoPairs(t *testing.T) {

	expectedRank := "two-pairs"
	expectedCardsOut := []string{"K", "K", "Q", "Q", "J"}
	handMap := map[string]int{"K": 2, "Q": 2, "J": 1}
	rank, cardsOut := ranks.TwoPairs(handMap)

	if rank != expectedRank || !reflect.DeepEqual(cardsOut, expectedCardsOut) {
		t.Errorf("Expected %s, %q, got %s, %q", expectedRank, expectedCardsOut, rank, cardsOut)
	}

	handMap = map[string]int{"K": 2, "Q": 1, "J": 1, "6": 1}
	rank, cardsOut = ranks.TwoPairs(handMap)

	if len(rank) != 0 || len(cardsOut) != 0 {
		t.Errorf("Expected %s, %q, got %s, %q", expectedRank, expectedCardsOut, rank, cardsOut)
	}
}

func TestOnePair(t *testing.T) {

	expectedRank := "one-pair"
	expectedCardsOut := []string{"K", "K", "Q", "J", "T"}
	handMap := map[string]int{"K": 2, "Q": 1, "J": 1, "T": 1}
	rank, cardsOut := ranks.OnePair(handMap)

	if rank != expectedRank || !reflect.DeepEqual(cardsOut, expectedCardsOut) {
		t.Errorf("Expected %s, %q, got %s, %q", expectedRank, expectedCardsOut, rank, cardsOut)
	}
}

func TestHighCard(t *testing.T) {

	expectedRank := "high-card"
	expectedCardsOut := []string{"K", "Q", "J", "T", "6"}
	handMap := map[string]int{"K": 1, "Q": 1, "J": 1, "T": 1, "6": 1}
	rank, cardsOut := ranks.HighCard(handMap)

	if rank != expectedRank || !reflect.DeepEqual(cardsOut, expectedCardsOut) {
		t.Errorf("Expected %s, %q, got %s, %q", expectedRank, expectedCardsOut, rank, cardsOut)
	}
}
