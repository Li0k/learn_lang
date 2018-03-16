import com.sun.xml.internal.xsom.impl.Const;

import java.lang.reflect.Array;
import java.util.Arrays;

public class Solution5 {
//    String res;
//    int maxNum = 0;
//
//    public String longestPalindrome(String s) {
//        if (s.length() == 1) return s;
//
//        for (int i = 0; i < s.length(); i++) {
//            checkPalindrome(s, i, i); //奇数
//            checkPalindrome(s, i, i + 1);//偶数
//        }
//
//        return res;
//
//    }
//
//    public void checkPalindrome(String s, int low, int high) {
////        int low = i, high = i;
//
////        while (low >= 0 && high < s.length()) {
////            if (s.charAt(low) == s.charAt(high)) {
////                if (high - low + 1 > maxNum){
////                    maxNum = high-low+1;
////                    res = s.substring(low,high + 1);
////                }
////                    low--;
////                high++;
////            } else {
////                return;
////            }
////
////        }
//
//        while (low >= 0 && high < s.length() && s.charAt(low) == s.charAt(high)) {
//            low--;
//            high++;
//        }
//
//        if (high - low + 1 > maxNum) {
//            maxNum = high - low + 1;
//            res = s.substring(low, high + 1);
//        }
//    }

    //Manacher's Algorithm 马拉车算法 不成功......
//    int maxLen = 0, center = 0;
//
//    public String longestPalindrome(String s) {
////
//        StringBuilder ts = new StringBuilder("$#");
//        for (int i = 0;i < s.length();i++){
//            ts.append(s.charAt(i));
//            ts.append("#");
//        }
//
//        int[] p = new int[ts.length()];
//        Arrays.fill(p,0);
//        int mx = 0;
//        int id = 0; //mx为最右端，id为对称点
//
//        char[] t = ts.toString().toCharArray();
//        for (int i = 1; i < t.length; i++) {
//            //分为i在mx左右进行讨论
//            p[i] = mx > i ? Math.min(p[2 * id - i], mx - i) : 1;//p[2 * id - i]为对称点，根据对称性分析p[i]，分析点p[i] < mx - i, >= mx - i;
//            //若i在mx右边与扩展p[i]一起操作
//
//            while (t[i + p[i]] == t[i - p[i]]) ++p[i];
//
//            if (mx < i + p[i]) {
//                mx = i + p[i];
//                id = i;
//            }
//
//            if (maxLen < p[i]) {
//                maxLen = p[i];
//                center = i;
//            }
//
//        }
//
//        return s.substring(center- maxLen / 2, maxLen - 1);
//    }


    //动态规划
    /*dp[j][i]  =  true                j = i
                =  str[i] == str[j]    i - j = 1
                =  str[i] == str[j] && dp[j + 1][i - 1]   i - j > 1  回文串的子串也是回文串
    */
    public String longestPalindrome(String s) {
        if (s == null || s.length() == 1) return s;

        int size = s.length();
        boolean[][] dp = new boolean[size][size];

        int maxLen = 1;
        int left = 0, right = 0;

        for (int i = 0; i < s.length(); ++i) {
            for (int j = 0; j <= i; ++j) {
                if (i - j < 2) {
                    dp[j][i] = s.charAt(i) == s.charAt(j);
                } else {
                    dp[j][i] = (s.charAt(i) == s.charAt(j)) && dp[j + 1][i - 1];
                }

                if (dp[j][i] && maxLen < (i - j + 1)) {
                    maxLen = i - j + 1;
                    left = j;
                    right = i;
                }
            }
        }
//        System.out.println(left);
//        System.out.println(right);
//        System.out.println(maxLen);
        return s.substring(left, right + 1);
    }


    public static void main(String... args) {
        String s1 = "12212";
        String s2 = "waabwswfd";
        System.out.println(new Solution5().longestPalindrome(s1));
        System.out.println(new Solution5().longestPalindrome(s2));
    }
}
