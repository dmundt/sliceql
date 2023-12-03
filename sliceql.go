// Package sliceql implements query support for Go language slices.
//
// Source code and other details for the project are available at GitHub:
//
//	https://github.com/dmundt/sliceql
package sliceql

import (
	"fmt"
	"sort"
)

// A Query is a slice of comparable elements.
// The zero value of a Query is an empty slice.
type Query[E comparable] []E

// Create is a function that creates a new Query object.
//
// It takes two parameters: count, an integer representing the number of elements to create,
// and create, a function that takes an integer index and returns an element of type E.
//
// The function returns a pointer to a Query object.
func Create[E comparable](n int, f func(int) E) *Query[E] {
	if n <= 0 || f == nil {
		return &Query[E]{}
	}
	q := Query[E](make([]E, n))
	for i := 0; i < n; i++ {
		q[i] = f(i)
	}
	return &q
}

// NewQuery creates a new Query object.
//
// It takes a slice s of elements of type E and returns a pointer to a Query object.
func NewQuery[E comparable](v []E) *Query[E] {
	// Make shallow copy of slice elements.
	q := Query[E](v)
	return &q
}

// All checks if all elements in the query satisfy the given test.
//
// The test parameter is a function that takes an element of the query as input
// and returns a boolean value indicating whether the element satisfies the
// condition.
//
// The function returns true if all elements in the query satisfy the test.
// Otherwise, it returns false.
func (q *Query[E]) All(f func(E) bool) bool {
	if len(*q) < 1 || f == nil {
		return false
	}
	for _, e := range *q {
		if !f(e) {
			return false
		}
	}
	return true
}

// Any checks if any element in the Query satisfies the provided test.
//
// The test parameter is a function to apply to each element in the Query.
// Returns: true if any element satisfies the test, false otherwise.
func (q *Query[E]) Any(f func(E) bool) bool {
	if f == nil {
		return false
	}
	for _, e := range *q {
		if f(e) {
			return true
		}
	}
	return false
}

// At returns the element at the specified index.
//
// Parameter(s):
// - index: the index of the element to retrieve.
//
// Returns:
// - E: the element at the specified index.
func (q *Query[E]) At(i int) E {
	if len(*q) < 1 {
		panic("sliceql.At: empty list")
	}
	if i < 0 || i >= len(*q) {
		panic("sliceql.At: index out of bounds")
	}
	return (*q)[i]
}

// Contains checks if the query contains an element that satisfies the given function.
//
// Parameters:
// - f: a function that takes an element of type E and returns a boolean value.
//
// Returns:
// - bool: true if there is an element that satisfies the function, false otherwise.
func (q *Query[E]) Contains(f func(E) bool) bool {
	for _, e := range *q {
		if f(e) {
			return true
		}
	}
	return false
}

// Count returns the number of elements in the Query
// that satisfy the given Tester function.
//
// test: The function that tests each element in the Query.
// int: The number of elements that satisfy the given Tester function.
func (q *Query[E]) Count(f func(E) bool) int {
	n := 0
	for _, e := range *q {
		if f(e) {
			n++
		}
	}
	return n
}

// Each applies the given action function to each element in the Query.
//
// f: The function to apply to each element.
// Returns: The Query itself.
func (q *Query[E]) Each(f func(E) E) *Query[E] {
	for i := range *q {
		(*q)[i] = f((*q)[i])
	}
	return q
}

// Equal checks if the Query is equal to the given slice using the provided equality function.
//
// It takes a slice of type E and a function eq that takes two parameters of type E and returns a bool.
// The function returns a bool indicating whether the Query is equal to the given slice.
func (q *Query[E]) Equal(v []E, eq func(E, E) bool) bool {
	if len(*q) != len(v) {
		return false
	}
	for i := range *q {
		if !eq((*q)[i], v[i]) {
			return false
		}
	}
	return true
}

// First returns the first element of the Query.
//
// No parameters.
// Returns the element of type E.
func (q *Query[E]) First() E {
	if len(*q) < 1 {
		panic("sliceql.First: empty list")
	}
	return (*q)[0]
}

