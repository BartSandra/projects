package main

import (
	"container/list"
	"fmt"
)

type TreeNode struct {
	HasToy bool
	Left   *TreeNode
	Right  *TreeNode
}

func UnrollGarland(root *TreeNode) (result []bool) {
	nodeQueue := list.New()
	nodeQueue.PushBack(root)
	var leftToRight = false

	for nodeQueue.Len() > 0 {
		size := nodeQueue.Len()
		tempQueue := list.New()

		for i := 0; i < size; i++ {
			if nodeQueue.Front().Value.(*TreeNode) == nil {
				nodeQueue.Remove(nodeQueue.Front())
				continue
			}
			result = append(result, nodeQueue.Front().Value.(*TreeNode).HasToy)

			if leftToRight {
				tempQueue.PushFront(nodeQueue.Front().Value.(*TreeNode).Left)
				tempQueue.PushFront(nodeQueue.Front().Value.(*TreeNode).Right)
			} else {
				tempQueue.PushFront(nodeQueue.Front().Value.(*TreeNode).Right)
				tempQueue.PushFront(nodeQueue.Front().Value.(*TreeNode).Left)
			}

			nodeQueue.Remove(nodeQueue.Front())
		}

		leftToRight = !leftToRight
		nodeQueue = tempQueue
	}

	return
}

func main() {
	root := &TreeNode{
		HasToy: true,
		Left: &TreeNode{
			HasToy: true,
			Left: &TreeNode{
				HasToy: true,
			},
			Right: &TreeNode{
				HasToy: false,
			},
		},
		Right: &TreeNode{
			HasToy: false,
			Left: &TreeNode{
				HasToy: true,
			},
			Right: &TreeNode{
				HasToy: true,
			},
		},
	}

	fmt.Println(UnrollGarland(root))

	/*root := TreeNode{HasToy: true}
	root.Left = &TreeNode{HasToy: true}
	root.Left.Left = &TreeNode{HasToy: true}
	root.Left.Right = &TreeNode{HasToy: false}
	root.Right = &TreeNode{HasToy: false}
	root.Right.Left = &TreeNode{HasToy: true}
	root.Right.Right = &TreeNode{HasToy: true}

	fmt.Println(UnrollGarland(&root))*/
}
