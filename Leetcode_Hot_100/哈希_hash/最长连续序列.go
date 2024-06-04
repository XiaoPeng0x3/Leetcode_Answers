func longestConsecutive(nums []int) int {
    // 先把元素插到哈希表里面
    mp := map[int]bool{}
    for _, val := range nums {
        mp[val] = true
    }
    ans := 0
    // 插好了，现在就可以寻找了
    for k, _ := range mp {
        // 为false说明k是端点
        if !mp[k - 1] {
            cnt := k
            for mp[cnt] {
                cnt++
            }
            // 下面我们可以说一下， cnt返回的是最后端点+1的值，k是起点，正好相减即可
            ans = max(ans, cnt - k)
        }
    } 
    return ans
}