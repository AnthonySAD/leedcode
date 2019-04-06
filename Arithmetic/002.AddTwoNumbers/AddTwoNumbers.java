import ListNode;

public class AddTwoNumbers
{

    //10 ms	45 MB
    public ListNode AddTwoNumbers(ListNode l1, ListNode l2)
    {
        ListNode dummyHead = new ListNode(0);
        currentNode = dummyHead;
        int carry = 0;

        while(l1 != null || l2 != null){
            int num1 = l1 != null ? l1.val : 0;
            int num2 = l2 != null ? l2.val : 0;
            int sum = num1 + num2 + carry;
            ListNode carry = sum/10;
            currentNode.next = new ListNode(sum - 10 * carry);
            currentNode = currentNode.next;

            l1 = l1 == null ? l1 : l1.next;
            l2 = l2 == null ? l2 : l2.next;
        }

        if(carry > 0){
            currentNode.next = new ListNode(carry);
        }

        return dummyHead.next;

}