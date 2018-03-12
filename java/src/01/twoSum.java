import sun.applet.Main;

import java.util.HashMap;

public class twoSum {
    public int[] TwoSum(int[] nums, int target) {
        HashMap<Integer,Integer> m = new HashMap<>();
        int[] res = new int[2];

        for (int i = 0; i < nums.length; i++){
            if (m.containsKey(target - nums[i])){
                res[0] = i;
                res[1] = m.get(target-nums[i]);
                break;
            }

            m.put(nums[i],i);
        }
        return res;
    }

    public static void main(String...args){
        int[] nums = {2, 7, 11, 15};
        int target = 9;
//        new twoSum().TwoSum(nums,target);
        System.out.println(new twoSum().TwoSum(nums,target));
    }
}
