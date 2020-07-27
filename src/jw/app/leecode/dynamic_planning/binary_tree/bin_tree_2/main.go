package main

import "fileget/util"

//
type TreeNode2 struct {
	Val      int32
	Children [2]*TreeNode2
}

func setNode(val int32, node *TreeNode2) {
	//tmp := 0
	if node.Val > val {
		if node.Children[0] == nil {
			node.Children[0] = &TreeNode2{Val: val}
		} else {
			setNode(val, node.Children[0])
		}
	} else {
		if node.Children[1] == nil {
			node.Children[1] = &TreeNode2{Val: val}
		} else {
			setNode(val, node.Children[1])
		}
	}
	//util.Lg.Debugf("%v", tmp)
}

func geneBSTree(input []int32) *TreeNode2 {
	root := &TreeNode2{Val: input[0]}

	for i := range input {
		if i == 0 {continue}
		setNode(input[i], root)
	}
	//util.Lg.Debugf("%v", root)

	return root
}

func getNodeByPreorder(n *TreeNode2) {
	//util.Lg.Debugf("%v", n.Val)
	arr = append(arr, n.Val)
	if n.Children[0] != nil {
		getNodeByPreorder(n.Children[0])
	}

	if n.Children[1] != nil {
		getNodeByPreorder(n.Children[1])
	}
}

var arr = make([]int32, 0)
func Preorder(root *TreeNode2) {
	getNodeByPreorder(root)
}

func getNodeByInorder(n *TreeNode2) {
	//util.Lg.Debugf("%v", n.Val)
	if n.Children[0] != nil {
		getNodeByInorder(n.Children[0])
	}

	arr = append(arr, n.Val)

	if n.Children[1] != nil {
		getNodeByInorder(n.Children[1])
	}
}

func Inorder(root *TreeNode2) {
	getNodeByInorder(root)
}

func getNodeByPostorder(n *TreeNode2) {
	//util.Lg.Debugf("%v", n.Val)
	if n.Children[0] != nil {
		getNodeByPostorder(n.Children[0])
	}

	if n.Children[1] != nil {
		getNodeByPostorder(n.Children[1])
	}
	arr = append(arr, n.Val)
}

func Postorder(root *TreeNode2) {
	getNodeByPostorder(root)
}

func main() {
	nums := []int32{5, 2, 4,  1, 3, 8, 6, 7, 9}
	//nums := []int32{1, 2, 3, 4, 5}
	//nums := []int32{5, 4, 3, 2, 1}
	mt := geneBSTree(nums)

	// preorder
	Preorder(mt)
	util.Lg.Debugf("%-9v: %v", "Preorder", arr)

	// inorder
	arr = make([]int32, 0)
	Inorder(mt)
	util.Lg.Debugf("%-9v: %v", "Inorder", arr)

	// postorder
	arr = make([]int32, 0)
	Postorder(mt)
	util.Lg.Debugf("%-9v: %v", "Postorder", arr)

}
