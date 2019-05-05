public class Solution
{
    public static void main(String [] args)
    {
        String s = "abacab";
        String ans = answer(s);
        System.out.print(ans);
    }
    //103 ms,   65.7 MB
    public static String answer(String s)
    {
        int n = s.length();
        boolean[][] m = new boolean[n][n];

        int start = 0;
        int end = 0;
        int diff = 0;

        if(n < 2){
            return s;
        }

        if(n == 2){
            if(s.charAt(0) == s.charAt(1)){
                return s;
            }else{
                return s.substring(0, 1);
            }
        }

        for(int i = n - 2; i >= 0; i --){
            m[i][i] = true;

            for(int j = i + 1; j < n; j ++){

                if(s.charAt(i) == s.charAt(j) && (j - i < 3 || m[i+1][j-1])){
                    m[i][j] = true;
                    if(j - i > diff){
                        start = i;
                        end = j;
                        diff = j - i;
                    }
                }
            }
        }

        return s.substring(start, end + 1);
    }

}
