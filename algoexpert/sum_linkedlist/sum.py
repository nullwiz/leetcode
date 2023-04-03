# This is an input class. Do not edit.
class LinkedList:
    def __init__(self, value):
        self.value = value
        self.next = None


def sumOfLinkedLists(linkedListOne, linkedListTwo):
    # Write your code here.
    sum = int(returnIntOfList(linkedListOne)) + int(returnIntOfList(linkedListTwo))
    return createLinkedList(sum)
    

def createLinkedList(values):
    values = [int(i) for i in reversed(str(values))]
    newLinkedList = LinkedList(0) 
    head = newLinkedList
    for i in range(0,len(values)):
        head.value = values[i]
        if i < len(values) - 1:
            newNode = LinkedList(0)
            head.next = newNode
            head = head.next
    return newLinkedList 
    
def returnIntOfList(linkedListOne):
    a = ""
    while linkedListOne != None:
        a += str(linkedListOne.value)
        linkedListOne = linkedListOne.next
    a=a[::-1]
    return int(a) 

    
