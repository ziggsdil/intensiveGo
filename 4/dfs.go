package main

import "fmt"

type Node struct {
	Val  int
	R, L *Node
}

// посчитать сумму всех элементов в дереве
func main() {
	root := &Node{
		Val: 3,
		R: &Node{
			Val: 2,
			R:   nil,
			L: &Node{
				Val: 1,
				R: &Node{
					Val: 12,
					R:   nil,
					L:   nil,
				},
				L: &Node{
					Val: 10,
					R:   nil,
					L:   nil,
				},
			},
		},
		L: &Node{
			Val: 4,
			L:   nil,
			R:   nil,
		},
	}

	res := dfs(root)
	fmt.Println(res)

}

func dfs(node *Node) int {
	if node == nil {
		return 0
	}

	leftSum := dfs(node.L)
	rightSum := dfs(node.R)

	return node.Val + leftSum + rightSum
}
