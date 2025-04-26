package avl

type TreeNode struct {
	Val    int
	Left   *TreeNode
	Right  *TreeNode
	height int
}

func New(val int, left *TreeNode, right *TreeNode) *TreeNode {
	return &TreeNode{Val: val, Left: left, Right: right, height: 1}
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
	} else if root.Val < val {
		root.Right = root.Right.Insert(val)
	} else {
		return root
	}

	root.updateHeight()

	hl := root.Left.Height()
	hr := root.Right.Height()

	if hr-hl == 2 {
		if root.Right.Left.Height() <= root.Right.Right.Height() {
			root = root.smallLeftRotate()
		} else {
			root = root.bigLeftRotate()
		}
	}
	if hl-hr == 2 {
		if root.Left.Right.Height() <= root.Left.Left.Height() {
			root = root.smallRightRotate()
		} else {
			root = root.bigRightRotate()
		}
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

	root.updateHeight()

	hl := root.Left.Height()
	hr := root.Right.Height()

	if hr-hl == 2 {
		if root.Right.Left.Height() <= root.Right.Right.Height() {
			root = root.smallLeftRotate()
		} else {
			root = root.bigLeftRotate()
		}
	}
	if hl-hr == 2 {
		if root.Left.Right.Height() <= root.Left.Left.Height() {
			root = root.smallRightRotate()
		} else {
			root = root.bigRightRotate()
		}
	}
	return root
}

func (root *TreeNode) Height() int {
	if root == nil {
		return 0
	}
	return root.height
}

func (root *TreeNode) updateHeight() {
	root.height = 1 + max(root.Left.Height(), root.Right.Height())
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

//				    root                          pivot
//			       /    \                        /    \
//		        pivot    C      =>              A     root
//			   /   \                                  /  \
//	          A     B                                B    C
func (root *TreeNode) smallRightRotate() *TreeNode {
	pivot := root.Left
	root.Left = pivot.Right
	pivot.Right = root

	root.updateHeight()
	pivot.updateHeight()
	return pivot
}

//			       root                           pivot
//			      /    \                         /    \
//		         A     pivot      =>          root     C
//	            /        \                   /   \
//			   B          C                 A     B
func (root *TreeNode) smallLeftRotate() *TreeNode {
	pivot := root.Right
	root.Right = pivot.Left
	pivot.Left = root

	pivot.updateHeight()
	root.updateHeight()
	return pivot
}

func (root *TreeNode) bigLeftRotate() *TreeNode {
	root.Left = root.smallRightRotate()
	return root.smallLeftRotate()
}

func (root *TreeNode) bigRightRotate() *TreeNode {
	root.Left = root.smallLeftRotate()
	return root.smallRightRotate()
}
