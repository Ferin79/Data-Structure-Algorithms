package main
import "fmt"

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func RemoveKthNodeFromEnd(head *LinkedList, k int) {
	length := 0

    node := head

    for node != nil {
        length++
        node = node.Next
    }

    nodeToRemove := length - k + 1


    if nodeToRemove == 1 {
        head.Value = head.Next.Value
        head.Next = head.Next.Next
        fmt.Println("")
    } else {
        currPos := 1

        previous := head
        for head != nil && currPos != nodeToRemove {
            currPos++
            previous = head
            head = head.Next
        }
       

        if head != nil && head.Next != nil {
            previous.Next = head.Next
        } else {
            previous.Next = nil
        }
    }
 

}
