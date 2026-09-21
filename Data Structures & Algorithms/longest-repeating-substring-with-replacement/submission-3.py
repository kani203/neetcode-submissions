class Solution:
    def characterReplacement(self, s: str, k: int) -> int:
        cnt = [0] * 26
        l = 0
        maxLength = 0
        for r in range(len(s)):
            cnt[ord(s[r]) - 65] += 1
            while (r-l+1) - max(cnt) > k:
                cnt[ord(s[l]) - 65] -= 1
                l += 1
            maxLength = max(maxLength, (r-l+1))
        return maxLength
        