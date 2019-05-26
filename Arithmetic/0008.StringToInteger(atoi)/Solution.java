import java.util.regex.Pattern;
import java.util.regex.Matcher;

class Solution
{
    public static void main(String[] args) {
        String str = "-42";
        int res = myAtoi(str);
        System.out.println(res);
    }

    public static int myAtoi(String str)
    {
        str = str.trim();
        int symbol, res = 0, i = 0, length = str.length();

        if(length < 1)
        {
            return 0;
        }

        if(str.charAt(0) == '+')
        {
            symbol = 1;
            i = 1;
        }
        else if(str.charAt(0) == '-')
        {
            symbol = -1;
            i = 1;
        }
        else if(str.charAt(0) >= '0' && str.charAt(0) <= '9')
        {
            symbol = 1;
        }
        else{
            return 0;
        }
        for(;i < length; i ++)
        {
            if(str.charAt(i) < '0' || str.charAt(i) > '9')
            {
                return res;
            }

            int temp = Integer.parseInt(String.valueOf(str.charAt(i))) * symbol;

            if(res > Integer.MAX_VALUE / 10 || (res == Integer.MAX_VALUE / 10 && temp > 7))
            {
                return Integer.MAX_VALUE;
            }

            if(res < Integer.MIN_VALUE / 10 || (res == Integer.MIN_VALUE / 10 && temp < -8))
            {
                return Integer.MIN_VALUE;
            }

            res = 10 * res + temp;
        }

        return res;
    }

    public int myAtoi2(String str) {
        str = str.trim();

        String pattern = "^[\\+\\-]\\d+";

        Pattern p = Pattern.compile(pattern);

        Matcher m = p.matcher(str);

        String res;
        if(m.find()){
            res = str.substring(m.start(),m.end());
        }else{
            return 0;
        }
        
        if(res.length() == 1 && (res.charAt(0) == '+' || res.charAt(0) == '-' ))
        {
            return 0;
        }
        
        try{
            int r = Integer.parseInt(res);
            return r;
        }catch(Exception e){
            return res.charAt(0) == '-' ? Integer.MIN_VALUE : Integer.MAX_VALUE;
        }
    }
}