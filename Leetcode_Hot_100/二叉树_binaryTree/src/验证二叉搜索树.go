// 确定好区间，所有的值都在两个范围之间
func isValidBST(root *TreeNode) bool {
    return dfs(root, math.MinInt, math.MaxInt)
}

func dfs(root *TreeNode, l int, r int) bool {
    if root == nil {
        return true
    }
    // 根节点的值大于做左区间，小于右区间
    return l < root.Val && r > root.Val && dfs(root.Left, l, root.Val) && dfs(root.Right, root.Val, r)
    
    // 除此之外，还要检查左右子树，检查左子树的时候，根节点就是区间的右端点
    // 检查右子树的时候，根节点就是区间的左端点
}