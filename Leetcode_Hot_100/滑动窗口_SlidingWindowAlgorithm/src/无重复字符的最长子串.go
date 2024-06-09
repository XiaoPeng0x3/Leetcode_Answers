func lengthOfLongestSubstring(s string) int {
    // 先声明一下答案和哈希表
    ans := 0
    cnt := [128]bool{}
    // 声明左指针
    l := 0
    for r, c := range s{
        // 如果当前字符存在的话，那么我们就删除
        // 删除的时候，一定是删除的左指针
        for cnt[c]{
            cnt[s[l]] = false
            l++
        }
        // 不存在的话，把这个字符加进去
        cnt[c] = true
        // 更新答案
        ans = max(ans, r - l + 1)
    }
}