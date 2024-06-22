func setZeroes(matrix [][]int) {
    // 首先是初始化好我们的行标记数组
    m := len(matrix)
    row := make([]bool, m) // 初始化一个长度为m的标记数组
    // 初始化好我们的列标记数组
    n := len(matrix[0])
    col := make([]bool, n)
    
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if matrix[i][j] == 0{
                row[i] = true
                col[i] = true
            }
        }
    }
    // 然后就可以直接赋值修改
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if row[i] || col[j] {
                matrix[i][j] = 0
            }
        }
    }
}