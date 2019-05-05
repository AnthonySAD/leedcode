public class Solution
{
    public static void main(String[] args) {
        int[] a = {1, 2};
        int[] b = {1, 2};

        double num = answer(a, b);

        System.out.print(num);
    }
    
    public static double answer(int[] a, int[] b){
        //确保数组a的长度小于等于b,否则交换数组a,b
        int aLen = a.length;
        int bLen = b.length;
        if(aLen > bLen){
            int[] tempArr = a;
            a = b;
            b = tempArr;
            int tempLen = aLen;
            aLen = bLen;
            bLen = tempLen;
        }

        //使用二分法查找时先初始化边界
        int aPMin = 0;
        int aPMax = aLen;
        int halfLen = (aLen + bLen + 1) / 2;

        //因为只要集合不都为空,则必然有解,所以不需要判断
        while (true) {
            //使用二分法求值
            int aP = (aPMin + aPMax) / 2;
            //通过公式计算出bP的值
            int bP = halfLen - aP;

            //先判断非边界情况
            //当b的左集合最大值大于a的右集合最小值时,需要增大a的右集合的最小值,既aP右移,所以把aP的左边界右移到aP+1
            //a,b的左集合长度为aP,bP,所以最大下标为aP-1,bP-1
            if (aP < aPMax && b[bP-1] > a[aP]){
                aPMin = aP + 1;
            }
            //当a的左集合最大值大于b的右集合最小值时,需要减小a的右集合的最大值,既aP左移,所以把aP的右边界右移到aP-1
            else if (aP > aPMin && a[aP-1] > b[bP]) {
                aPMax = aP - 1;
            }
            //当使用二分法求出的值为边界时，则此时二分法的左右边界已重合，已无法进行下去了，同时本题必然有解，所以此值就为正确答案
            else {
                int maxLeft;
                int minRight;
                //求出左集合的最大值
                if (aP == 0) {
                    maxLeft = b[bP - 1];
                }else if (bP == 0) {
                    maxLeft = a[aP - 1];
                }else{
                    maxLeft = Math.max(b[bP - 1], a[aP - 1]);
                }

                //当集合总数为奇数时，则中位数则为maxLeft。
                //因为之前计算半长的时候，实际数值增加了0.5，这将导致左集合比右集合多一个值，及中位数为maxLeft
                if ((aLen + bLen) % 2 == 1) {
                    return maxLeft;
                }

                //求右集合的最小值
                if (aP == aLen) {
                    minRight = b[bP];
                }else if(bP == bLen){
                    minRight = a[aP];
                }else{
                    minRight = Math.min(b[bP], a[aP]);
                }
                
                //当集合总数为偶数时,中位数为左集合的最大值与右集合的平均值
                //此处用到了自动类型转换,java会自动把基本数据类型向大数值转换
                return (maxLeft + minRight) / 2.0;
            }
        }
    }
}