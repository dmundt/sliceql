// Package sliceql implements query support for Go language slices.
//
// Source code and other details for the project are available at GitHub:
//
//	https://github.com/dmundt/sliceql
package sliceql

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

// Example initializes a slice of Person objects and
// performs a series of operations on that slice using
// the NewQuery type.
//
// It demonstrates the usage of the Where() method
// to filter elements based on a condition, the Sort() method
// to sort the elements, and the Last() method
// to retrieve the last element in the slice.
func Example() {
	s := NewQuery([]Person{
		{"Bob", 31},
		{"Jenny", 26},
		{"John", 42},
		{"Michael", 17},
	})
	fmt.Println(s)

	p := s.Where(func(p Person) bool {
		// Filter by age > 30.
		return p.Age > 30
	}).Sort(func(p1, p2 Person) bool {
		// Sort by ascending age.
		return p1.Age < p2.Age
	}).Last()
	fmt.Println(p)

	// Output:
	// [Bob: 31 Jenny: 26 John: 42 Michael: 17]
	// John: 42
}
