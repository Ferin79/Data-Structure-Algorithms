package main

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func RemoveKthNodeFromEnd(head *LinkedList, k int) {
	first := head
    second := head
    counter := 1
    
    for counter <= k {
        counter++
        second = second.Next
    }

    if second == nil {
        first.Value = first.Next.Value
        first.Next = first.Next.Next
        return
    }

    for second.Next != nil {
        second = second.Next
        first = first.Next
    }
    first.Next = first.Next.Next
}
