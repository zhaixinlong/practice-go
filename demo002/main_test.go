package main

import (
	"reflect"
	"testing"
)

func Test_returnNewSlice(t *testing.T) {
	type args struct {
		slice []int
	}
	tests := []struct {
		name  string
		args  args
		want  []int
		want1 int
	}{
		// TODO: Add test cases.
		{name: "test1", args: args{slice: []int{0, 1, 1, 2, 3, 3, 3, 3, 3, 3, 3, 4, 5}}, want: []int{0, 1, 2, 3, 4, 5}, want1: 13},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := returnNewSlice(tt.args.slice)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("returnNewSlice() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("returnNewSlice() got1 = %v, want %v", got1, tt.want1)
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
