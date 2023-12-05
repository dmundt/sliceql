// Copyright 2023 Daniel Mundt. All rights reserved.
// Use of this source code is governed by a
// MIT license that can be found in the LICENSE file.

package sliceql

import (
	"reflect"
	"testing"
)

func Test_Create(t *testing.T) {
	type args struct {
		n int
		f func(int) int
	}
	tests := []struct {
		name string
		args args
		want *Query[int]
	}{
		{
			name: "negative count",
			args: args{
				n: -1,
				f: func(index int) int {
					return index
				},
			},
			want: &Query[int]{},
		},
		{
			name: "nil creater",
			args: args{
				n: 10,
				f: nil,
			},
			want: &Query[int]{},
		},
		{
			name: "zero count slice",
			args: args{
				n: 0,
				f: func(index int) int {
					return index
				},
			},
			want: &Query[int]{},
		},
		{
			name: "indexed slice",
			args: args{
				n: 5,
				f: func(index int) int {
					return index + 1
				},
			},
			want: &Query[int]{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Create(tt.args.n, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_NewQuery(t *testing.T) {
	type args struct {
		v []int
	}
	tests := []struct {
		name string
		args args
		want *Query[int]
	}{
		{
			name: "empty slice",
			args: args{
				v: []int{},
			},
			want: &Query[int]{},
		},
		{
			name: "non-empty slice",
			args: args{
				v: []int{1, 2, 3, 4, 5},
			},
			want: &Query[int]{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewQuery(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuery_All(t *testing.T) {
	type args struct {
		q *Query[int]
		f func(int) bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "nil predicate",
			args: args{
				q: &Query[int]{1, 2, 3, 4, 5},
				f: nil,
			},
			want: false,
		},
		{
			name: "nil sequence",
			args: args{
				q: &Query[int]{},
				f: func(num int) bool {
					return true
				},
			},
			want: false,
		},
		{
			name: "empty slice (false condition)",
			args: args{
				q: &Query[int]{},
				f: func(num int) bool {
					return false
				},
			},
			want: false,
		},
		{
			name: "empty slice (true condition)",
			args: args{
				q: &Query[int]{},
				f: func(num int) bool {
					return true
				},
			},
			want: false,
		},
		{
			name: "false condition",
			args: args{
				q: &Query[int]{1, 2, 3, 4, 5},
				f: func(num int) bool {
					return false
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.q.All(tt.args.f); got != tt.want {
				t.Errorf("Query.All() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuery_Any(t *testing.T) {
	type args struct {
		q *Query[int]
		f func(int) bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "nil predicate",
			args: args{
				q: &Query[int]{1, 2, 3, 4, 5},
				f: nil,
			},
			want: false,
		},
		{
			name: "nil sequence",
			args: args{
				q: &Query[int]{},
				f: func(num int) bool {
					return true
				},
			},
			want: false,
		},
		{
			name: "empty slice (false condition)",
			args: args{
				q: &Query[int]{},
				f: func(num int) bool {
					return false
				},
			},
			want: false,
		},
		{
			name: "empty slice (true condition)",
			args: args{
				q: &Query[int]{},
				f: func(num int) bool {
					return true
				},
			},
			want: false,
		},
		{
			name: "true condition",
			args: args{
				q: &Query[int]{1, 2, 3, 4, 5},
				f: func(num int) bool {
					return true
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.q.Any(tt.args.f); got != tt.want {
				t.Errorf("Query.Any() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuery_At(t *testing.T) {
	type args struct {
		q *Query[int]
		i int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "valid index",
			args: args{
				q: &Query[int]{1, 2, 3, 4, 5},
				i: 2,
			},
			want: 3,
		},
		// {
		// 	name: "negative index",
		// 	args: args{
		// 		q: &Query[int]{1, 2, 3, 4, 5},
		// 		i: -1,
		// 	},
		// 	want: nil,
		// },
		// {
		// 	name: "index out of range",
		// 	args: args{
		// 		q: &Query[int]{1, 2, 3, 4, 5},
		// 		i: 10,
		// 	},
		// 	want: nil,
		// },
		// {
		// 	name: "empty slice",
		// 	args: args{
		// 		q: &Query[int]{},
		// 		i: 0,
		// 	},
		// 	want: nil,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.q.At(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Query.At() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuery_Contains(t *testing.T) {
	type args struct {
		q *Query[int]
		f func(int) bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "nil slice",
			args: args{
				q: &Query[int]{},
				f: func(v int) bool {
					return v == 0
				},
			},
			want: false,
		},
		{
			name: "empty slice",
			args: args{
				q: &Query[int]{},
				f: func(v int) bool {
					return v == 0
				},
			},
			want: false,
		},
		{
			name: "unknown value",
			args: args{
				q: &Query[int]{1, 2, 3, 4, 5},
				f: func(v int) bool {
					return v == 10
				},
			},
			want: false,
		},
		{
			name: "first index",
			args: args{
				q: &Query[int]{1, 2, 3, 4, 5},
				f: func(v int) bool {
					return v == 1
				},
			},
			want: true,
		},
		{
			name: "some index",
			args: args{
				q: &Query[int]{1, 2, 3, 4, 5},
				f: func(v int) bool {
					return v == 3
				},
			},
			want: true,
		},
		{
			name: "last index",
			args: args{
				q: &Query[int]{1, 2, 3, 4, 5},
				f: func(v int) bool {
					return v == 5
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.q.Contains(tt.args.f); got != tt.want {
				t.Errorf("Query.Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuery_Count(t *testing.T) {
	type args struct {
		q *Query[int]
		f func(int) bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {
		// 	name: "nil predicate",
		// 	args: args{
		// 		q: &Query[int]{1, 2, 3, 4, 5},
		// 		f: nil,
		// 	},
		// 	want: 0,
		// },
		{
			name: "nil sequence",
			args: args{
				q: &Query[int]{},
				f: func(num int) bool {
					return true
				},
			},
			want: 0,
		},
		{
			name: "empty slice (false condition)",
			args: args{
				q: &Query[int]{},
				f: func(num int) bool {
					return false
				},
			},
			want: 0,
		},
		{
			name: "empty slice (true condition)",
			args: args{
				q: &Query[int]{},
				f: func(num int) bool {
					return true
				},
			},
			want: 0,
		},
		{
			name: "true condition",
			args: args{
				q: &Query[int]{1, 2, 3, 4, 5},
				f: func(num int) bool {
					return true
				},
			},
			want: 5,
		},
		{
			name: "even condition",
			args: args{
				q: &Query[int]{1, 2, 3, 4, 5},
				f: func(num int) bool {
					return num%2 == 0
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.q.Count(tt.args.f); got != tt.want {
				t.Errorf("Query.Count() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuery_Each(t *testing.T) {
	type args struct {
		q *Query[int]
		f func(int) int
	}
	tests := []struct {
		name string
		args args
		want *Query[int]
	}{
		// {
		// 	name: "nil action",
		// 	args: args{
		// 		q: &Query[int]{1, 2, 3, 4, 5},
		// 		f: nil,
		// 	},
		// 	want: &Query[int]{1, 2, 3, 4, 5},
		// },
		{
			name: "nil sequence",
			args: args{
				q: &Query[int]{},
				f: func(num int) int {
					t.Errorf("Unexpected action function call")
					return 0
				},
			},
			want: &Query[int]{},
		},
		{
			name: "empty slice",
			args: args{
				q: &Query[int]{},
				f: func(num int) int {
					t.Errorf("Unexpected action function call")
					return 0
				},
			},
			want: &Query[int]{},
		},
		{
			name: "non empty slice",
			args: args{
				q: &Query[int]{1, 2, 3, 4, 5},
				f: func(num int) int {
					return num + 1
				},
			},
			want: &Query[int]{2, 3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.args.q.Each(tt.args.f); !reflect.DeepEqual(tt.args.q, tt.want) {
				t.Errorf("Query.Each() = %v, want %v", tt.args.q, tt.want)
			}
		})
	}
}

func TestQuery_Equal(t *testing.T) {
	type args struct {
		q *Query[int]
		v []int
		f func(int, int) bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty slices",
			args: args{
				q: &Query[int]{},
				v: []int{},
				f: func(a, b int) bool {
					return a == b
				},
			},
			want: true,
		},
		{
			name: "different lengths",
			args: args{
				q: &Query[int]{1, 2, 3, 4, 5},
				v: []int{1, 2, 3, 4},
				f: func(a, b int) bool {
					return a == b
				},
			},
			want: false,
		},
		{
			name: "equal",
			args: args{
				q: &Query[int]{1, 2, 3, 4, 5},
				v: []int{1, 2, 3, 4, 5},
				f: func(a, b int) bool {
					return a == b
				},
			},
			want: true,
		},
		{
			name: "not equal (elements)",
			args: args{
				q: &Query[int]{1, 2, 3, 4, 5},
				v: []int{1, 2, 3, 4, 6},
				f: func(a, b int) bool {
					return a == b
				},
			},
			want: false,
		},
		{
			name: "equal (disjunct w/ neq)",
			args: args{
				q: &Query[int]{1, 2, 3, 4, 5},
				v: []int{6, 7, 8, 9, 10},
				f: func(a, b int) bool {
					return a != b
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.q.Equal(tt.args.v, tt.args.f); got != tt.want {
				t.Errorf("Query.Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuery_First(t *testing.T) {
	type args struct {
		q *Query[int]
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {
		// 	name: "nil slice",
		// 	args: args{
		// 		q: &Query[int]{},
		// 	},
		// 	want: nil,
		// },
		{
			name: "one item",
			args: args{
				q: &Query[int]{1}},
			want: 1,
		},
		{
			name: "many items",
			args: args{
				q: &Query[int]{1, 2, 3, 4, 5},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.q.First(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Query.First() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuery_Fold(t *testing.T) {
	type args struct {
		q *Query[int]
		v int
		f func(int, int) int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "empty slice",
			args: args{
				q: &Query[int]{},
				v: 0,
				f: func(acc int, num int) int {
					return acc + num
				},
			},
			want: 0,
		},
		{
			name: "empty slice (w/ offset)",
			args: args{
				q: &Query[int]{},
				v: 10,
				f: func(acc int, num int) int {
					return acc + num
				},
			},
			want: 10,
		},
		{
			name: "nil combiner",
			args: args{
				q: &Query[int]{},
				v: 0,
				f: nil,
			},
			want: 0,
		},
		{
			name: "nil combiner  (w/ offset)",
			args: args{
				q: &Query[int]{},
				v: 10,
				f: nil,
			},
			want: 10,
		},
		{
			name: "one item",
			args: args{
				q: &Query[int]{1},
				v: 0,
				f: func(acc int, num int) int {
					return acc + num
				},
			},
			want: 1,
		},
		{
			name: "one item (w/ offset)",
			args: args{
				q: &Query[int]{1},
				v: 10,
				f: func(acc int, num int) int {
					return acc + num
				},
			},
			want: 11,
		},
		{
			name: "many items",
			args: args{
				q: &Query[int]{1, 2, 3, 4, 5},
				v: 0,
				f: func(acc int, num int) int {
					return acc + num
				},
			},
			want: 15,
		},
		{
			name: "many items (w/ offset)",
			args: args{
				q: &Query[int]{1, 2, 3, 4, 5},
				v: 10,
				f: func(acc int, num int) int {
					return acc + num
				},
			},
			want: 25,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.q.Fold(tt.args.v, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Query.Fold() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuery_Index(t *testing.T) {
	type args struct {
		q *Query[int]
		f func(int) bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "empty slice",
			args: args{
				q: &Query[int]{},
				f: func(v int) bool {
					return v == 0
				},
			},
			want: -1,
		},
		{
			name: "unknown value",
			args: args{
				q: &Query[int]{1, 2, 3, 4, 5},
				f: func(v int) bool {
					return v == 10
				},
			},
			want: -1,
		},
		{
			name: "first index",
			args: args{
				q: &Query[int]{1, 2, 3, 4, 5},
				f: func(v int) bool {
					return v == 1
				},
			},
			want: 0,
		},
		{
			name: "some index",
			args: args{
				q: &Query[int]{1, 2, 3, 4, 5},
				f: func(v int) bool {
					return v == 3
				},
			},
			want: 2,
		},
		{
			name: "last index",
			args: args{
				q: &Query[int]{1, 2, 3, 4, 5},
				f: func(v int) bool {
					return v == 5
				},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.q.Index(tt.args.f); got != tt.want {
				t.Errorf("Query.Index() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuery_Last(t *testing.T) {
	type args struct {
		q *Query[int]
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {
		// 	name: "nil slice",
		// 	args: args{
		// 		q: &Query[int]{}},
		// 	want: nil,
		// },
		{
			name: "one item",
			args: args{
				q: &Query[int]{1},
			},
			want: 1,
		},
		{
			name: "many items",
			args: args{
				q: &Query[int]{1, 2, 3, 4, 5},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.q.Last(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Query.Last() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuery_Reverse(t *testing.T) {
	type args struct {
		q *Query[int]
	}
	tests := []struct {
		name string
		args args
		want *Query[int]
	}{
		{
			name: "empty slice",
			args: args{
				q: &Query[int]{},
			},
			want: &Query[int]{},
		},
		{
			name: "non-empty slice",
			args: args{
				q: &Query[int]{1, 2, 3, 4, 5},
			},
			want: &Query[int]{5, 4, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.q.Reverse(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Query.Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuery_Skip(t *testing.T) {
	defer func() { _ = recover() }()
	type args struct {
		q *Query[int]
		n int
	}
	tests := []struct {
		name string
		args args
		want *Query[int]
	}{
		// {
		// 	name: "negative n",
		// 	args: args{
		// 		q: &Query[int]{1, 2, 3, 4, 5},
		// 		n: -1,
		// 	},
		// 	want: &Query[int]{},
		// // },
		// {
		// 	name: "empty slice",
		// 	args: args{
		// 		q: &Query[int]{},
		// 		n: 1,
		// 	},
		// 	want: &Query[int]{},
		// },
		{
			name: "first element",
			args: args{
				q: &Query[int]{1, 2, 3, 4, 5},
				n: 0,
			},
			want: &Query[int]{1, 2, 3, 4, 5},
		},
		{
			name: "any element",
			args: args{
				q: &Query[int]{1, 2, 3, 4, 5},
				n: 3,
			},
			want: &Query[int]{4, 5},
		},
		{
			name: "last element",
			args: args{
				q: &Query[int]{1, 2, 3, 4, 5},
				n: 5,
			},
			want: &Query[int]{},
		},
		// {
		// 	name: "n overflow",
		// 	args: args{
		// 		q: &Query[int]{1, 2, 3, 4, 5},
		// 		n: 10,
		// 	},
		// 	want: &Query[int]{},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.q.Skip(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Query.Skip() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuery_Sort(t *testing.T) {
	type args struct {
		q *Query[int]
		f func(int, int) bool
	}
	tests := []struct {
		name string
		args args
		want *Query[int]
	}{
		{
			name: "nil action",
			args: args{
				q: &Query[int]{5, 4, 3, 2, 1},
				f: nil,
			},
			want: &Query[int]{5, 4, 3, 2, 1},
		},
		{
			name: "nil sequence",
			args: args{
				q: &Query[int]{},
				f: func(i, j int) bool {
					t.Errorf("Unexpected action function call")
					return i < j
				},
			},
			want: &Query[int]{},
		},
		{
			name: "empty slice",
			args: args{
				q: &Query[int]{},
				f: func(i, j int) bool {
					t.Errorf("Unexpected action function call")
					return i < j
				},
			},
			want: &Query[int]{},
		},
		{
			name: "ascending sort",
			args: args{
				q: &Query[int]{3, 2, 5, 4, 1},
				f: func(i, j int) bool {
					return i < j
				},
			},
			want: &Query[int]{1, 2, 3, 4, 5},
		},
		{
			name: "descending sort",
			args: args{
				q: &Query[int]{3, 2, 5, 4, 1},
				f: func(i, j int) bool {
					return i > j
				},
			},
			want: &Query[int]{5, 4, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.args.q.Sort(tt.args.f); !reflect.DeepEqual(tt.args.q, tt.want) {
				t.Errorf("Query.Sort() = %v, want %v", tt.args.q, tt.want)
			}
		})
	}
}

func TestQuery_String(t *testing.T) {
	type args struct {
		q *Query[int]
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "nil slice",
			args: args{
				q: &Query[int]{},
			},
			want: "[]",
		},
		{
			name: "empty slice",
			args: args{
				q: &Query[int]{},
			},
			want: "[]",
		},
		{
			name: "non empty slice",
			args: args{
				q: &Query[int]{1, 2, 3, 4, 5},
			},
			want: "[1 2 3 4 5]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.q.String(); got != tt.want {
				t.Errorf("Query.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuery_Take(t *testing.T) {
	type args struct {
		q *Query[int]
		n int
	}
	tests := []struct {
		name string
		args args
		want *Query[int]
	}{
		// {
		// 	name: "negative count",
		// 	args: args{
		// 		q: &Query[int]{1, 2, 3, 4, 5},
		// 		n: -1,
		// 	},
		// 	want: &Query[int]{},
		// },
		// {
		// 	name: "empty slice",
		// 	args: args{
		// 		q: &Query[int]{},
		// 		n: 1,
		// 	},
		// 	want: &Query[int]{},
		// },
		{
			name: "first element",
			args: args{
				q: &Query[int]{1, 2, 3, 4, 5},
				n: 0,
			},
			want: &Query[int]{},
		},
		{
			name: "any element",
			args: args{
				q: &Query[int]{1, 2, 3, 4, 5},
				n: 3,
			},
			want: &Query[int]{1, 2, 3},
		},
		{
			name: "last element",
			args: args{
				q: &Query[int]{1, 2, 3, 4, 5},
				n: 5,
			},
			want: &Query[int]{1, 2, 3, 4, 5},
		},
		// {
		// 	name: "index out of bounds",
		// 	args: args{
		// 		q: &Query[int]{1, 2, 3, 4, 5},
		// 		n: 10,
		// 	},
		// 	want: &Query[int]{1, 2, 3, 4, 5},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.q.Take(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Query.Take() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuery_ToSlice(t *testing.T) {
	type args struct {
		q *Query[int]
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "empty slice",
			args: args{
				q: &Query[int]{},
			},
			want: []int{},
		},
		{
			name: "complete slice",
			args: args{
				q: &Query[int]{1, 2, 3, 4, 5},
			},
			want: []int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.q.ToSlice(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Query.ToSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuery_Where(t *testing.T) {
	type args struct {
		q *Query[int]
		f func(int) bool
	}
	tests := []struct {
		name string
		args args
		want *Query[int]
	}{
		{
			name: "nil tester",
			args: args{
				q: &Query[int]{1, 2, 3, 4, 5},
				f: nil,
			},
			want: &Query[int]{},
		},
		{
			name: "empty slice",
			args: args{
				q: &Query[int]{},
				f: func(num int) bool {
					return num > 0
				},
			},
			want: &Query[int]{},
		},
		{
			name: "non empty slice",
			args: args{
				q: &Query[int]{1, 2, 3, 4, 5},
				f: func(num int) bool {
					return num%2 == 0
				},
			},
			want: &Query[int]{2, 4},
		},
		{
			name: "false condition",
			args: args{
				q: &Query[int]{1, 2, 3, 4, 5},
				f: func(num int) bool {
					return num > 5
				},
			},
			want: &Query[int]{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.q.Where(tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Query.Where() = %v, want %v", got, tt.want)
			}
		})
	}
}
