// 贪心
func maxSubArray(nums []int) int {
    // 这道题目还是有印象的，一种做法是贪心做法，就是去维护一个sum,如果sum < 0 的话，肯定不是我们的最大值
    sum := 0
    ans := -100010
    for _, val := range nums{
        sum += val
        ans = max(ans, sum)
        //小于0不去计入数组和
        if sum < 0{
            sum = 0
        }
    }
    return ans
}

//分支递归
func maxSubArray(nums []int) int {
    n := len(nums)
    return maxSum(nums, 0, n - 1)
}

func maxSum(nums []int, s int, e int) int{
    if s >= e{
        return nums[s]
    }
    m := (s + e) / 2
    leftMax := maxSum(nums,s, m)
    rightMax := maxSum(nums, m + 1, e)
    // 跨越区间的最大值
    leftSum, rightSum := 0, 0
    lSum, rSum := -100086, -100086 //跨越区间
    for i := m; i >= s; i--{
        leftSum += nums[i]
        lSum = max(leftSum, lSum) //记录连续左区间最大值
    }
    // 同理计算rightSUm
    for j := m + 1; j <= e; j++{
        rightSum += nums[j]
        rSum = max(rSum, rightSum) //记录连续右区间最大值
    }
    // 再计算一下跨越区间的最大值
    allSum := lSum + rSum
    return max(max(leftMax, rightMax), allSum)
    
}