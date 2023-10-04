package sliceql

import (
	"reflect"
	"testing"
)

func TestFrom(t *testing.T) {
	type args struct {
		s []int
	}
	tests := []struct {
		name string
		args args
		want *Query[int]
	}{
		{
			name: "nil slice",
			args: args{
				s: nil,
			},
			want: &Query[int]{s: nil},
		},
		{
			name: "zero slice",
			args: args{
				s: []int{},
			},
			want: &Query[int]{s: []int{}},
		},
		{
			name: "any slice",
			args: args{
				s: []int{1, 2, 3, 4, 5},
			},
			want: &Query[int]{s: []int{1, 2, 3, 4, 5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := From(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("From() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreate(t *testing.T) {
	type args struct {
		count  int
		create Creater[int]
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
				create: func(index int) int {
					return index
				},
			},
			want: &Query[int]{s: []int{}},
		},
		{
			name: "nil creater",
			args: args{
				count:  10,
				create: nil,
			},
			want: &Query[int]{s: []int{}},
		},
		{
			name: "zero count slice",
			args: args{
				count: 0,
				create: func(index int) int {
					return index
				},
			},
			want: &Query[int]{s: []int{}},
		},
		{
			name: "indexed slice",
			args: args{
				count: 5,
				create: func(index int) int {
					return index + 1
				},
			},
			want: &Query[int]{s: []int{1, 2, 3, 4, 5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Create(tt.args.count, tt.args.create); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() = %v, want %v", got, tt.want)
			}
		})
	}
}
