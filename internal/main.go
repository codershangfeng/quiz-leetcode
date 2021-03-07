package internal

func twoSum(nums []int, target int) []int {
	hash := map[int]int{}
	for i := 0; i < len(nums); i++ {
		if _, ok := hash[target-nums[i]]; ok {
			return []int{hash[target-nums[i]], i}
		}
		hash[nums[i]] = i
	}

	return []int{}
}

type ListNode struct {
    Val int
    Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l := ListNode{}

	add(&l, l1, l2, false)

	return &l
}

func add(l *ListNode, l1 *ListNode, l2 *ListNode, offset bool) *ListNode {
	if !offset {
		if l1 == nil || l2 == nil {
			return l
		}
		l.Val = l1.Val + l2.Val
		
		
		next := ListNode{}
		if l1.Next != nil || l2.Next != nil {
			l.Next = &next
		}
		if l.Val < 10 {
			if l1.Next == nil && l2.Next == nil {
				return l
			}
			l.Next = &next
			add(l.Next, l1.Next, l2.Next, false)
		} else {
			l.Val %= 10
			add(l.Next, l1.Next, l2.Next, true)
		}
		return l
	}
	
	if l1 == nil && l2 == nil {
		l.Next = &ListNode{Val: 1, Next: nil}
		return l
	}
	
	next := ListNode{}
	if l1 == nil {
		l.Val = l2.Val + 1
		if l.Val >= 10 {
			l.Next = &next
			l.Val %= 10
			add(l.Next, nil, nil, true)
		}
		return l
	}

	if l2 == nil {
		l.Val = l1.Val + 1
		if l.Val >= 10 {
			l.Next = &next
			l.Val %= 10
			add(l.Next, nil, nil, true)
		}
		return l
	}
	
	l.Val = l1.Val + l2.Val + 1
	if l.Val >= 10 {
		add(l.Next, l1.Next, l2.Next, true)
	} else {
		add(l.Next, l1.Next, l2.Next, false)
	}
	return l
}

