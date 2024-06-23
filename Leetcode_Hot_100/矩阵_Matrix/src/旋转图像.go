func rotate(matrix [][]int)  {
    // 先按照对角线旋转
    m := len(matrix)
    n := len(matrix[0])
    for i :=0; i < m; i++ {
        for j := i + 1; j < n; j++ { // 要注意交换的循序，每个元素只做一次交换
            matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
        }
    }
    // 再把每一行旋转一下
    for i := 0; i < m; i++ {
        swapRow(matrix[i])
    }
}

func swapRow(nums []int){
    l, r := 0, len(nums) - 1
    for l <= r{
        nums[l], nums[r] = nums[r], nums[l]
        l++
        r--
    }
}