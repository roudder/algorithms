class Solution:
    max = []

    def longestPalindrome(self, s: str) -> str:
        for index in range(0, len(s)):
            temp = []
            d = Solution.define_repeating_middle(index, s)
            start_repeating_index_left = index - d["left"]
            start_repeating_index_right = index + d["right"]
            count_of_adding = start_repeating_index_right - start_repeating_index_left + 1
            for _ in range(count_of_adding):
                temp.append(s[index])
            while start_repeating_index_left > 0 and start_repeating_index_right < len(s) - 1 and s[
                start_repeating_index_left - 1] == s[start_repeating_index_right + 1]:
                temp.insert(0, s[start_repeating_index_left - 1])
                temp.insert(len(temp), s[start_repeating_index_right + 1])
                start_repeating_index_left = start_repeating_index_left - 1
                start_repeating_index_right = start_repeating_index_right + 1
            if len(temp) > len(Solution.max):
                Solution.max = temp
        result = ''.join(Solution.max)
        return result

    @staticmethod
    def define_repeating_middle(index, s) -> dict:
        primary_index = index
        left_count = int()
        right_count = int()
        while index > 0 and s[index] == s[index - 1]:
            left_count = left_count + 1
            index = index - 1
        index = primary_index
        while index < len(s) - 1 and s[index] == s[index + 1]:
            right_count = right_count + 1
            index = index + 1
        d = dict()
        d["left"] = left_count
        d["right"] = right_count
        return d

