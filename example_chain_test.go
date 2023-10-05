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

func Example() {
	people := []Person{
		{"Bob", 31},
		{"Jenny", 26},
		{"John", 42},
		{"Michael", 17},
	}
	s := New(&people)
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
