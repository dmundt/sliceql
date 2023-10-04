package sliceql

import "fmt"

type Actioner[E any] func(*E)

type Combiner[E any] func(E, E) E

type Creater[E any] func(int) E

type Tester[E any] func(E) bool

type Query[E comparable] struct {
	s []E
}

func (q *Query[E]) As() []E {
	return q.s
}

func (q *Query[E]) To() []E {
	return q.s
}

func From[E comparable](s []E) *Query[E] {
	return &Query[E]{s: s}
}

func Create[E comparable](count int, create Creater[E]) *Query[E] {
	if count <= 0 || create == nil {
		return &Query[E]{s: make([]E, 0)}
	}
	result := make([]E, count)
	for i := 0; i < count; i++ {
		result[i] = create(i)
	}
	return &Query[E]{s: result}
}

func (q *Query[E]) All(test Tester[E]) bool {
	if len(q.s) == 0 || test == nil {
		return false
	}
	for _, e := range q.s {
		if !test(e) {
			return false
		}
	}
	return true
}

func (q *Query[E]) Any(test Tester[E]) bool {
	if test == nil {
		return false
	}
	for _, e := range q.s {
		if test(e) {
			return true
		}
	}
	return false
}

func (q *Query[E]) At(index int) E {
	if index < 0 || index >= len(q.s) {
		return *new(E)
	}
	return q.s[index]
}

func (q *Query[E]) Count(test Tester[E]) int {
	if test == nil {
		return 0
	}
	count := 0
	for _, e := range q.s {
		if test(e) {
			count++
		}
	}
	return count
}

func (q *Query[E]) Each(action Actioner[E]) *Query[E] {
	if action == nil {
		return q
	}
	for i := range q.s {
		action(&q.s[i])
	}
	return q
}

func (q *Query[E]) First() E {
	if len(q.s) == 0 {
		return *new(E)
	}
	return q.s[0]
}

func (q *Query[E]) Fold(initial E, combine Combiner[E]) E {
	if len(q.s) == 0 || combine == nil {
		return *new(E)
	} else if len(q.s) == 1 {
		return q.s[0]
	}
	result := initial
	for i := 0; i < len(q.s); i++ {
		result = combine(result, q.s[i])
	}
	return result
}

func (q *Query[E]) Index(s []E, e E) int {
	for i, v := range s {
		// v and e are type S, which has the comparable
		// constraint, so we can use == here.
		if v == e {
			return i
		}
	}
	return -1
}

func (q *Query[E]) Last() E {
	if len(q.s) == 0 {
		return *new(E)
	}
	return q.s[len(q.s)-1]
}

func (q *Query[E]) Skip(count int) *Query[E] {
	if len(q.s) == 0 || count < 0 {
		q.s = make([]E, 0)
		return q
	}
	m := min(count, len(q.s))
	q.s = q.s[m:]
	return q
}

func (q *Query[E]) Take(count int) *Query[E] {
	if len(q.s) == 0 || count < 0 {
		q.s = make([]E, 0)
		return q
	}
	m := min(count, len(q.s))
	q.s = q.s[:m]
	return q
}

func (q *Query[E]) Where(test Tester[E]) *Query[E] {
	if len(q.s) == 0 || test == nil {
		q.s = make([]E, 0)
		return q
	}
	result := make([]E, 0, len(q.s))
	for _, e := range q.s {
		if test(e) {
			result = append(result, e)
		}
	}
	q.s = result
	return q
}

func (q *Query[E]) String() string {
	return fmt.Sprintf("%v", q.s)
}
