package main


 func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
   var newList  = &ListNode{}
    var out = newList
    for l1 != nil && l2 != nil {
        if l1.Val < l2.Val {
            newList.Next = l1 //make the next element the current element of l1
            l1 = l1.Next //move l1 to the next Node
            newList = newList.Next //move newList to the most current node 
        }else{
            newList.Next = l2
            l2 = l2.Next
            newList = newList.Next
        }
    }
    if l1 != nil {
        newList.Next = l1
    }else if l2 != nil {
        newList.Next = l2 //append the rest of l2 if l1 is empty
    }
    
    return out.Next
}

