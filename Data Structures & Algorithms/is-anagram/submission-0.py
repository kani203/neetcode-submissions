class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        if len(s) != len(t):
            return False
        chars = {}
        for i in range(len(s)):
            chars[s[i]] = chars.get(s[i], 0) + 1
        for i in range(len(t)):
            if (not chars.get(t[i])) or chars.get(t[i]) == 0:
                return False
            else:
                chars[t[i]] -= 1
        return True