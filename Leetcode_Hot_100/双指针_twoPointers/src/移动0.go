func moveZeroes(nums []int)  {
    l := 0
    r := 0
    lens := len(nums)
    for r < lens{
        // 找到第一个不为0的数
        if nums[r] != 0{
            //交换(和python一样啊)
            nums[l], nums[r] = nums[r], nums[l]
            //移动慢指针
            l++
        }
        r++
    }
}