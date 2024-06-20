package lambda

import (
	"testing"
)

func TestMap(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	cb := func(i int) int {
		return i * 2
	}
	expected := []int{2, 4, 6, 8, 10}
	actual := Map(arr, cb)
	for i := 0; i < len(arr); i++ {
		if actual[i] != expected[i] {
			t.Errorf("expected %d but got %d", expected[i], actual[i])
		}
	}
}

func TestReverse(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	expected := []int{5, 4, 3, 2, 1}
	actual := Reverse(arr)
	for i := 0; i < len(arr); i++ {
		if actual[i] != expected[i] {
			t.Errorf("expected %d but got %d", expected[i], actual[i])
		}
	}
}

func TestFilter(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	cb := func(i int) bool {
		return i%2 == 0
	}
	expected := []int{2, 4}
	actual := Filter(arr, cb)
	for i := 0; i < len(actual); i++ {
		if actual[i] != expected[i] {
			t.Errorf("expected %d but got %d", expected[i], actual[i])
		}
	}
}

func TestHead(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	expected := 1
	actual := Head(arr)
	if actual != expected {
		t.Errorf("expected %d but got %d", expected, actual)
	}
}

func TestTail(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	expected := []int{2, 3, 4, 5}
	actual := Tail(arr)
	for i := 0; i < len(actual); i++ {
		if actual[i] != expected[i] {
			t.Errorf("expected %d but got %d", expected[i], actual[i])
		}
	}
}

func TestReduce(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	init := 0
	cb := func(acc int, i int) int {
		return acc + i
	}
	expected := 15
	actual := Reduce(arr, init, cb)
	if actual != expected {
		t.Errorf("expected %d but got %d", expected, actual)
	}
}
