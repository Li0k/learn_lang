import java.lang.reflect.Array;
import java.util.*;

public class Solution03 {
    public int lengthOfLongestSubstring(String s) {
        //1.利用set的元素唯一性来排除重复
//        Set<Character> set = new HashSet<>();
//        int res = 0;
//        int left = 0,right = 0;
//
//        while (right < s.length()){
//            char c = s.charAt(right);
//            if (!set.contains(c)){
//                set.add(c);
//                res = Math.max(res,set.size());
//                right++;
//            } else {
//                //遇到重复时，从左端开始删除元素，直到无重复
//                set.remove(s.charAt(left++));
//            }
//        }
//
//        return res;

        //2.利用数组

        int res = 0;
        int[] map = new int[256];
        Arrays.fill(map, -1);

        for (int i = 0, j = 0; j < s.length(); j++) {
            i = Math.max(map[s.charAt(j)] + 1, i);
            res = Math.max(res, j - i + 1);
            map[s.charAt(j)] = j;
        }

        //3.利用map,记录元素与index的映射
//        int res = 0;
//        Map<Character, Integer> map = new HashMap<>();
//        for (int i = 0, j = 0; j < s.length(); j++) {
//            char c = s.charAt(j);
//            if (map.containsKey(c)) {
//                i = Math.max(map.get(c) + 1, i);
//            }
//            res = Math.max(res, j - i + 1);
//            map.put(c,j);
//        }


        return res;
    }

    public static void main(String... args) {
        String s = "pwwkew";
        System.out.println(new Solution03().lengthOfLongestSubstring(s));
        String s1 = "aab";
        System.out.println(new Solution03().lengthOfLongestSubstring(s1));
        String s2 = "dvdf";
        System.out.print(new Solution03().lengthOfLongestSubstring(s2));
    }
}
