func searchMatrix(matrix [][]int, target int) bool {
    m := len(matrix)
    n := len(matrix[0])
    l, r := 0, m * n - 1
    x, mid := 0, 0
    
    for l <= r {
        mid = l + (r - l) / 2
        x = matrix[mid / n][mid % n]
        if x == target {
            return true
        } else if x < target {
            l = mid + 1
        } else {
            r = mid - 1
        }
    }
    return false
    
}