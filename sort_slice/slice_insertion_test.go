package sort_slice

import (
	"reflect"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{name: "t1", args: args{nums: []int{9, 2, 5, 4, 1, 2, 3, 4, 5}}, want: []int{9, 5, 5, 4, 4, 3, 2, 2, 1}},
		{name: "t2", args: args{nums: []int{1, 2, 3, 4, 5}}, want: []int{5, 4, 3, 2, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InsertionSort(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_swap(t *testing.T) {
	type args struct {
		nums []int
		i    int
		j    int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{name: "t1", args: args{nums: []int{9, 2}, i: 0, j: 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			swap(tt.args.nums, tt.args.i, tt.args.j)
		})
	}
}
