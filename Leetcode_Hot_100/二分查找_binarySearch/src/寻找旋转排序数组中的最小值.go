func findMin(nums []int) int {
    n := len(nums)
    l, r := 0, n - 2
    for l <= r {
        mid := l + (r - l) / 2
        if nums[mid] < nums[n -1] {
            r = mid - 1
        } else {
            l = mid + 1
        }
    }
    return nums[l]
}