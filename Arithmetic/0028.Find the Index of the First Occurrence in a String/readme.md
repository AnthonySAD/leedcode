# 题目

给你两个字符串 haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串的第一个匹配项的下标（下标从 0 开始）。如果 needle 不是 haystack 的一部分，则返回  -1 。



# 分析

字符串算法不熟，直接暴力解法。时间复杂度是m*n，的确拉了点。看了下评论，写了个kmp算法。

kmp太难了，看懂了思路，但自己写就老是写错。