# 题目

给定一个排序数组，你需要在 原地 删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。

不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。

示例 1:

给定数组 nums = [1,1,2], 

函数应该返回新的长度 2, 并且原数组 nums 的前两个元素被修改为 1, 2。 

你不需要考虑数组中超出新长度后面的元素。
示例 2:

给定 nums = [0,0,1,1,1,2,2,3,3,4],

函数应该返回新的长度 5, 并且原数组 nums 的前五个元素被修改为 0, 1, 2, 3, 4。

你不需要考虑数组中超出新长度后面的元素。
 
说明:

为什么返回数值是整数，但输出的答案是数组呢?

请注意，输入数组是以「引用」方式传递的，这意味着在函数里修改输入数组对于调用者是可见的。

你可以想象内部操作如下:
// 根据你的函数返回的长度, 它会打印出数组中该长度范围内的所有元素。
for (int i = 0; i < len; i++) {
    print(nums[i]);
}

# 分析

只要遍历一遍数组，并交换不相等元素即可，和排序算法相比算比较简单的，写起来也很简单。

# 解答

solution.go的时间复杂度为O(n)，空间复杂度为O(1)。我比较奇怪的是，我只用了1个int变量，竟然内存消耗只打败了9.09%的用户，估计是leetcode的内存消耗计算是有误差的。

执行用时 :8 ms, 在所有 Go 提交中击败了93.29%的用户
内存消耗 :4.6 MB, 在所有 Go 提交中击败了9.09%的用户