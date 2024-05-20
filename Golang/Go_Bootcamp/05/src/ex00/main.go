//Дерево считается сбалансированным, если количество игрушек в левом поддереве равно количеству игрушек в правом поддереве

package main

import "fmt"

type TreeNode struct {
	HasToy bool
	Left   *TreeNode
	Right  *TreeNode
}

func AreToysBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	leftCounter, rightCounter := 0, 0

	countToys(root.Left, &leftCounter)
	countToys(root.Right, &rightCounter)

	return leftCounter == rightCounter
}

func countToys(root *TreeNode, counter *int) {

	if root == nil {
		return
	}

	if root.HasToy {
		*counter++
	}

	countToys(root.Left, counter)
	countToys(root.Right, counter)
}

func main() {

	t := TreeNode{HasToy: true}
	t.Left = new(TreeNode)
	t.Left.HasToy = true
	t.Right = new(TreeNode)
	t.Right.HasToy = true
	fmt.Println(AreToysBalanced(&t))

	/*root := TreeNode{HasToy: true}
	root.Left = &TreeNode{HasToy: true}
	root.Right = &TreeNode{HasToy: false}
	fmt.Println(AreToysBalanced(&root))*/

	/*root := &TreeNode{HasToy: false}
	root.Left = &TreeNode{HasToy: true}
	root.Right = &TreeNode{HasToy: false}
	root.Right.Right = &TreeNode{HasToy: true}
	root.Left.Right = &TreeNode{HasToy: true}
	fmt.Println(AreToysBalanced(root))*/

}
