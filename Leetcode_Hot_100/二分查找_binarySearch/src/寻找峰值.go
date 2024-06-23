// 直接遍历，但是不符合时间复杂度的要求
func findPeakElement(nums []int) int {
    max := math.MinInt
    index := 0
    for i, val := range nums {
        if val > max {
            max = val
            index = i
        }
    }
    return index
}

// 二分法，实现O(logn)
func findPeakElement(nums []int) int {
    n := len(nums)
    l, r, mid := 0, n - 1, 0
    
    get := func(i int) {
        if i == -1 || i == n{
            return math.MinInt
        }
        return nums[i]
    }
    
    for l <= r {
        mid = l + (r - l) / 2
        // 比两边的元素都要大
        if get(mid) > get(mid - 1) && get(mid) > get(mid + 1) {
            return mid
        } else if get(mid) > get(mid - 1) {
            l = mid + 1
        } else {
            r = mid - 1
        }
    }
    return -1
}