# This is an input class. Do not edit.
class LinkedList:
    def __init__(self, value):
        self.value = value
        self.next = None


def reverseLinkedList(head):
    # Write your code here.
    current = head
    previous, nextp = None, None 
    while current != None: 
        nextp = current.next
        current.next = previous
        previous = current
        current = nextp
    return previous 
