class Solution
{
    public static void main(String[] args) {
        int number = 12321;
        boolean res = solution(number);
        System.out.println(res);
    }

    public static boolean solution(int x){
        String str = String.valueOf(x);
        int len = str.length();
        int halfLen = len / 2;
        for(int i = 0; i < halfLen; i ++)
        {
            if(str.charAt(i) != str.charAt(len - i - 1)){
                return false;
            }
        }
        return true;
    }

    public static boolean solution2(int x){
        if(x < 0){
            return false;
        }
        if(x == 0){
            return true;
        }
        if(x % 10 == 0){
            return false;
        }
        int right = 0;
        int left = x;
        while(left >= right){
            if(left / 10 == right){
                return true;
            }
            right = 10 * right + left % 10;
            left /= 10;
            if(left == right){
                return true;
            }
        }

        return false;
    }

}