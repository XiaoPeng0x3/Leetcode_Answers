func firstMissingPositive(nums []int) int {
    // 好像是有一些思路了
    // 我们要做的就是原地进行哈希表的交换，把nums[i]交换到nums[nums[i] - 1]的地方去
    // 也就是说，在一个数组里面，把这个数组的所有元素全部安排好，然后再看下标哪里不对应

    n := len(nums)
    for i := 0; i < n; i++{
        // 我们所交换的数必须都是1~n+1范围内的
        for nums[i] <= n && 1 <= nums[i] && nums[nums[i] - 1] != nums[i]{
            // 已经在正确位置的数不必交换
            nums[i], nums[nums[i] - 1] = nums[nums[i] - 1], nums[i]
        }
    }

    for i := 0; i < n; i++{
        if i + 1 != nums[i]{
            return i + 1
        }
    }
    return n + 1
}