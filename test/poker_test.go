package test

import (
	"go-poker/poker"
	"reflect"
	"testing"
)

func TestIsValid(t *testing.T) {

	expectedVal := true
	val := poker.IsValid("KKQQ5")

	if val != expectedVal {
		t.Errorf("Expected %t, got %t", expectedVal, val)
	}

	expectedVal = false
	val = poker.IsValid("KKQQ1")

	if val != expectedVal {
		t.Errorf("Expected %t, got %t", expectedVal, val)
	}

	expectedVal = false
	val = poker.IsValid("KKQQJJ")

	if val != expectedVal {
		t.Errorf("Expected %t, got %t", expectedVal, val)
	}
}

func TestGetRank(t *testing.T) {

	expectedRank := "two-pairs"
	expectedCards := []string{"K", "K", "Q", "Q", "5"}
	rank, cards := poker.GetRank("QQKK5")

	if rank != expectedRank {
		t.Errorf("Expected %s, got %s", expectedRank, rank)
	}

	if !reflect.DeepEqual(cards, expectedCards) {
		t.Errorf("Expected %q, got %q", expectedCards, cards)
	}

}

func TestHandCompare(t *testing.T) {

	//Same rank - high-card but hand1 has high weight card than hand2.
	expectedWinHand1, expectedWinHand2 := true, false
	winHand1, windHand2 := poker.HandCompare("KT9Q5", "KT5Q7")

	if winHand1 != expectedWinHand1 || windHand2 != expectedWinHand2 {
		t.Errorf("Expected %t and %t, got %t and %t", expectedWinHand1, expectedWinHand2, winHand1, windHand2)
	}

	//Same rank - high-card but hand2 has high weight card than hand1.
	expectedWinHand1, expectedWinHand2 = false, true
	winHand1, windHand2 = poker.HandCompare("KT7Q5", "KT5Q9")

	if winHand1 != expectedWinHand1 || windHand2 != expectedWinHand2 {
		t.Errorf("Expected %t and %t, got %t and %t", expectedWinHand1, expectedWinHand2, winHand1, windHand2)
	}

	//Hand1 rank - four-of-a-kind is higher than hand2 rank - high-card.
	expectedWinHand1, expectedWinHand2 = true, false
	winHand1, windHand2 = poker.HandCompare("KKKK5", "KT5Q7")

	if winHand1 != expectedWinHand1 || windHand2 != expectedWinHand2 {
		t.Errorf("Expected %t and %t, got %t and %t", expectedWinHand1, expectedWinHand2, winHand1, windHand2)
	}

	//It's a tie.
	expectedWinHand1, expectedWinHand2 = true, true
	winHand1, windHand2 = poker.HandCompare("KKKK5", "KKKK5")

	if winHand1 != expectedWinHand1 || windHand2 != expectedWinHand2 {
		t.Errorf("Expected %t and %t, got %t and %t", expectedWinHand1, expectedWinHand2, winHand1, windHand2)
	}

	//Hand2 rank - four-of-a-kind is higher than hand1 rank - two-pairs.
	expectedWinHand1, expectedWinHand2 = false, true
	winHand1, windHand2 = poker.HandCompare("KKQQ5", "99992")

	if winHand1 != expectedWinHand1 || windHand2 != expectedWinHand2 {
		t.Errorf("Expected %t and %t, got %t and %t", expectedWinHand1, expectedWinHand2, winHand1, windHand2)
	}

	//Invalid hands
	expectedWinHand1, expectedWinHand2 = false, false
	winHand1, windHand2 = poker.HandCompare("KKQQ5J", "99991")

	if winHand1 != expectedWinHand1 || windHand2 != expectedWinHand2 {
		t.Errorf("Expected %t and %t, got %t and %t", expectedWinHand1, expectedWinHand2, winHand1, windHand2)
	}

}
