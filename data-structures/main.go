package main

import (
	"data-structures/queue"
	bst "data-structures/trees"
	"fmt"
)

func main() {
	bst := bst.NewBinarySearchTree()

	bst.AddNode(8)
	bst.AddNode(3)
	bst.AddNode(1)
	bst.AddNode(6)
	bst.AddNode(4)
	bst.AddNode(10)
	bst.AddNode(14)
	bst.AddNode(13)
	bst.AddNode(7)

	fmt.Printf("       %d\n", bst.Root.Value)

	fmt.Println("     /   \\ ")
	fmt.Printf("    %d", bst.Root.Left.Value)
	fmt.Printf("     %d\n", bst.Root.Right.Value)

	fmt.Println("   / \\     \\ ")
	fmt.Printf("  %d", bst.Root.Left.Left.Value)
	fmt.Printf("   %d", bst.Root.Left.Right.Value)

	fmt.Printf("    %d\n", bst.Root.Right.Right.Value)

	fmt.Println("     / \\   /")
	fmt.Printf("    %d", bst.Root.Left.Right.Left.Value)
	fmt.Printf("   %d", bst.Root.Left.Right.Right.Value)

	fmt.Printf("  %d", bst.Root.Right.Right.Left.Value)

	queue := queue.NewQueue()

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	queue.Enqueue(5)

	fmt.Println()
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())

	value, err := queue.Dequeue()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(value)
}