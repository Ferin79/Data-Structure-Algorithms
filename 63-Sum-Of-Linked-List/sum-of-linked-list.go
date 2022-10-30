package main

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func SumOfLinkedLists(linkedListOne *LinkedList, linkedListTwo *LinkedList) *LinkedList {    
        
    head := &LinkedList{Value:0}
    currentNode := head
    carry := 0
    
    for linkedListOne != nil || linkedListTwo != nil || carry != 0 {
        var valueOne, valueTwo int

        if linkedListOne != nil {
            valueOne = linkedListOne.Value
            linkedListOne = linkedListOne.Next
        }
        if linkedListTwo != nil {
            valueTwo = linkedListTwo.Value
            linkedListTwo = linkedListTwo.Next
        }
        addition := valueOne + valueTwo + carry
        carry = addition / 10
        addition = addition % 10
       
        node := &LinkedList{Value: addition}

		currentNode.Next = node
		currentNode = currentNode.Next
    
    }    
    return head.Next
}
