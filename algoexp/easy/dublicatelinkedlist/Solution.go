package main

import "fmt"

// LinkedList This is an input struct. Do not edit.
type LinkedList struct {
	Value int
	Next  *LinkedList
}

// RemoveDuplicatesFromLinkedList O(n) time | O(1) space
func RemoveDuplicatesFromLinkedList(linkedList *LinkedList) *LinkedList {
	if linkedList.Next != nil {
		if linkedList.Value == linkedList.Next.Value {
			linkedList.Next = linkedList.Next.Next
			RemoveDuplicatesFromLinkedList(linkedList)
		} else {
			RemoveDuplicatesFromLinkedList(linkedList.Next)
		}
		// Write your code here.
	}
	return linkedList

}

func main() {
	list := LinkedList{
		Value: 1, Next: &LinkedList{
			Value: 1, Next: &LinkedList{
				Value: 1, Next: &LinkedList{
					Value: 2, Next: &LinkedList{
						Value: 3, Next: &LinkedList{
							Value: 4, Next: &LinkedList{
								Value: 4, Next: &LinkedList{},
							},
						},
					},
				},
			},
		},
	}
	RemoveDuplicatesFromLinkedList(&list)
	fmt.Println("end")
}
