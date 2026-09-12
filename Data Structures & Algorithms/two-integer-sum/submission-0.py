class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        occurance = {}
        for i, num in enumerate(nums):
            if num in occurance.keys():
                occurance[num].append(i)
            else:
                occurance[num] = [i]

        for i, num in enumerate(nums):
            if target - num == num:
                if len(occurance[num]) == 2:
                    return occurance[num]
            else:
                if (target - num) in occurance.keys():
                    return [i, occurance[target - num][0]]


