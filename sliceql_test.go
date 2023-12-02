// Package sliceql implements query support for Go language slices.
//
// Source code and other details for the project are available at GitHub:
//
//	https://github.com/dmundt/sliceql
package sliceql

import (
	"reflect"
	"testing"
)

func Test_Create(t *testing.T) {
	type args struct {
		count int
		f     Creater[int]
	}
	tests := []struct {
		name string
		args args
		want *Query[int]
	}{
		{
			name: "negative count",
			args: args{
				count: -1,
				f: func(index int) int {
					return index
				},
			},
			want: &Query[int]{},
		},
		{
			name: "nil creater",
			args: args{
				count: 10,
				f:     nil,
			},
			want: &Query[int]{},
		},
		{
			name: "zero count slice",
			args: args{
				count: 0,
				f: func(index int) int {
					return index
				},
			},
			want: &Query[int]{},
		},
		{
			name: "indexed slice",
			args: args{
				count: 5,
				f: func(index int) int {
					return index + 1
				},
			},
			want: &Query[int]{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Create(tt.args.count, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuery_All(t *testing.T) {
	type args struct {
		q *Query[int]
		f Tester[int]
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
		f Tester[int]
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
		q     *Query[int]
		index int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "valid index",
			args: args{
				q:     &Query[int]{1, 2, 3, 4, 5},
				index: 2,
			},
			want: 3,
		},
		{
			name: "negative index",
			args: args{
				q:     &Query[int]{1, 2, 3, 4, 5},
				index: -1,
			},
			want: 0,
		},
		{
			name: "index out of range",
			args: args{
				q:     &Query[int]{1, 2, 3, 4, 5},
				index: 10,
			},
			want: 0,
		},
		{
			name: "empty slice",
			args: args{
				q:     &Query[int]{},
				index: 0,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.q.At(tt.args.index); got != tt.want {
				t.Errorf("Query.At() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuery_Count(t *testing.T) {
	type args struct {
		q *Query[int]
		f Tester[int]
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "nil predicate",
			args: args{
				q: &Query[int]{1, 2, 3, 4, 5},
				f: nil,
			},
			want: 0,
		},
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
		f Actioner[int]
	}
	tests := []struct {
		name string
		args args
		want *Query[int]
	}{
		{
			name: "nil action",
			args: args{
				q: &Query[int]{1, 2, 3, 4, 5},
				f: nil,
			},
			want: &Query[int]{1, 2, 3, 4, 5},
		},
		{
			name: "nil sequence",
			args: args{
				q: &Query[int]{},
				f: func(num *int) {
					t.Errorf("Unexpected action function call")
				},
			},
			want: &Query[int]{},
		},
		{
			name: "empty slice",
			args: args{
				q: &Query[int]{},
				f: func(num *int) {
					t.Errorf("Unexpected action function call")
				},
			},
			want: &Query[int]{},
		},
		{
			name: "non empty slice",
			args: args{
				q: &Query[int]{1, 2, 3, 4, 5},
				f: func(num *int) {
					(*num)++
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

func TestQuery_First(t *testing.T) {
	type args struct {
		q *Query[int]
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "nil slice",
			args: args{
				q: &Query[int]{},
			},
			want: 0,
		},
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

func TestQuery_Index(t *testing.T) {
	type args struct {
		q *Query[int]
		e int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "nil slice",
			args: args{
				q: &Query[int]{},
				e: 0,
			},
			want: -1,
		},
		{
			name: "empty slice",
			args: args{
				q: &Query[int]{},
				e: 0,
			},
			want: -1,
		},
		{
			name: "unknown value",
			args: args{
				q: &Query[int]{1, 2, 3, 4, 5},
				e: 10,
			},
			want: -1,
		},
		{
			name: "first index",
			args: args{
				q: &Query[int]{1, 2, 3, 4, 5},
				e: 1,
			},
			want: 0,
		},
		{
			name: "some index",
			args: args{
				q: &Query[int]{1, 2, 3, 4, 5},
				e: 3,
			},
			want: 2,
		},
		{
			name: "last index",
			args: args{
				q: &Query[int]{1, 2, 3, 4, 5},
				e: 5,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.q.Index(tt.args.e); got != tt.want {
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
		{
			name: "nil slice",
			args: args{
				q: &Query[int]{}},
			want: 0,
		},
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

func TestQuery_Skip(t *testing.T) {
	type args struct {
		q     *Query[int]
		count int
	}
	tests := []struct {
		name string
		args args
		want *Query[int]
	}{
		{
			name: "negative count",
			args: args{
				q:     &Query[int]{1, 2, 3, 4, 5},
				count: -1,
			},
			want: &Query[int]{},
		},
		{
			name: "empty slice",
			args: args{
				q:     &Query[int]{},
				count: 1,
			},
			want: &Query[int]{},
		},
		{
			name: "non empty slice 1",
			args: args{
				q:     &Query[int]{1, 2, 3, 4, 5},
				count: 0,
			},
			want: &Query[int]{1, 2, 3, 4, 5},
		},
		{
			name: "non empty slice 2",
			args: args{
				q:     &Query[int]{1, 2, 3, 4, 5},
				count: 3,
			},
			want: &Query[int]{4, 5},
		},
		{
			name: "complete slice",
			args: args{
				q:     &Query[int]{1, 2, 3, 4, 5},
				count: 5,
			},
			want: &Query[int]{},
		},
		{
			name: "large index",
			args: args{
				q:     &Query[int]{1, 2, 3, 4, 5},
				count: 10,
			},
			want: &Query[int]{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.q.Skip(tt.args.count); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Query.Skip() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuery_Take(t *testing.T) {
	type args struct {
		q     *Query[int]
		count int
	}
	tests := []struct {
		name string
		args args
		want *Query[int]
	}{
		{
			name: "negative count",
			args: args{
				q:     &Query[int]{1, 2, 3, 4, 5},
				count: -1,
			},
			want: &Query[int]{},
		},
		{
			name: "empty slice",
			args: args{
				q:     &Query[int]{},
				count: 1,
			},
			want: &Query[int]{},
		},
		{
			name: "non empty slice",
			args: args{
				q:     &Query[int]{1, 2, 3, 4, 5},
				count: 0,
			},
			want: &Query[int]{},
		},
		{
			name: "non empty slice",
			args: args{
				q:     &Query[int]{1, 2, 3, 4, 5},
				count: 3,
			},
			want: &Query[int]{1, 2, 3},
		},
		{
			name: "complete slice",
			args: args{
				q:     &Query[int]{1, 2, 3, 4, 5},
				count: 5,
			},
			want: &Query[int]{1, 2, 3, 4, 5},
		},
		{
			name: "complete slice",
			args: args{
				q:     &Query[int]{1, 2, 3, 4, 5},
				count: 10,
			},
			want: &Query[int]{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.q.Take(tt.args.count); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Query.Take() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuery_Where(t *testing.T) {
	type args struct {
		q *Query[int]
		f Tester[int]
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
			name: "nil sequence",
			args: args{
				q: &Query[int]{},
				f: func(num int) bool {
					return num > 0
				},
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
