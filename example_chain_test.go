// Package sliceql implements query support for Go language slices.
//
// Source code and other details for the project are available at GitHub:
//
//   https://github.com/dmundt/sliceql
//
package sliceql

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

type ByAge Query[Person]

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

// Example initializes a slice of Person objects and
// performs a series of operations on that slice using
// the NewQuery type.

// It demonstrates the usage of the Where() method
// to filter elements based on a condition, the Sort() method
// to sort the elements, and the Last() method
// to retrieve the last elementin the slice.
func Example() {
	people := []Person{
		{"Bob", 31},
		{"Jenny", 26},
		{"John", 42},
		{"Michael", 17},
	}
	s := NewQuery(people)
	fmt.Println(s)

	// Where() returns a new slice with the elements that match all persons
	// that have an age greater than 30.
	s.Where(func(p Person) bool {
		return p.Age > 30
	})
	fmt.Println(s)

	// Sort by age.
	sort.Sort(ByAge(*s))
	fmt.Println(s.Last())

	// Output:
	// [Bob: 31 Jenny: 26 John: 42 Michael: 17]
	// [Bob: 31 John: 42]
	// John: 42
}
