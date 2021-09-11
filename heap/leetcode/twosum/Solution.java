package algorithms.heap.leetcode.twosum;

import java.util.HashMap;
import java.util.stream.IntStream;

class Solution {
    public int[] twoSum(int[] nums, int target) {
        HashMap<Integer, Integer> hashMap = new HashMap<>();
        int[] result = new int[2];
        IntStream.range(0, nums.length)
                .peek(index -> {
                            Integer lookingIndex = hashMap.get(target - nums[index]);
                            if (lookingIndex != null) {
                                result[0] = lookingIndex;
                                result[1] = index;
                            } else {
                                hashMap.put(nums[index], index);
                            }
                        }
                )
                .count();
        return result;
    }

    public static void main(String[] args) throws Exception {
        int a = 3;
    }
}