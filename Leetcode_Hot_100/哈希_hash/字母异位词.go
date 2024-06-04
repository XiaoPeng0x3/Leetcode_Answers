func groupAnagrams(strs []string) [][]string {
    // 先观察一下参数，给定的是字符串类型的数组，返回一个二维数组
    // 按照我们的思路，首先对这些数组进行排序
    mp := make(map[string][]string) // 创建好hash表，键是string类型的，值是[]string数组
    for _, str := range strs{
        //首先一点是进行排序啊
        //这里用到了切片的知识，我们可以把字符串转换为字符串切片
        s := []byte(str)
        // 然后每次都对切片进行排序
        sort.Slice(s, func (i, j int) bool { return s[i] < s[j]} )
        // 拍好序后，再转换为字符串的形式
        s1 := string(s)
        //转换为后插入到键里面,把对应的值插到键里面
        mp[s1] = append(mp[s1], str)
    }
    // 然后，经过上面那个循环，我们就已经准备好了，所以下一步直接创建好返回值
    ans := make([][]string, 0, len(mp)) //这是必须的，就像是vector里面的初始化一样
    for _, v := range mp{
        ans = append(ans, v)
    }
    return ans
}