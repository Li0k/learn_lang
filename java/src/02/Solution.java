public class Solution {

    public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        ListNode newList = new ListNode(-1);//头节点方便操作
        ListNode cur = newList;//标定当前节点
        int flg = 0;//判断进位

        while (l1 != null || l2 != null) {
            int v1 = l1 == null ? 0 : l1.val;
            int v2 = l2 == null ? 0 : l2.val;
            int sum = v1 + v2 + flg;

//            if (sum >= 10) flg = 1;
            flg = (sum >= 10) ? 1 : 0;
            newList.val = (sum % 10);
//            curr = newList.next;
            cur.next = new ListNode(sum % 10);

            if (l1 != null) l1 = l1.next;
            if (l2 != null) l2 = l2.next;

            cur = cur.next;
        }

        //处理最后的进位
        if (flg == 1) cur.next = new ListNode(1);
        return newList.next;
    }

    public  static void main(String...args){
        ListNode l1 = new ListNode(2);
        l1.next = new ListNode(4);
        l1.next.next = new ListNode(3);

        ListNode l2 = new ListNode(5);
        l2.next = new ListNode(6);
        l2.next.next = new ListNode(4);

        ListNode l3 = new Solution().addTwoNumbers(l1,l2);
        System.out.println(l3);

    }

}

class ListNode {
    int val;
    ListNode next;

    ListNode(int x) {
        val = x;
    }

    @Override
    public String toString() {
        return "ListNode{" +
                "val=" + val +
                ", next=" + next +
                '}';
    }

}
