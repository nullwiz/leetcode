# This is an input class. Do not edit.
class LinkedList:
    def __init__(self, value):
        self.value = value
        self.next = None


def mergingLinkedLists(linkedListOne, linkedListTwo):
    # If lengths are the same
    node1, node2 = linkedListOne, linkedListTwo
    len1, len2 = length(node1), length(node2)
    if len1 != len2:
        node1, node2 = matchLengths(node1,len1,node2,len2)
    while node1!=None and node2!=None:
        if node1.value == node2.value:
            return node1 
        node1, node2 = node1.next, node2.next
    return None
        
def length(linkedList):
    sum = 0 
    head = linkedList
    while head!=None:
        sum+=1
        head = head.next
    return sum 

def matchLengths(l1,len1,l2,len2):
    # match lengths: biggest linked list should get
    # sliced by the difference
    # returns len1, len2
    if len1 > len2:
        head = l1
        for i in range(len1-len2):
            head = head.next
        return head, l2
    if len2 > len1:
        head = l2
        for i in range(len2-len1):
            head = head.next
        return l1,head
