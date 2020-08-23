package main

import (
	"fileget/util"
	"math"
)

/*
	breadth first search / 广度优先
*/

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

func BFS(root *TreeNode2, target int32, step *int32) int32 {
	if root == nil {
		return -1
	}
	nodeQue := make([]*TreeNode2, 0)     // 节点队列
	height := int32(1)                   // 树的层数记录
	visited := make(map[*TreeNode2]bool) // 记录遍历过的节点

	nodeQue = append(nodeQue, root) // root 入队
	visited[root] = true            // 记录root 遍历过
	//var step int32

	for len(nodeQue) > 0 {
		sz := len(nodeQue)
		for i := 0; i < sz; i++ {
			*step++
			cur := nodeQue[0]
			if cur == nil {
				continue
			}
			if cur.Val == target {
				return height
			}

			if cur.Children[0] != nil {
				if _, ok := visited[cur.Children[0]]; !ok {
					nodeQue = append(nodeQue, cur.Children[0])
				}
			}
			if cur.Children[1] != nil {
				if _, ok := visited[cur.Children[1]]; !ok {
					nodeQue = append(nodeQue, cur.Children[1])
				}
			}
			nodeQue = nodeQue[1:]
		}

		height++
	}

	util.Lg.Debugf("step: %v", step)

	return -1
}

func main() {
	util.Lg.Debugf("this is for binary tree basic traverse ways")
	nums := []int32{5, 2, 4, 1, 3, 8, 6, 7, 9}
	//nums := []int32{1, 2, 3, 4, 5}
	//nums := []int32{5, 4, 3, 2, 1}
	mt := geneBSTree(nums)

	var step int32

	util.Lg.Debugf("height: %v, step: %v", BFS(mt, 5, &step), step)
	util.Lg.Debugf("height: %v, step: %v", BFS(mt, 8, &step), step)
	util.Lg.Debugf("height: %v, step: %v", BFS(mt, 9, &step), step)
	util.Lg.Debugf("height: %v, step: %v", BFS(mt, 7, &step), step)

	util.Lg.Debugf("tmp: %v", math.Ceil(7.08))

}
