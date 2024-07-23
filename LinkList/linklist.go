package linklist

import "fmt"

type node struct {
	data string
	next *node
}

type linklist struct {
	head *node
	size int
}

type ILinklist interface {
	Len() int
	Insert(data string)
	Print()
}

func InitLinkList() ILinklist {
	return &linklist{
		head: nil,
		size: 0,
	}
}

func (l *linklist) Len() int { return (*l).size }

func (l *linklist) Insert(data string) {
	newNode := &node{data: data, next: nil}
	if (*l).head == nil {
		(*l).head = newNode
	} else {
		p := (*l).head
		newNode.next = p
		(*l).head = newNode
	}
	(*l).size += 1
}

func (l *linklist) Print() {
	p := (*l).head
	for p != nil {
		fmt.Printf("->%v", p.data)
		p = p.next
	}
	fmt.Println()
}
