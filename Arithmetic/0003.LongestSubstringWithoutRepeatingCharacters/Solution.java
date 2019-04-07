import java.util.HashMap;
import java.util.HashSet;
import java.util.Set;

public class Solution
{
    public static void main(String[] args) {
        String s = "abcabcdddddahstgasewr";
        int num1 = solution1(s);
        System.out.print(num1);
        System.out.print("-");
        int num2 = solution2(s);
        System.out.print(num2);
        System.out.print("-");
        int num3 = solution3(s);
        System.out.print(num3);
    }

    public static int solution1(String s){
        int len = s.length();
        HashSet<Character> set = new HashSet();
        int max = 0, i = 0, j = 0;

        while(i < len && j < len){
            if(!set.contains(s.charAt(j))){
                set.add(s.charAt(j));
                max = Math.max(max, j - i + 1);
                j ++;
            }else{
                set.remove(s.charAt(i));
                i ++;
            }
        }
        return max;
    }

    public static int solution2(String s){
        int len = s.length(), max = 0;
        HashMap<Character, Integer> map = new HashMap();
        for(int start = 0, end = 0; end < len; end ++){
            if(map.containsKey(s.charAt(end)) && map.get(s.charAt(end)) >= start){
                start = map.get(s.charAt(end)) + 1;
            }else{
                max = Math.max(max, end - start + 1);
            }
            map.put(s.charAt(end), end);
        }

        return max;

    }

    public static int solution3(String s){
        int len = s.length(), max = 0;
        int[] arr = new int[128];
        for(int start = 0, end = 0; end < len; end ++){
            if(arr[s.charAt(end)] != 0 && arr[s.charAt(end)] >= start){
                start = arr[s.charAt(end)] + 1;
            }else{
                max = Math.max(max, end - start + 1);
            }
            arr[s.charAt(end)] = end;
        }

        return max;
    }
}