// Fold applies the combine function to each element in the Query
// and returns the accumulated result.
//
// It takes an initial value and a combine function as parameters.
// The initial value is the starting value for the accumulation.
// The combine function is a function that takes two arguments:
// the accumulated result and the current element in the Query.
// It returns the accumulated result after applying the combine
// function to each element in the Query.
// The return type is the same as the type of the initial value.
func (q *Query[E]) Fold(v E, f func(E, E) E) E {
	result := v
	for _, e := range *q {
		result = f(result, e)
	}
	return result
}

// Index returns the index of the first element in the Query[E] that satisfies the given condition.
//
// The parameter `f` is a function that takes an element of type E and returns
// a boolean value indicating whether the element satisfies the condition or not.
// The function returns an integer value representing the index of the first element
// that satisfies the condition. If no element satisfies the condition, it returns -1.
func (q *Query[E]) Index(f func(E) bool) int {
	for i, e := range *q {
		// v and e are type E, which has the comparable
		// constraint, so we can use == here.
		if f(e) {
			return i
		}
	}
	return -1
}

// Last returns the last element of the Query.
//
// It does not take any parameters.
// It returns the type E.
func (q *Query[E]) Last() E {
	if len(*q) < 1 {
		panic("sliceql.Last: empty list")
	}
	return (*q)[len(*q)-1]
}

// Reverse reverses the elements in the Query.
//
// No parameters.
// Returns a pointer to a Query[E].
func (q *Query[E]) Reverse() *Query[E] {
	for i, j := 0, len(*q)-1; i < j; i, j = i+1, j-1 {
		(*q)[i], (*q)[j] = (*q)[j], (*q)[i]
	}
	return q
}

// Skip removes the first 'count' elements from the Query.
//
// n: the number of elements to skip.
// Returns: a pointer to the modified Query.
func (q *Query[E]) Skip(n int) *Query[E] {
	if len(*q) < 1 {
		panic("sliceql.Skip: empty list")
	}
	if n < 0 || n > len(*q) {
		panic("sliceql.Skip: index out of bounds")
	}
	m := min(n, len(*q))
	*q = (*q)[m:]
	return q
}

// Sort sorts the elements in the Query slice using the provided less function.
//
// The less function should return true if the element at index i should be
// placed before the element at index j in the sorted slice.
//
// Returns a pointer to the original Query slice after sorting.
func (q *Query[E]) Sort(cmp func(E, E) bool) *Query[E] {
	if len(*q) < 1 || cmp == nil {
		return q
	}
	sort.Slice(*q, func(i, j int) bool {
		return cmp((*q)[i], (*q)[j])
	})
	return q
}

// String returns a string representation of the Query object.
func (q *Query[_]) String() string {
	return fmt.Sprintf("%v", *q)
}

// Take removes elements from the Query and returns
// a new Query with the specified number of elements.
//
// n: the number of elements to take from the Query.
// *Query[E]: a new Query with the specified number of elements.
func (q *Query[E]) Take(n int) *Query[E] {
	if len(*q) < 1 {
		panic("sliceql.Take: empty list")
	}
	if n < 0 || n > len(*q) {
		panic("sliceql.Take: index out of bounds")
	}
	m := min(n, len(*q))
	*q = (*q)[:m]
	return q
}

// ToSlice returns the Query[E] as a slice of E.
//
// No parameters.
// Returns a slice of type E.
func (q *Query[E]) ToSlice() []E {
	return []E(*q)
}

// Where filters the elements in the Query based on
// the provided test function.
//
// The test function is used to determine if an element
// should be included in the filtered result.
// The test function should accept an element of the
// Query's element type and return a boolean value.
//
// The function returns a pointer to the filtered Query.
func (q *Query[E]) Where(f func(E) bool) *Query[E] {
	if len(*q) < 1 || f == nil {
		return &Query[E]{}
	}
	result := Query[E](make([]E, 0))
	for _, e := range *q {
		if f(e) {
			result = append(result, e)
		}
	}
	*q = result
	return q
}
