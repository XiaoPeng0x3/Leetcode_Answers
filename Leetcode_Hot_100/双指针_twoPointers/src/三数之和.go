func threeSum(nums []int) [][]int {
    //先准备好答案
    ans := make([][]int, 0)
    n := len(nums)
    //对数组进行排序
    sort.Ints(nums)
    for i := 0; i < n - 2; i++ {
        //先去重nums[i],和上一个数作比较
        if i > 0 && nums[i] == nums[i - 1] {
            continue
        }
        // 否则则开始寻找
        x := nums[i]
        j := i + 1
        k := n - 1
        // 定i不动，移动j和k
        for j < k {
            if x + nums[j] + nums[k] < 0{
                j++
            } else if x + nums[j] + nums[k] > 0{
                k--
            }else{
                // 找到答案了
                ans = append(ans,[]int{x, nums[j], nums[k]})
                //然后开始对这两个nums[j]和nums[k]去重
                j++
                for j < k && nums[j] == nums[j - 1]{
                        j++
                }
                k--
                for j < k && nums[k] == nums[k + 1]{
                        k--
                }
            }

        }
        
    }
    return ans
}