func trap(height []int) int {
    n := len(height)
    ans := 0
    //双指针
    l := 0
    r := n - 1
    // 记录一下左右两侧的最大值
    lmax := 0
    rmax := 0
    // 不越界
    for l <= r{
        lmax = max(height[l], lmax)
        rmax = max(height[r], rmax)
        if lmax < rmax{
            // 左边是短板，那么取的水就主要取决于左边的木板
            // 柱子高度也会占用水的地方
            ans += lmax - height[l]
            l++
        }else{
            // 右边是短板
            ans += rmax- height[r]
            r--
        }
    }
    return ans
}