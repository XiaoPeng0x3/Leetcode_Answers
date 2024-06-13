 func productExceptSelf(nums []int) []int {
     n := len(nums)
     // 先准备好数组
     L := make([]int, n)
     R := make([]int, n)
     ans := make([]int, n)
     // 最左边那个数是没有乘积的，所以我们在收集的时候，最左边那个就是1
     L[0] = 1
     //同理，左边那个数也是
     R[n - 1] = 1
     //此时就可以开始计算了
     for i := 1; i < n; i++{
         L[i] = L[i - 1] * nums[i - 1]
     }
     //此时也可以计算出来左边的数
     //最后那个数的右区间也是1，所以不用从最后一个数开始，直接从倒数第二个数开始
     for j := n - 2; j >= 0; j--{
         R[j] = R[j + 1] * nums[j + 1]
     }
     //此时就把所有的收集好了，然后我们就直接乘积即可
     
     for i := 0; i < n; i++{
         ans[i] = L[i] * R[i]
     }
     return ans
 }

 func productExceptSelf(nums []int) []int {
     //左边还是一样的
     n := len(nums)
     // 先准备好数组
     ans := make([]int, n)
     // 最左边那个数是没有乘积的，所以我们在收集的时候，最左边那个就是1
     ans[0] = 1
     //此时就可以开始计算了
     for i := 1; i < n; i++{
         ans[i] = ans[i - 1] * nums[i - 1]
     }
     // 计算好左边之后，右边的乘积直接在答案上面进行构造
     rnum := 1 // 辅助变量
     for j := n - 1; j >= 0; j--{
         ans[j] = ans[j] * rnum //最右边的数是左边的乘积乘以rnum
         rnum *= nums[j]
     }
     return ans
 }