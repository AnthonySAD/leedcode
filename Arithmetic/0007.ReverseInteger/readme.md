# 题目

给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。

假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−2<sup>31</sup>,  2<sup>31</sup> − 1]。请根据这个假设，如果反转后整数溢出那么就返回 0。

# 分析

第一感觉就是先把整数转换成字符串,然后生成新的字符串,然后转出整数就好了。看了下标答，原来可以通过数学计算的方式得出，不用额外的储存空间，太巧妙了。