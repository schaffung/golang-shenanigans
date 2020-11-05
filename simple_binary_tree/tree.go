package main

import "fmt"

type tree struct {
	value  int
	leftC  *tree
	rightC *tree
}

func addValueToTree(value int, treeNode **tree) {
	if *treeNode == nil {
		*treeNode = new(tree)
		(*treeNode).value = value
	} else if value < (*treeNode).value {
		addValueToTree(value, &(*treeNode).leftC)
	} else {
		addValueToTree(value, &(*treeNode).rightC)
	}
}

func preOrderTraversal(treeNode *tree) {
	if treeNode == nil {
		return
	}
	fmt.Print(treeNode.value, "->")
	if treeNode.leftC != nil {
		preOrderTraversal(treeNode.leftC)
	}
	if treeNode.rightC != nil {
		preOrderTraversal(treeNode.rightC)
	}
	return
}

func main() {
	var rootNode *tree
	addValueToTree(1, &rootNode)
	addValueToTree(2, &rootNode)
	addValueToTree(3, &rootNode)
	addValueToTree(5, &rootNode)

	preOrderTraversal(rootNode)
	fmt.Println()
}
