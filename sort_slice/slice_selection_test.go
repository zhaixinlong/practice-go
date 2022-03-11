package sort_slice

import (
	"reflect"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{name: "t1", args: args{nums: []int{9, 2, 5, 4, 1, 2, 3, 4, 5}}, want: []int{1, 2, 2, 3, 4, 4, 5, 5, 9}},
		{name: "t1", args: args{nums: []int{1, 2, 3, 4, 5}}, want: []int{1, 2, 3, 4, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SelectionSort(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectionSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
