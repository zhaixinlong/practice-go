package main

import (
	"reflect"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		list1 *ListNode
		list2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		// l1 l2 都有 并且 长度一致
		{name: "Test_mergeTwoLists_1", args: args{list1: GenerateListNode([]int{1, 2}), list2: GenerateListNode([]int{3, 4})}, want: GenerateListNode([]int{1, 2, 3, 4})},
		// l1 为nil l2 不为nil
		{name: "Test_mergeTwoLists_2", args: args{list1: nil, list2: GenerateListNode([]int{3, 4})}, want: GenerateListNode([]int{3, 4})},
		// l1 不为nil l2 为nil
		{name: "Test_mergeTwoLists_3", args: args{list1: GenerateListNode([]int{1, 2}), list2: nil}, want: GenerateListNode([]int{1, 2})},
		// l1 l2 都为nil
		{name: "Test_mergeTwoLists_4", args: args{list1: nil, list2: nil}, want: nil},
		// l1 l2 都有 并且 l1 长
		{name: "Test_mergeTwoLists_5", args: args{list1: GenerateListNode([]int{1, 2, 3}), list2: GenerateListNode([]int{3, 4})}, want: GenerateListNode([]int{1, 2, 3, 3, 4})},
		// l1 l2 都有 并且 l2 长
		{name: "Test_mergeTwoLists_6", args: args{list1: GenerateListNode([]int{1, 2}), list2: GenerateListNode([]int{2, 3, 4})}, want: GenerateListNode([]int{1, 2, 2, 3, 4})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoLists(tt.args.list1, tt.args.list2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mergeTwoListsTraverse(t *testing.T) {
	type args struct {
		list1 *ListNode
		list2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		// l1 l2 都有 并且 长度一致
		{name: "Test_mergeTwoListsTraverse_1", args: args{list1: GenerateListNode([]int{1, 2}), list2: GenerateListNode([]int{3, 4})}, want: GenerateListNode([]int{1, 2, 3, 4})},
		// l1 为nil l2 不为nil
		{name: "Test_mergeTwoListsTraverse_2", args: args{list1: nil, list2: GenerateListNode([]int{3, 4})}, want: GenerateListNode([]int{3, 4})},
		// l1 不为nil l2 为nil
		{name: "Test_mergeTwoListsTraverse_3", args: args{list1: GenerateListNode([]int{1, 2}), list2: nil}, want: GenerateListNode([]int{1, 2})},
		// l1 l2 都为nil
		{name: "Test_mergeTwoListsTraverse_4", args: args{list1: nil, list2: nil}, want: nil},
		// l1 l2 都有 并且 l1 长
		{name: "Test_mergeTwoListsTraverse_5", args: args{list1: GenerateListNode([]int{1, 2, 3}), list2: GenerateListNode([]int{3, 4})}, want: GenerateListNode([]int{1, 2, 3, 3, 4})},
		// l1 l2 都有 并且 l2 长
		{name: "Test_mergeTwoListsTraverse_6", args: args{list1: GenerateListNode([]int{1, 2}), list2: GenerateListNode([]int{2, 3, 4})}, want: GenerateListNode([]int{1, 2, 2, 3, 4})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoListsTraverse(tt.args.list1, tt.args.list2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoListsTraverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerateListNode(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{name: "TestGenerateListNode1", args: args{arr: []int{1, 2}}, want: GenerateListNode([]int{1, 2})},
		{name: "TestGenerateListNode2", args: args{arr: []int{}}, want: GenerateListNode([]int{})},
		{name: "TestGenerateListNode3", args: args{arr: nil}, want: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateListNode(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateListNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetListNode1(t *testing.T) {
	tests := []struct {
		name string
		want *ListNode
	}{
		// TODO: Add test cases.
		{name: "TestGetListNode1", want: GenerateListNode([]int{1, 2, 3, 4})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetListNode1(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetListNode1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetListNode2(t *testing.T) {
	tests := []struct {
		name string
		want *ListNode
	}{
		// TODO: Add test cases.
		{name: "TestGetListNode2Succ", want: GenerateListNode([]int{1, 3, 4, 5})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetListNode2(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetListNode2() = %v, want %v", got, tt.want)
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
