package main

import "fmt"

type node struct {
	number int
	next   *node
}

type linkedList struct {
	head *node
	size int
}

var node_count int

func (l *linkedList) search(number int) bool {
	head := l.head

	if head == nil {
		return false
	}

	if head.number == 0 {
		return false
	}

	for head.number != number && head.next != nil {
		node_count++
		fmt.Println("Node numero", node_count)

		head = head.next
	}

	if head.next == nil {
		return false
	}

	node_count = 0
	fmt.Println(head)
	return true
}

func (l *linkedList) prepend(number int) {
	second := l.head
	l.head = &node{number, second}

	l.size++
}

func (l *linkedList) remove(number int) {
	head := l.head

	if head.number == number {
		head = head.next
		return
	}

	for head.next.number != number && head.next.next != nil {
		head = head.next
	}

	if head.next.next != nil {
		head.next = head.next.next
		l.size--
	}
}

func main() {
	list := &linkedList{}
	fmt.Println(list)

	list.prepend(13)
	list.prepend(25)
	list.prepend(2)
	list.prepend(67)
	list.prepend(83)
	list.prepend(15)
	list.prepend(118)
	list.prepend(1320)

	fmt.Println(list, list.head, list.head.next, list.head.next.next, list.head.next.next.next)

	list.remove(118)

	fmt.Println(list, list.head, list.head.next, list.head.next.next, list.head.next.next.next)

	fmt.Println(list.search(99))
}
