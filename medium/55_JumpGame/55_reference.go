package leetcode

/*
解法: 贪心
只要存在一个位置 x，它本身可以到达，并且它跳跃的最大长度为 x+nums[x]，这个值大于等于 y ，那么位置 y 也可以到达。
换句话说，对于每一个可以到达的位置 x，它使得 x+1,x+2,⋯,x+nums[x] 这些连续的位置都可以到达。
这样以来，我们依次遍历数组中的每一个位置，并实时维护 最远可以到达的位置。
对于当前遍历到的位置 x，如果它在 最远可以到达的位置 的范围内，那么我们就可以从起点通过若干次跳跃到达该位置，因此我们可以用 x+nums[x] 更新 最远可以到达的位置。
在遍历的过程中，如果 最远可以到达的位置 大于等于数组中的最后一个位置，那就说明最后一个位置可达，我们就可以直接返回 True 作为答案。反之，如果在遍历结束后，最后一个位置仍然不可达，我们就返回 False 作为答案。


结果: 执行用时 :12 ms 内存消耗 :4.2 MB
*/


func referenceCanJump(nums []int) bool {
	var length, maxIndex = len(nums), 0
	for i := 0; i < length; i++ {
		if i <= maxIndex {
			if maxIndex < i + nums[i] {
				maxIndex = i + nums[i]
			}
			if maxIndex >= length - 1 {
				return true
			}
		}
	}
	return false
}