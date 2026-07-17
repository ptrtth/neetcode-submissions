/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
	visited := map[*ListNode]struct{}{}

	for node := head; node != nil; node = node.Next {
		if _, exists := visited[node]; exists {
			return true
		}

		visited[node] = struct{}{}
	}

	return false
}
