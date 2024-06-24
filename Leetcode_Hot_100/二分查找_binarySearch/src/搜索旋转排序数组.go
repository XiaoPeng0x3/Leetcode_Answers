func search(nums []int, target int) int {
    l, r := 0, len(nums) - 1
    x, mid := 0, 0
    for l <= r {
        mid = l + (r - l) / 2
        x = nums[mid]
        if x == target {
            return mid
        }else if x < nums[r] {
            // 右边有序

            // 看一看target在哪个区间里面
            // 确实是在右区间
            if x < target && target <= nums[r]{
                l = mid + 1
            } else {
                r = mid - 1
            }
        } else {
            // 确实是在左区间
            if x > target && nums[l] <= target {
                r = mid - 1 //去左边寻找
            } else {
                l = mid + 1
            }
        }
    }
    return -1
}