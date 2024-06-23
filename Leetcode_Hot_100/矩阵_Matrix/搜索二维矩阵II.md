# 搜索二维矩阵II

## 描述

编写一个高效的算法来搜索 `*m* x *n*` 矩阵 `matrix` 中的一个目标值 `target` 。该矩阵具有以下特性：

- 每行的元素从左到右升序排列。
- 每列的元素从上到下升序排列。

 

**示例 1：**

![img](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2020/11/25/searchgrid2.jpg)

```
输入：matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]], target = 5
输出：true
```

**示例 2：**

![img](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2020/11/25/searchgrid.jpg)

```
输入：matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]], target = 20
输出：false
```

 

**提示：**

- `m == matrix.length`
- `n == matrix[i].length`
- `1 <= n, m <= 300`
- `-109 <= matrix[i][j] <= 109`
- 每行的所有元素从左到右升序排列
- 每列的所有元素从上到下升序排列
- `-109 <= target <= 109`



## 思路

首先要利用好已经排好序的这个特征，对于右上角(也可以是其他参考值)来说，如果`target`大于右上角的元素，那么右上角所在的那一行都不在寻找范围之内；同理，如果小于右上角元素，那么所在的那一列也不在我们的寻找范围之内。这样，我们每次通过简单的比较就可以省去很多无用的元素



## 代码

```go
func searchMatrix(matrix [][]int, target int) bool {
    
    // 从右上角开始
    m, n := len(matrix), len(matrix[0])
    i, j := 0, n - 1 // 从右上角元素开始看
    for i < m && j >= 0 {
        if matrix[i][j] == target {
            return true
        } else if matrix[i][j] < target {
            i++
        } else {
            j--
        }
    }
    return false

}
```

