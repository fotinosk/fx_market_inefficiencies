package graph

import (
	"testing"
)

func TestHappyPath(t *testing.T) {
	happy_path := []string{"eur", "usd", "rub", "eur"}
	_, err := TraverseGraph(happy_path)
	if err != nil {
		t.Fatalf("Happy path returns error: %s", err)
	}
}

func TestUnhappyPathWithCycle(t *testing.T) {
	unhappy_path := []string{"eur", "usd", "eur", "rub", "eur"}
	_, err2 := TraverseGraph(unhappy_path)
	if err2 == nil {
		t.Fatalf("Unhappy path does not return error")
	}
}

func TestUnhappyPathWithDifferentEnd(t *testing.T) {
	unhappy_path := []string{"eur", "usd", "rub"}
	_, err2 := TraverseGraph(unhappy_path)
	if err2 == nil {
		t.Fatalf("Unhappy path does not return error")
	}
}