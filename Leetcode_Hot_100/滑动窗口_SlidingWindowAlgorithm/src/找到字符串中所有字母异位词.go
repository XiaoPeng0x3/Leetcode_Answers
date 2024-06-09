func findAnagrams(s string, p string) []int {
    ans := []int{}
    // s是我们要寻找的串，p是给出的模式串
    
    //先获取一下长度
    slen := len(s)
    plen := len(p)
    if slen < plen{
        return nil
    }
    // 创建好我们的哈希表
    window := [27]int{}
    cnt := [27]int{}
    //把p加入
    for i, _ := range p{
        // cnt记录p
        cnt[p[i] - 'a']++
        // 窗口记录s
        window[s[i] - 'a']++
        
    }
    if cnt == window{
        ans = append(ans, 0)
    }
    //开始搜索p
    for i := 0; i < slen - plen; i++{
        //第一个元素马上划走
        window[s[i] - 'a']--
        // plen + i 位置进来
        window[s[i + plen] - 'a']++
        if window == cnt{
            ans = append(ans,i + 1)
        }
    }
    return ans
}