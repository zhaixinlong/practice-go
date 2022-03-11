package sort_slice

import (
	"reflect"
	"testing"
)

func Test_partition(t *testing.T) {
	type args struct {
		arr  []int
		low  int
		high int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "t1", args: args{arr: []int{9, 2, 5, 4, 1, 2, 3, 4, 5}, low: 0, high: 8}, want: 7},
		{name: "t1", args: args{arr: []int{1, 2, 3, 4, 5}, low: 0, high: 4}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partition(tt.args.arr, tt.args.low, tt.args.high); got != tt.want {
				t.Errorf("partition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuickSortRange(t *testing.T) {
	type args struct {
		arr  []int
		low  int
		high int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{name: "t1", args: args{arr: []int{9, 2, 5, 4, 1, 2, 3, 4, 5}, low: 0, high: 8}},
		{name: "t1", args: args{arr: []int{1, 2, 3, 4, 5}, low: 0, high: 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			QuickSortRange(tt.args.arr, tt.args.low, tt.args.high)
		})
	}
}

func TestQuickSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{name: "t1", args: args{arr: []int{9, 2, 5, 4, 1, 2, 3, 4, 5}}, want: []int{1, 2, 2, 3, 4, 4, 5, 5, 9}},
		{name: "t1", args: args{arr: []int{1, 2, 3, 4, 5}}, want: []int{1, 2, 3, 4, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QuickSort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuickSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
