func searchRange(nums []int, target int) []int {
    // 这道题只要我们分开寻找即可
    ans := make([]int, 0, 2)
    // 先开始寻找左边
    n := len(nums)
    l1, r1, l2, r2 := 0, n - 1, 0, n - 1
    // 先找左边第一次出现的那个数
    m1, m2 := 0, 0
    x1, x2 := 0, 0
    index1, index2 := -1, -1
    
    for l1 <= r1 {
        m1 = l1 + (r1 - l1) / 2
        x1 = nums[m1]
        if x1 == target {
            r1 = m1 - 1
            index1 = m1
        } else if x1 < target {
            l1 = m1 + 1
        } else {
            r1 = m1 - 1
        }
    }
    ans = append(ans, index1)

    for l2 <= r2 {
        m2 = l2 + (r2 - l2) / 2
        x2 = nums[m2]
        if x2 == target {
            l2 = m2 + 1
            index2 = m2
        } else if x2 < target {
            l2 = m2 + 1
        } else {
            r2 = m2 - 1
        }
    }
    ans = append(ans, index2)

    return ans
}