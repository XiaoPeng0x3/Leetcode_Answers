func minWindow(s string, t string) string {
    // 创建好哈希表
    // 对于这类题目来说，都有着固定的套路
    window := make(map[byte]int)
    need := make(map[byte]int)
    
    // 开始遍历字符串t
    tLen := len(t)
    sLen := len(s)
    minLen := sLen + 1
    start := 0
    for i := 0; i < tLen; i++ {
        need[t[i]]++ // 得到每个字符出现的个数
    }
    
    // 开始滑动
    l, r := 0, 0
    v := 0
    for r < sLen {
        sc := s[r]
        r++
        // 如果该字符出现在了need里面，那么我们就++
        if _, ok := need[sc]; ok {
            window[sc]++
            // 该字符出现的次数如果一样的话，那么这个字符就满足我们的要求
            if window[sc] == need[sc] {
                v++ // v的个数和len(need)一样的话，证明我们找到了一个合适的窗口
            }
        }
        // 出窗口
        // 要想一想出窗口要做什么事
        // 只要当前这个长的窗口一直包含我们需要的字符的话，那么就一直向左移动
        // 直到这个窗口不包含所有我们需要的字符了，那么我们就得重新向左去寻找
        for v == len(need) {
            // 进来这个循环，说明当前窗口内的字符已经满足我们的要求了
            if r - l < minLen {
                // 我们就得记录一下返回的过程
                start = l
                minLen = r - l // 最后返回的时候返回的字符串长度就是 s[start : start + minLen]
            }
            // 还是要向左移动窗口
            tc := s[l]
            l++
            if _, ok := need[tc]; ok {
                // 如果说左边的字符是我们需要的字符的话
                // 那么我们就得从滑动窗口里面删去这个值
                // 左侧字符串的个数减少
                window[tc]--
                // 小于说明我们减过头了，所以就又得继续扩大窗口
                if window[tc] < need[tc] {
                    v-- // 字符串得种类减少了一个
                }
            }
        }
    }
    if minLen == sLen + 1 {
        return ""
    }
    return s[start: start + minLen]
}
