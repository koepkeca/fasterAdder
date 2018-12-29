package fasterAdder

import (
	"testing"
)

func TestBasicAdder(t *testing.T) {
	tList := []int64{1, 2, 3, 4, 5}
	tmpSum := Summer(tList, 1)
	if tmpSum != 15 {
		t.Fatalf("Error: Expected 15 got %d\n", tmpSum)
	}
	return
}

func TestConcurrAdder(t *testing.T) {
	tList := []int64{-36, -18, 18, 36}
	tmpSum := Summer(tList, 2)
	if tmpSum != 0 {
		t.Fatalf("Error: Expected 0 got %d\n", tmpSum)
	}
	return
}

func TestBigConcurrAdder(t *testing.T) {
	tList := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25}
	tmpSum := Summer(tList, 2)
	if tmpSum != 325 {
		t.Fatalf("Error: Expected 325 got %d\n", tmpSum)
	}
	return
}

func TestEmptySlice(t *testing.T) {
	tList := []int64{}
	tmpSum := Summer(tList, 4)
	if tmpSum != 0 {
		t.Fatalf("Error: Expected 0 got %d\n", tmpSum)
	}
	return
}

func TestOddConcurrAdder(t *testing.T) {
	tList := []int64{2, 4, 6, 8, 10, 12, 14, 16, 18}
	tmpSum := Summer(tList, 2)
	if tmpSum != 90 {
		t.Fatalf("Error: Expected 90 got %d\n", tmpSum)
	}
	return
}
