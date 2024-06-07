func maxArea(height []int) int {
    ans := 0
    l := 0
    lens := len(height)
    r := lens - 1
    for l <= r {
        area := min(height[l], height[r]) * (r - l)
        ans = max(area, ans)
        if height[l] < height[r]{
            l++
        }else{
            r--
        }
    }
    return ans
}