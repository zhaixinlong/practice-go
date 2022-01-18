package main

import (
	"reflect"
	"testing"
)

func Test_sortBubble(t *testing.T) {
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
			if got := sortBubble(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortBubble() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "main"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
