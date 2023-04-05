
// Definition for singly-linked list.
class ListNode {
    val: number
    next: ListNode | null
    constructor(val?: number, next?: ListNode | null) {
        this.val = (val===undefined ? 0 : val)
        this.next = (next===undefined ? null : next)
    }
}

function addTwoNumbers(l1: ListNode | null, l2: ListNode | null): ListNode | null { 

    let l1Current = l1;
    let l2Current = l2;

    let resultNode: ListNode = new ListNode();
    let currentResultNode: ListNode | null = resultNode;
    let previousResultNode = currentResultNode;

    while (l1Current !== null || l2Current !== null) {
        let sum = (l1Current?.val ?? 0) + (l2Current?.val ?? 0) + (currentResultNode?.val ?? 0);
        if (sum === 0 && (!l1Current?.next && !l2Current?.next)) {
            break;
        }

        if (sum > 9) {
            let sumString = sum + '';
            let firstNum = new Number(sumString.substring(0, 1));
            let secondNum = new Number(sumString.substring(1));;
            let nextNode = new ListNode()
            nextNode.val = firstNum.valueOf();
            currentResultNode.val = secondNum.valueOf();
            currentResultNode.next = nextNode;
            previousResultNode = currentResultNode;
            currentResultNode = currentResultNode.next;
        } else {
            currentResultNode.val = sum;
            currentResultNode.next = new ListNode();
            previousResultNode = currentResultNode;
            currentResultNode = currentResultNode.next;
        }

        l1Current = l1Current?.next ?? null;
        l2Current = l2Current?.next ?? null;
    }
    
    if (currentResultNode.val === 0) {
        previousResultNode.next = null;
    }

    return resultNode;

};