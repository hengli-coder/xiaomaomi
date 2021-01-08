package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {

	var arr []*ListNode
	for head != nil {
		arr = append(arr, head)
		head = head.Next
	}

	if len(arr) == 0 {
		return head
	}

	arr[0].Next = nil
	for i := 1; i < len(arr); i++ {
		arr[i].Next = arr[i-1]
	}

	return arr[len(arr)-1]
}

func reverse(head *ListNode) *ListNode {

	var preNode, curNode, nextNode *ListNode
	curNode = head

	for curNode != nil {

		nextNode = curNode.Next
		curNode.Next = preNode
		preNode = curNode
		curNode = nextNode
	}

	return preNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {

	var mPre, mNode, nNode, nNext *ListNode
	var preNode, curNode, NextNode *ListNode
	curNode = head
	for i := 1; i <= n; i++ {
		if i == m-1 {
			mPre = curNode
		}

		if i == m {
			mNode = curNode
		}

		if i == n {
			nNode = curNode
			nNext = nNode.Next
		}

		if i >= m && i <= n {
			NextNode = curNode.Next
			curNode.Next = preNode
			preNode = curNode
			curNode = NextNode
		} else {
			curNode = curNode.Next
		}

		if i > n {
			break
		}
	}

	if mPre == nil {
		head = nNode
	} else {
		mPre.Next = nNode
	}

	mNode.Next = nNext
	return head
}

func getDecimalValue(head *ListNode) int {

	var s int
	for head != nil {
		s = s<<1 | head.Val
		head = head.Next
	}
	return s
}

func removeDuplicateNodes(head *ListNode) *ListNode {
	var m = make(map[int]*ListNode)
	var pre *ListNode
	var cur = head

	for cur != nil {
		if _, exist := m[cur.Val]; exist {
			pre.Next = cur.Next
			cur = cur.Next
		} else {
			m[cur.Val] = cur
			pre = cur
			cur = cur.Next
		}

	}
	return head
}

func reorderList(head *ListNode) {
	var s []*ListNode
	cur := head
	for cur != nil {
		s = append(s, cur)
		cur = cur.Next
	}

	cur = head

	for i := 0; i < len(s); i++ {
		s[i].Next = s[len(s)-1-i]
	}
}

func main() {

}
