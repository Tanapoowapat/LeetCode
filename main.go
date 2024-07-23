package main

import (
	"fmt"

	linklist "github.com/Tanapoowapat/LeetCode/LinkList"
)

func main() {
	// s := stack.InitStack()
	// s.Push("Hello")
	// s.Push(",")
	// s.Push("World")
	// fmt.Println(s)

	// q := queue.InitQueue(5)
	// q.Print()

	l := linklist.InitLinkList()
	l.Insert("1")
	l.Insert("2")
	l.Insert("3")
	l.Insert("4")

	l.Print()
	fmt.Println(l.Len())

}
