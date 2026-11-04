package main

import (
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	slice := generateRandomElements(5)
	if len(slice) != 5 {
		t.Errorf("Ожидали длину 5, получили %d", len(slice))
	}

	empty := generateRandomElements(0)
	if empty != nil {
		t.Errorf("Ожидали nil, получили срез")
	}
}

func TestMaximum(t *testing.T) {
	data := []int{1, 5, 9, 3, 2}
	if got := maximum(data); got != 9 {
		t.Errorf("Ожидали 9, получили %d", got)
	}

	negative := []int{-10, -5, -20}
	if got := maximum(negative); got != -5 {
		t.Errorf("Ожидали -5, получили %d", got)
	}

	var empty []int
	if got := maximum(empty); got != 0 {
		t.Errorf("Ожидали 0, получили %d", got)
	}
}

func TestMaxChunks(t *testing.T) {
	data := []int{
		1, 2, 3, 4, 5, 6, 7, 8,
		10, 20, 99, 40, 50, 60, 70, 80,
	}
	if got := maxChunks(data); got != 99 {
		t.Errorf("Ожидали 99, получили %d", got)
	}

	firstMax := []int{100, 2, 3, 4, 5, 6, 7, 8}
	if got := maxChunks(firstMax); got != 100 {
		t.Errorf("Ожидали 100, получили %d", got)
	}

	lastMax := []int{1, 2, 3, 4, 5, 6, 7, 100}
	if got := maxChunks(lastMax); got != 100 {
		t.Errorf("Ожидали 100, получили %d", got)
	}

	var empty []int
	if got := maxChunks(empty); got != 0 {
		t.Errorf("Ожидали 0, получили %d", got)
	}
}
