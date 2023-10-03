/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 import (

 "fmt"
//  "math"
 )

func middleNode(head *ListNode) *ListNode {
    n  := head 
    nodeLen := 1
    for n != nil {
        nodeLen = nodeLen+1
        n = n.Next
    }

    quot := int(math.Ceil(float64(nodeLen)/2))

    var result *ListNode = head
    nodeLen = 1
    for result != nil {
        if nodeLen >= quot{
            break
        }
        result = result.Next
        nodeLen = nodeLen+1
    }

    return result
}