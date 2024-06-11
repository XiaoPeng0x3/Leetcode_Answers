func merge(intervals [][]int) [][]int {
    // 先排序
    sort.Slice(intervals, func(i, j int)bool{
        return intervals[i][0] < intervals[j][0]
    })
    ans := [][]int{}
    n := len(intervals)
    //先把区间的端点值记录下来
    start, end := intervals[0][0], intervals[0][1]
    for i := 1; i < n; i++{
        //如果intervals[1][0]（x2 < y1）的情况，不需要合并，直接添加结果即可
        if intervals[i][0] > end{
            ans = append(ans,[]int{start, end})
            //直接向后迭代，看一看后面的区间的情况
            start = intervals[i][0]
            end = intervals[i][1]
        } else {
            // 需要合并的情况，即 x2 <= y1,那么，左端点一定是x1,而右端点就不知道是谁了。第一个左端点一定是最小的，如果说有n个区间可以合并([x1,y1]、[x2,y2]、[x3,y3]....[xn,yn]都是可以合并为一个的)，那么合并后的区间的左端点一定是x1,而合并后的右端点一定是max(y1,y2,y3,.....yn)
            end = max(end, intervals[i][1])
            
        }
    }
    // 最后把这几个区间加进来
    ans = append(ans, []int{start, end})
    return ans
    
}