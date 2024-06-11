// 暴力解法
func subarraySum(nums []int, k int) int {
    n := len(nums)
    ans := 0
    for s := 0; s < n; s++{
        sum := 0
        for end := s; end < n; end++{
            sum += nums[end]
            if sum == k{
                ans++
            }
        }
    }
    return ans
}
// 哈希表+前缀和优化
func subarraySum(nums []int, k int) int {
    ans := 0
    // 创建好哈希表
    cnt := map[int]int{0:1} //前缀和为0的出现次数是1次
    //计算好一个前缀和数组
    perSum := 0
    for _, val := range nums{
        perSum += val
        // 从哈希表里面查询perSum[i]-k是否存在
        if a, ok := cnt[perSum - k];ok{
            //存在，那么就加上这个前缀出现的次数即可
            ans += a
        }
        //在这里添加前缀的出现次数
        cnt[perSum]++
    }
    return ans
}