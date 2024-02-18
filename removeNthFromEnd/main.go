package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	node := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, &ListNode{7, &ListNode{8, &ListNode{9, &ListNode{10, nil}}}}}}}}}}
	modifiedNode := removeNthFromEnd(node, 7)

	print(modifiedNode)
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	pointer1 := head
	pointer2 := head

	for i := 0; i < n; i++ {
		pointer1 = pointer1.Next
	}

	if pointer1 == nil { // Means remove the first one
		return head.Next
	}

	for pointer1.Next != nil {
		pointer1 = pointer1.Next
		pointer2 = pointer2.Next
	}

	pointer2.Next = pointer2.Next.Next

	return head
}
