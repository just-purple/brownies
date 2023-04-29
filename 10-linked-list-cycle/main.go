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

	ll := &LinkedList{}
	ll.append(101)
	ll.append(102)
	ll.append(103)
	ll.append(104)

	// fmt.Printf("length is  %d \n", ll.length) //length is 6
	fmt.Printf("head is %d \n", ll.head.Val) //head is 101
	fmt.Printf("tail is %d \n", ll.tail.Val) //tail is 104
}

// Definition for singly-linked list
// Structure that represents each node in the linked list.
// Each node will have an integer value called Val and a second value called Next, which is a pointer to the next node
type ListNode struct {
	Val  int
	Next *ListNode
}

// The linked list struct have head and tail nodes pointer references of the linked list. It will also have a length value that keeps track of how long the linked list is.
type LinkedList struct {
	head *ListNode
	tail *ListNode
	// length int
}

// implementare una funzione che inserisce in coda usando il puntatore di coda (quindi non serve il for)
// The append method receives the integer Val that needs to be added into a new node to the end of the linked list.
func (list *LinkedList) append(data int) {
	//creates a new node that will be added to the linkedList
	// It first creates a ListNode from the value and assigns the pointer reference to the variable newNode.
	newNode := &ListNode{
		Val: data,
	}
	// The next step is to check if the current linked list has no head and,
	// if it does, it assigns the newNode created to the head and tail of the linked list.
	// If there is a head node present, the tail node’s next value will be assigned the newNode and the new node created earlier will be made the tail node.
	if list.head == nil {
		list.head = newNode
		list.tail = newNode
	} else {
		list.tail.Next = newNode
		list.tail = newNode
	}
	// Lastly increase the value of the list length in order to keep track of the linked list size.
	// list.length++
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

	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
