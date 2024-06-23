func searchInsert(nums []int, target int) int {
    // 假设target在我们的列表里面，那么就是最普通的二分法
    l, r := 0, len(nums) - 1
    x, mid := 0, 0
    for l <= r{
        mid = l + (r - l) / 2
        x = nums[mid]
        if x == target {
            return mid
        } else if x < target {
            l = mid + 1
        } else {
            r = mid - 1
        }
    }
    // 如果找不到，那么我们应该返回哪个下标索引呢
    // 可以想一下，退出循环的条件就是 l > r,更确切来说，是 l = r + 1
    // 现在想一下，当 l == r 时，如果x 比 target要小，那么执行的就是 l = mid + 1
    // 反之，则会执行 r = mid - 1
    // 可以发现，最后执行的那次循环的x就是要排在我们target的前面的那个数
    // 而最后 l执行了++操作，所以最后返回l即可
    // 也可以返回 r + 1
    return l
}