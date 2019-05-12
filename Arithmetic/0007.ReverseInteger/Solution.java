class Solution
{
    public static void main(String[] args) {
        
    }

    public reverse(int x){
        int ans = 0;
        while(x != 0)
        {
            int num = x % 10;
            x /= 10;
            //正数情况
            if(ans > Integer.MAX_VALUE / 10 ){
                return 0;
            }
            //max为9223372036854775807,所以>7就超了
            if(ans == Integer.MAX_VALUE / 10 && num > 7){
                return 0;
            }
            
            //负数情况
            if(ans < Integer.MIN_VALUE / 10 ){
                return 0;
            }
            //min为-9223372036854775808,所以<-8就超了
            if(ans == Integer.MIN_VALUE / 10 && num < -8){
                return 0;
            }

            ans = ans * 10 + num;
        }

        return ans;
    }
}