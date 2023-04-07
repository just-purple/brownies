package main

import (
	"fmt"
)

func main() {

	sl := &ListNode{}

	for i := 0; i < 5; i++ {
		sl.Insert(i)
	}

	Print(sl)

	// creating a loop in the above linked list
	sl.Next.Next.Next.Next.Next = sl // comment this then for no loop

	if hasCycle(sl) {
		fmt.Println("found loop")
	} else {
		fmt.Println("no loop")
	}
}

// Definition for singly-linked list
// node e linked list struct
type ListNode struct {
	Val  int
	Next *ListNode
}

// inserimento, push to a linked list
func (l *ListNode) Insert(Val int) {
	list := &ListNode{Val: Val, Next: nil}
	if l == nil {
		l = list
	} else {
		p := l
		for p.Next != nil {
			p = p.Next
		}
		p.Next = list
	}
}

// stampa della lista
func Print(l *ListNode) {
	p := l
	for p != nil {
		fmt.Printf("-> %v ", p.Val)
		p = p.Next
	}
}

// Detect loop in linkedlist using floyd's cycle finding algorithm
func hasCycle(head *ListNode) bool {
	// se la lista è vuota o composta da un solo nodo non connesso con se stesso
	if head == nil || head.Next == nil {
		return false
	}

	slow := head
	fast := head

	// dato il controllo che faccio all'inizio posso togliere la prima condizione?
	for slow != nil && fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
