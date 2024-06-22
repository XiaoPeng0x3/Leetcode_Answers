func spiralOrder(matrix [][]int) []int {
    // 总共需要 m * n 个元素
    m := len(matrix) // 行数
    n := len(matrix[0]) // 列数
    rowStart, rowEnd, colStart, colEnd := 0, m - 1, 0, n - 1
    count := 0
    ans := []int{}
    for count < m * n{
        // 横着开始
        for i := colStart; i <= colEnd; i++ {
            ans = append(ans, matrix[rowStart][i])
            count++
        }
        // 第一行读取完毕
        rowStart++
        // 注意要多加一层判断，判断是否越界
        if rowStart > rowEnd{
            break
        }

        // 竖着最后一列
        for i := rowStart; i <= rowEnd; i++ {
            ans = append(ans, matrix[i][colEnd])
            count++
        }
        colEnd--
        if colStart > colEnd{
            break
        }

        // 最后一行
        for i := colEnd; i >= colStart; i-- {
            ans = append(ans, matrix[rowEnd][i])
            count++
        }
        rowEnd--

        // 竖着读第一列
        for i := rowEnd; i >= rowStart; i-- {
            ans = append(ans, matrix[i][colStart])
            count++
        }
        colStart++
    }
    return ans
}