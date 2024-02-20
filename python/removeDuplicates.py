from typing import Dict, Optional


class ListNode:

    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

    def print(self):
        while self.next != None:
            print(self.val)
            self = self.next

        print(self.val)


def deleteDuplicates(head: Optional[ListNode]) -> Optional[ListNode]:
    seen: Dict[int, bool] = {}
    cur = head
    output: Optional[ListNode] = None
    curOutput = ListNode()

    while cur != None:
        if cur.val not in seen:
            seen[cur.val] = True
            if output == None:
                output = ListNode(cur.val)
                curOutput = output
            else:
                curOutput.next = ListNode(cur.val)
                curOutput = curOutput.next

        cur = cur.next

    return output


h = ListNode(1, ListNode(1, ListNode(2, ListNode(2, ListNode(4, None)))))
o: Optional[ListNode] = deleteDuplicates(h)

if o != None:
    o.print()
else:
    print("output is None, printing original")
    h.print()
