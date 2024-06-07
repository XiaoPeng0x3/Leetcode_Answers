// 暴力搜索
func twoSum(nums []int, target int) []int {
    for i := range nums{
        for j := range nums{
            if i != j && nums[i] + nums[j] == target{
                return []int{i,j} // 满足情况
            }
        }
     }
    return nil
}
//哈希表
func twoSum(nums []int, target int) []int {
    hash := make(map[int]int)
    for i, j := range nums{
        if p,ok := hash[target - j];ok{
            return []int{i, p}
        }
        hash[j] = i
    }
    return nil
}