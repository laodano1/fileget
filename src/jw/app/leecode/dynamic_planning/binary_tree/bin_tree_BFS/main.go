package main

import (
	"fileget/util"
)

type TreeNode2 struct {
	Val      int32
	Children [2]*TreeNode2
}

func buildBinaryTree(nums []int) {

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
		if i == 0 {
			continue
		}
		setNode(input[i], root)
	}
	//util.Lg.Debugf("%v", root)

	return root
}

//func DoBFSearch(node *TreeNode2, nodeQue []*TreeNode2, target, height int32) int32 {
//	if node.Val == target {
//		return height
//	}
//
//	if node.Children[0] != nil {
//		nodeQue = append(nodeQue, node.Children[0])
//	}
//	if node.Children[1] != nil {
//		nodeQue = append(nodeQue, node.Children[1])
//	}
//
//	height++
//	for i := range nodeQue {
//		DoBFSearch()
//	}
//
//}

func BFS(root *TreeNode2, target int32) int32 {
	nodeQue := make([]*TreeNode2, 0)
	height := int32(0)
	visited := make(map[*TreeNode2]bool)

	nodeQue = append(nodeQue, root)

	for len(nodeQue) > 0 {
		//sz := len(nodeQue)
		for i := range nodeQue {
			cur := nodeQue[i]
			if cur == nil {
				continue
			}
			if cur.Val == target {
				return height
			}

			if _, ok := visited[cur]; !ok {
				nodeQue = append(nodeQue, cur.Children[0])
				nodeQue = nodeQue[1:]
			}
		}

		height++
	}

	//if root.Val == target {
	//	return height
	//} else {
	//return DoBFSearch(root, nodeQue, target, height)
	//}
	return -1

}

func main() {
	util.Lg.Debugf("this is for binary tree basic traverse ways")
	nums := []int32{5, 2, 4, 1, 3, 8, 6, 7, 9}
	//nums := []int32{1, 2, 3, 4, 5}
	//nums := []int32{5, 4, 3, 2, 1}
	mt := geneBSTree(nums)

	util.Lg.Debugf("height: %v", BFS(mt, 7))

}
