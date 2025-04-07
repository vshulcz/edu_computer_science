package bst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func New(val int, left *TreeNode, right *TreeNode) *TreeNode {
	return &TreeNode{Val: val, Left: left, Right: right}
}

func (root *TreeNode) Destroy() {
	if root == nil {
		return
	}
	root.Left.Destroy()
	root.Right.Destroy()
	root.Left = nil
	root.Right = nil
}

func (root *TreeNode) Find(val int) *TreeNode {
	if root == nil || val == root.Val {
		return root
	}

	if val < root.Val {
		return root.Left.Find(val)
	}
	return root.Right.Find(val)
}

func (root *TreeNode) Insert(val int) *TreeNode {
	if root == nil {
		return &TreeNode{
			Val: val,
		}
	}

	if root.Val > val {
		root.Left = root.Left.Insert(val)
	} else {
		root.Right = root.Right.Insert(val)
	}
	return root
}

func (root *TreeNode) Erase(key int) *TreeNode {
	if root == nil {
		return nil
	}

	if key < root.Val {
		root.Left = root.Left.Erase(key)
	} else if key > root.Val {
		root.Right = root.Right.Erase(key)
	} else {
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		minNode := root.Right
		for minNode.Left != nil {
			minNode = minNode.Left
		}
		root.Val = minNode.Val
		root.Right = root.Right.Erase(minNode.Val)
	}
	return root
}

// Inorder: Left -> Root -> Right
func (root *TreeNode) Inorder() []int {
	if root == nil {
		return []int{}
	}
	result := []int{}
	result = append(result, root.Left.Inorder()...)
	result = append(result, root.Val)
	result = append(result, root.Right.Inorder()...)
	return result
}

// Preorder: Root -> Left -> Right
func (root *TreeNode) Preorder() []int {
	if root == nil {
		return []int{}
	}
	result := []int{}
	result = append(result, root.Val)
	result = append(result, root.Left.Preorder()...)
	result = append(result, root.Right.Preorder()...)
	return result
}

// Postorder: Left -> Right -> Root
func (root *TreeNode) Postorder() []int {
	if root == nil {
		return []int{}
	}
	result := []int{}
	result = append(result, root.Left.Postorder()...)
	result = append(result, root.Right.Postorder()...)
	result = append(result, root.Val)
	return result
}
