# https://leetcode.com/explore/featured/card/30-day-leetcoding-challenge/528/week-1/3285/

class Solution:
    def maxSubArray(self, nums):
        if len(nums) == 1:
            return nums[0]
        l = 0
        r = 1
        max_till_now = nums[0]
        sliding_sum = nums[0]
        while r < len(nums):
            print(l, r, max_till_now, sliding_sum)
            next_num = nums[r]
            if sliding_sum + next_num < sliding_sum:
                while l <= r:
                    last_num = nums[l]
                    sliding_sum -= last_num
                    max_till_now = max(max_till_now, sliding_sum)
                    l += 1
            sliding_sum += next_num
            max_till_now = max(max_till_now, sliding_sum)
            r += 1
        return max_till_now