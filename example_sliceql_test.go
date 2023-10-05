package sliceql

import (
	"fmt"
)

func ExampleNew() {
	// New() returns a slice with the given elements.
	s := New[int]([]int{1, 2, 3, 4, 5})
	fmt.Println(s.String())

	// Output:
	// [1 2 3 4 5]
}

func ExampleCreate() {
	// Create() returns a slice with the given count of elements.
	s := Create[int](10, func(index int) int {
		return index + 1
	})
	fmt.Println(s.String())

	// Output:
	// [1 2 3 4 5 6 7 8 9 10]
}

func ExampleQuery_All() {
	// All() returns true if the predicate is true for all elements in the sequence.
	s := New[int]([]int{1, 2, 3, 4, 5})
	fmt.Println(s.All(func(e int) bool {
		return e%2 == 0
	}))

	// Output:
	// false
}

func ExampleQuery_Any() {
	// Any() returns true if the predicate is true for any element in the sequence.
	s := New[int]([]int{1, 2, 3, 4, 5})
	fmt.Println(s.Any(func(e int) bool {
		return e%2 == 0
	}))

	// Output:
	// true
}

func ExampleQuery_At() {
	// At() returns the element at the given index.
	s := New[int]([]int{1, 2, 3, 4, 5})
	fmt.Println(s.At(2))

	// Output:
	// 3
}
