package list_node_slice

import (
	"reflect"
	"testing"

	"github.com/zhaixinlong/practice-go/model"
)

func TestBubbleSort(t *testing.T) {
	type args struct {
		head *model.ListNode
	}
	tests := []struct {
		name string
		args args
		want *model.ListNode
	}{
		// TODO: Add test cases.
		{name: "t1", args: args{head: model.GenerateListNode([]int{9, 2, 5, 4, 1, 2, 3, 4, 5})}, want: model.GenerateListNode([]int{1, 2, 2, 3, 4, 4, 5, 5, 9})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BubbleSort(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BubbleSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
