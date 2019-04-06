import java.util.HashMap;
import java.util.Map;

public class TwoSum
{
    public static void main(String[] args)
    {
        int[] nums = {1, 2, 3, 4, 5};
        int target = 3;
        int[] res = twoSum2(nums, target);
        System.out.print(res[0]);
        System.out.print(res[1]);
    }

    //9 ms	38.4 MB
    public static int[] twoSum(int[] nums, int target)
    {
        Map<Integer, Integer> map = new HashMap<>();
        for(int i = 0; i < nums.length; i ++){
            map.put(nums[i], i);
        }

        for(int i = 0; i < nums.length; i ++){
            int targetNum = target - nums[i];
            if(map.containsKey(targetNum) && map.get(targetNum) != i){
                return new int []{i, map.get(targetNum)};
            }
        }

        throw new IllegalArgumentException("wrong");
    }

    //6 ms	38.1 MB
    public static int[] twoSum2(int[] nums, int target) {
        Map<Integer, Integer> map = new HashMap<>();
        for (int i = 0; i < nums.length; i++) {
            int complement = target - nums[i];
            if (map.containsKey(complement)) {
                return new int[] { map.get(complement), i };
            }
            map.put(nums[i], i);
        }
        throw new IllegalArgumentException("No two sum solution");
    }

}
