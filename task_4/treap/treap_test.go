package main

import (
	"reflect"
	"testing"
)

func TestTreap_BuildTreap(t *testing.T) {
	type args struct {
		nodes []Node
	}
	tests := []struct {
		name       string
		args       args
		wantParent []int
		wantLeft   []int
		wantRight  []int
	}{
		{
			name: "Case 1",
			args: args{
				nodes: []Node{
					{A: 5, B: 4, Idx: 1},
					{A: 2, B: 2, Idx: 2},
					{A: 3, B: 9, Idx: 3},
					{A: 0, B: 5, Idx: 4},
					{A: 1, B: 3, Idx: 5},
					{A: 6, B: 6, Idx: 6},
					{A: 4, B: 11, Idx: 7},
				},
			},
			wantParent: []int{2, 0, 1, 5, 2, 1, 3},
			wantLeft:   []int{3, 5, 0, 0, 4, 0, 0},
			wantRight:  []int{6, 1, 7, 0, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotParent, gotLeft, gotRight := BuildTreap(tt.args.nodes)
			if !reflect.DeepEqual(gotParent, tt.wantParent) {
				t.Errorf("BuildTreap() gotParent = %v, want %v", gotParent, tt.wantParent)
			}
			if !reflect.DeepEqual(gotLeft, tt.wantLeft) {
				t.Errorf("BuildTreap() gotLeft = %v, want %v", gotLeft, tt.wantLeft)
			}
			if !reflect.DeepEqual(gotRight, tt.wantRight) {
				t.Errorf("BuildTreap() gotRight = %v, want %v", gotRight, tt.wantRight)
			}
		})
	}
}
