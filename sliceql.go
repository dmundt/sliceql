// Package sliceql implements query support for Go language slices.
//
// Source code and other details for the project are available at GitHub:
//
//   https://github.com/dmundt/sliceql
//
package sliceql

import "fmt"

// The Actioner interface is implemented by any function that
// takes a pointer to an element of type E. The element may be
// modified in place.
type Actioner[E any] func(*E)

// The Combiner interface is implemented by any function that
// takes two elements of type E and returns an element of type
// E by combining them.
type Combiner[E any] func(E, E) E

// The Creater interface is implemented by any function that
// takes the number of elements to create and returns an element
// of type E.
type Creater[E any] func(int) E

// The Tester interface is implemented by any function that
// takes an element of type E, tests whether it satisfies 
// the given condition and returns a boolean value.
type Tester[E any] func(E) bool

// A Query is a slice of comparable elements.
// The zero value of a Query is an empty slice.
type Query[E comparable] []E

// NewQuery creates a new Query object.
//
// It takes a slice s of elements of type E and returns a pointer to a Query object.
func NewQuery[E comparable](s []E) *Query[E] {
	// Make shallow copy of slice elements.
	q := Query[E](s)
	return &q
}

// From creates a new Query from the given slice.
//
// It makes a shallow copy of the slice elements.
// Returns a pointer to the created Query.
func From[E comparable](s []E) *Query[E] {
	// Make shallow copy of slice elements.
	q := Query[E](s)
	return &q
}

// As returns the Query[E] as a slice of E.
//
// No parameters.
// Returns a slice of E.
func (q *Query[E]) As() []E {
	return []E(*q)
}

// To returns the Query object as a slice of type E.
//
// No parameters.
// Returns a slice of type E.
func (q *Query[E]) To() []E {
	return []E(*q)
}

// Create is a function that creates a new Query object.
//
// It takes two parameters: count, an integer representing the number of elements to create,
// and create, a function that takes an integer index and returns an element of type E.
//
// The function returns a pointer to a Query object.
func Create[E comparable](count int, create Creater[E]) *Query[E] {
	if count <= 0 || create == nil {
		return &Query[E]{}
	}
	result := Query[E](make([]E, count))
	for i := 0; i < count; i++ {
		result[i] = create(i)
	}
	return &result
}

// All checks if all elements in the query satisfy the given test.
//
// The test parameter is a function that takes an element of the query as input
// and returns a boolean value indicating whether the element satisfies the
// condition.
//
// The function returns true if all elements in the query satisfy the test.
// Otherwise, it returns false.
func (q *Query[E]) All(test Tester[E]) bool {
	if len(*q) == 0 || test == nil {
		return false
	}
	for _, e := range *q {
		if !test(e) {
			return false
		}
	}
	return true
}

// Any checks if any element in the Query satisfies the provided test.
//
// The test parameter is a function to apply to each element in the Query.
// Returns: true if any element satisfies the test, false otherwise.
func (q *Query[E]) Any(test Tester[E]) bool {
	if test == nil {
		return false
	}
	for _, e := range *q {
		if test(e) {
			return true
		}
	}
	return false
}

// At returns the element at the specified index.
//
// Parameter(s):
// - index: the index of the element to retrieve.
// Return type(s):
// - E: the element at the specified index.
func (q *Query[E]) At(index int) E {
	if index < 0 || index >= len(*q) {
		return *new(E)
	}
	return (*q)[index]
}

// Count returns the number of elements in the Query
// that satisfy the given Tester function.
//
// test: The function that tests each element in the Query.
// int: The number of elements that satisfy the given Tester function.
func (q *Query[E]) Count(test Tester[E]) int {
	if test == nil {
		return 0
	}
	count := 0
	for _, e := range *q {
		if test(e) {
			count++
		}
	}
	return count
}

// Each applies the given action function to each element in the Query.
//
// action: The function to apply to each element.
// Returns: The Query itself.
func (q *Query[E]) Each(action Actioner[E]) *Query[E] {
	if action == nil {
		return q
	}
	for i := range *q {
		action(&(*q)[i])
	}
	return q
}

// First returns the first element of the Query.
//
// No parameters.
// Returns the element of type E.
func (q *Query[E]) First() E {
	if len(*q) == 0 {
		return *new(E)
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
func (q *Query[E]) Fold(initial E, combine Combiner[E]) E {
	if len(*q) == 0 || combine == nil {
		return *new(E)
	} else if len(*q) == 1 {
		return (*q)[0]
	}
	result := initial
	for i := 0; i < len(*q); i++ {
		result = combine(result, (*q)[i])
	}
	return result
}

// Index finds the index of an element in the Query.
//
// It takes an element e of type E as a parameter and returns
// the index of the first occurrence of e in the Query. If e
// is not found, it returns -1.
func (q *Query[E]) Index(e E) int {
	for i, v := range *q {
		// v and e are type S, which has the comparable
		// constraint, so we can use == here.
		if v == e {
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
	if len(*q) == 0 {
		return *new(E)
	}
	return (*q)[len(*q)-1]
}

// Skip removes the first 'count' elements from the Query.
//
// count: the number of elements to skip.
// Returns: a pointer to the modified Query.
func (q *Query[E]) Skip(count int) *Query[E] {
	if len(*q) == 0 || count < 0 {
		*q = make([]E, 0)
		return q
	}
	m := min(count, len(*q))
	*q = (*q)[m:]
	return q
}

// Take removes elements from the Query and returns
// a new Query with the specified number of elements.
//
// count: the number of elements to take from the Query.
// *Query[E]: a new Query with the specified number of elements.
func (q *Query[E]) Take(count int) *Query[E] {
	if len(*q) == 0 || count < 0 {
		*q = make([]E, 0)
		return q
	}
	m := min(count, len(*q))
	*q = (*q)[:m]
	return q
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
func (q *Query[E]) Where(test Tester[E]) *Query[E] {
	if len(*q) == 0 || test == nil {
		*q = make([]E, 0)
		return q
	}
	result := make([]E, 0, len(*q))
	for _, e := range *q {
		if test(e) {
			result = append(result, e)
		}
	}
	*q = result
	return q
}

// String returns a string representation of the Query object.
func (q *Query[_]) String() string {
	return fmt.Sprintf("%v", *q)
}
