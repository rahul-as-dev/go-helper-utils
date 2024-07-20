package main

import (
	"fmt"

	"github.com/rahul-as-dev/go-helper-utils/cmd/utils"
)

func main() {
	fmt.Printf("Hello World\n")
	//utils.ConvertIntToBits()
	//utils.ConvertBitsToInt()
	tree := utils.NewBinaryTree(10)
	tree.Left = utils.NewBinaryTree(5)
	tree.Right = utils.NewBinaryTree(15)
	fmt.Println(tree.NodeCount())
}
