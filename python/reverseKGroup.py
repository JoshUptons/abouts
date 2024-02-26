from typing import Optional


class ListNode:

    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

    def print(self):
        vals = []
        while self != None:
            vals.append(self.val)
            self = self.next

        print(vals)


head = ListNode(1, ListNode(2, ListNode(3, ListNode(4, ListNode(5, None)))))


def reverseKGroup(head: Optional[ListNode], k: int) -> Optional[ListNode]:
    if head == None:
        return None

    # memory array to store the current values to swap
    # we are going to have to use 2 pointers to swap the items
    # here as well inside the nxt loop
    m = []
    cur = head
    nxt = head

    i = 0
    while i < k:
        if nxt == None:
            break
        m.append(nxt.val)
        nxt = nxt.next
        i += 1

    while cur != None:

        # initialize the memory array pointers
        p1 = 0
        p2 = len(m) - 1

        # reverse the memory array
        while p1 < p2:
            m[p1], m[p2] = m[p2], m[p1]
            p1 += 1
            p2 -= 1

        while len(m) > 0 and cur != None:
            cur.val = m[0]
            m = m[1:]
            cur = cur.next

        cur = nxt

        i = 0
        while i < k:
            if nxt == None:
                return head
            m.append(nxt.val)
            nxt = nxt.next
            i += 1

    return head


head.print()
reverseKGroup(head, 4)
head.print()
