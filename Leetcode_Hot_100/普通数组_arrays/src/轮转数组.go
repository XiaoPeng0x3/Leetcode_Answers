//直接原地交换
func rotate(nums []int, k int)  {
    n := len(nums)
    k = k % n
    r(nums, 0, n - 1)
    r(nums, 0, k - 1)
    r(nums, k, n - 1) 
}

func r(nums []int, i int, j int){
    for i <= j{
        nums[i], nums[j] = nums[j], nums[i]
        i++
        j--
    }
}

//使用额外数组
func rotate(nums []int, k int)  {
    n := len(nums)
    ans := make([]int, n)
    for i := 0; i < n; i++{
        ans[(i + k) % n] = nums[i]
    }
    copy(nums, ans)
}