package avl

import (
	"reflect"
	"testing"
)

func TestTreeNode_Find(t *testing.T) {
	tree := &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val:   5,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 7},
		},
		Right: &TreeNode{
			Val:   15,
			Right: &TreeNode{Val: 20},
		},
	}
	tests := []struct {
		name string
		val  int
		want *TreeNode
	}{
		{
			name: "Case 1",
			val:  10,
			want: tree,
		},
		{
			name: "Case 2",
			val:  5,
			want: tree.Left,
		},
		{
			name: "Case 3",
			val:  15,
			want: tree.Right,
		},
		{
			name: "Case 4",
			val:  3,
			want: tree.Left.Left,
		},
		{
			name: "Case 5",
			val:  20,
			want: tree.Right.Right,
		},
		{
			name: "Case 6",
			val:  100,
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tree.Find(tt.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TreeNode.Find(%d) = %v, want %v", tt.val, got, tt.want)
			}
		})
	}
}

func TestTreeNode_Insert(t *testing.T) {
	tests := []struct {
		name     string
		initial  []int
		insert   int
		expected []int
	}{
		{
			name:     "Case 1",
			initial:  []int{},
			insert:   5,
			expected: []int{5},
		},
		{
			name:     "Case 2",
			initial:  []int{10},
			insert:   5,
			expected: []int{5, 10},
		},
		{
			name:     "Case 3",
			initial:  []int{10},
			insert:   15,
			expected: []int{10, 15},
		},
		{
			name:     "Case 4",
			initial:  []int{10},
			insert:   10,
			expected: []int{10, 10},
		},
		{
			name:     "Case 5",
			initial:  []int{10, 5, 15},
			insert:   12,
			expected: []int{5, 10, 12, 15},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var root *TreeNode
			for _, val := range tt.initial {
				root = root.Insert(val)
			}
			root = root.Insert(tt.insert)

			got := treeToSlice(root)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("TreeNode.Insert(%d) = %v, want %v", tt.insert, got, tt.expected)
			}
		})
	}
}

func treeToSlice(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	return append(append(treeToSlice(root.Left), root.Val), treeToSlice(root.Right)...)
}

func TestTreeNode_Erase(t *testing.T) {
	tests := []struct {
		name     string
		initial  []int
		erase    int
		expected []int
	}{
		{
			name:     "Erase leaf node",
			initial:  []int{5, 3, 6, 2, 4, 7},
			erase:    2,
			expected: []int{3, 4, 5, 6, 7},
		},
		{
			name:     "Erase node with one child",
			initial:  []int{5, 3, 6, 2, 4, 7},
			erase:    3,
			expected: []int{2, 4, 5, 6, 7},
		},
		{
			name:     "Erase node with two children",
			initial:  []int{5, 3, 6, 2, 4, 7},
			erase:    5,
			expected: []int{2, 3, 4, 6, 7},
		},
		{
			name:     "Erase root node to empty",
			initial:  []int{1},
			erase:    1,
			expected: []int{},
		},
		{
			name:     "Erase non-existing node",
			initial:  []int{5, 3, 6},
			erase:    10,
			expected: []int{3, 5, 6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var root *TreeNode
			for _, val := range tt.initial {
				root = root.Insert(val)
			}

			root = root.Erase(tt.erase)
			got := treeToSlice(root)

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("Erase(%d): got %v, want %v", tt.erase, got, tt.expected)
			}
		})
	}
}

func TestTreeNode_Traversals(t *testing.T) {
	var root *TreeNode
	for _, val := range []int{5, 3, 6, 2, 4, 7} {
		root = root.Insert(val)
	}

	tests := []struct {
		name     string
		method   func(*TreeNode) []int
		expected []int
	}{
		{
			name:     "Inorder",
			method:   (*TreeNode).Inorder,
			expected: []int{2, 3, 4, 5, 6, 7},
		},
		{
			name:     "Preorder",
			method:   (*TreeNode).Preorder,
			expected: []int{5, 3, 2, 4, 6, 7},
		},
		{
			name:     "Postorder",
			method:   (*TreeNode).Postorder,
			expected: []int{2, 4, 3, 7, 6, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.method(root)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("%s() = %v, want %v", tt.name, result, tt.expected)
			}
		})
	}
}

func TestTreeNode_Destroy(t *testing.T) {
	tests := []struct {
		name    string
		initial []int
	}{
		{
			name:    "Destroy full tree",
			initial: []int{10, 5, 15, 3, 7, 12, 18},
		},
		{
			name:    "Destroy single node",
			initial: []int{42},
		},
		{
			name:    "Destroy empty tree",
			initial: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var root *TreeNode
			for _, val := range tt.initial {
				root = root.Insert(val)
			}

			root.Destroy()
			if root != nil && (root.Left != nil || root.Right != nil) {
				t.Errorf("Expected tree to be destroyed (Left and Right nil), got Left: %v, Right: %v", root.Left, root.Right)
			}
		})
	}
}
