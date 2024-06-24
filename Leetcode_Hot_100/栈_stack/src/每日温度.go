// 从右向左遍历

func dailyTemperatures(temperatures []int) []int {
    n := len(temperatures) //获得数组长度
    ans := make([]int, n)
    
    st := []int{}
    
    // 从右向左遍历
    for i := n - 1; i >= 0; i-- {
        // 栈不为空，并且栈顶元素小于遇到的元素是，要弹出，直到栈顶位置比遇到的那个元素大为止
        for len(st) > 0 && temperatures[i] >= temperatures[st[len(st) - 1]] {
            // 开始弹出栈元素
            st = st[ :len(st) - 1]
        }
        // 弹出后就是遇到了比自己大的那个数
        // 而且比自己大的那个数还是栈顶
        if len(st) > 0 {
            ans[i] = st[len(st) - 1] - i
        }
        // 添加新的元素入栈
        st = append(st, i)
    }
    return ans
    
}