# 题目
你这个学期必须选修 numCourses 门课程，记为 0 到 numCourses - 1 。

在选修某些课程之前需要一些先修课程。 先修课程按数组 prerequisites 给出，其中 prerequisites[i] = [ai, bi] ，表示如果要学习课程 ai 则 必须 先学习课程  bi 。

例如，先修课程对 [0, 1] 表示：想要学习课程 0 ，你需要先完成课程 1 。
请你判断是否可能完成所有课程的学习？如果可以，返回 true ；否则，返回 false 。

 

示例 1：

输入：numCourses = 2, prerequisites = [[1,0]]
输出：true
解释：总共有 2 门课程。学习课程 1 之前，你需要完成课程 0 。这是可能的。
示例 2：

输入：numCourses = 2, prerequisites = [[1,0],[0,1]]
输出：false
解释：总共有 2 门课程。学习课程 1 之前，你需要先完成​课程 0 ；并且学习课程 0 之前，你还应先完成课程 1 。这是不可能的。

# 分析
记得有个啥算法是适合求依赖关系的，但忘了。。。
想了下，应该先找出那些不需要依赖的课程。如果所有课程都需要依赖，则变成了循环依赖，必定无解。想了下暴力循环时间复杂度太高了。可以先记录每个课程需要几个依赖，然后遍历依赖数组，如果复合的减去依赖数，如果依赖数归0，则说明可学。
做完后发现还是太慢了，只击败了20%的人。
看了下答案，思路和我差不多，但看的是头大，原来这个叫拓扑排序。