func sortedArrayToBST(nums []int) *TreeNode {
    n := len(nums)
    return build(nums, 0, n - 1)
}

func build(nums[] int, l int, r int) *TreeNode {
    if l > r {
        return nil
    }
    m := (l + r) / 2
    // 构造出来根节点
    root := &TreeNode{Val:nums[m]}
    root.Left = build(nums, l, m - 1)
    root.Right = build(nums, m + 1, r)
    return root
}