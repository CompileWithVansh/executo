-- ─────────────────────────────────────────────
--  Seed Data: 3 Sample Problems
-- ─────────────────────────────────────────────

-- Clear existing data (safe for development)
TRUNCATE TABLE submissions RESTART IDENTITY CASCADE;
TRUNCATE TABLE problems RESTART IDENTITY CASCADE;

-- ── Problem 1: Two Sum (Easy) ─────────────────

INSERT INTO problems (
    title, slug, description, difficulty,
    examples, constraints, test_cases, function_signature
) VALUES (
    'Two Sum',
    'two-sum',
    'Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.',
    'easy',
    '[
        {"input": "nums = [2,7,11,15], target = 9", "output": "[0,1]", "explanation": "Because nums[0] + nums[1] == 9, we return [0, 1]."},
        {"input": "nums = [3,2,4], target = 6", "output": "[1,2]"},
        {"input": "nums = [3,3], target = 6", "output": "[0,1]"}
    ]',
    '["2 <= nums.length <= 10^4", "-10^9 <= nums[i] <= 10^9", "-10^9 <= target <= 10^9", "Only one valid answer exists."]',
    '[
        {"input": "4\n2 7 11 15\n9", "expected_output": "0 1"},
        {"input": "3\n3 2 4\n6", "expected_output": "1 2"},
        {"input": "2\n3 3\n6", "expected_output": "0 1"},
        {"input": "5\n1 2 3 4 5\n9", "expected_output": "3 4"},
        {"input": "2\n-1 -2\n-3", "expected_output": "0 1"}
    ]',
    '{
        "python3": "class Solution:\n    def twoSum(self, nums: list[int], target: int) -> list[int]:\n        # Your code here\n        pass",
        "java": "class Solution {\n    public int[] twoSum(int[] nums, int target) {\n        // Your code here\n        return new int[]{};\n    }\n}",
        "cpp": "#include <bits/stdc++.h>\nusing namespace std;\n\nclass Solution {\npublic:\n    vector<int> twoSum(vector<int>& nums, int target) {\n        // Your code here\n        return {};\n    }\n};",
        "javascript": "/**\n * @param {number[]} nums\n * @param {number} target\n * @return {number[]}\n */\nvar twoSum = function(nums, target) {\n    // Your code here\n};"
    }'
);

-- ── Problem 2: Valid Parentheses (Easy) ──────

INSERT INTO problems (
    title, slug, description, difficulty,
    examples, constraints, test_cases, function_signature
) VALUES (
    'Valid Parentheses',
    'valid-parentheses',
    'Given a string s containing just the characters ''('', '')'', ''{'', ''}'', ''['' and '']'', determine if the input string is valid.

An input string is valid if:
1. Open brackets must be closed by the same type of brackets.
2. Open brackets must be closed in the correct order.
3. Every close bracket has a corresponding open bracket of the same type.',
    'easy',
    '[
        {"input": "s = \"()\"", "output": "true"},
        {"input": "s = \"()[]{}\"", "output": "true"},
        {"input": "s = \"(]\"", "output": "false"},
        {"input": "s = \"([)]\"", "output": "false"},
        {"input": "s = \"{[]}\"", "output": "true"}
    ]',
    '["1 <= s.length <= 10^4", "s consists of parentheses only ''()[]{}''."]',
    '[
        {"input": "()", "expected_output": "true"},
        {"input": "()[]{}", "expected_output": "true"},
        {"input": "(]", "expected_output": "false"},
        {"input": "([)]", "expected_output": "false"},
        {"input": "{[]}", "expected_output": "true"},
        {"input": "", "expected_output": "true"},
        {"input": "((", "expected_output": "false"},
        {"input": "]", "expected_output": "false"}
    ]',
    '{
        "python3": "class Solution:\n    def isValid(self, s: str) -> bool:\n        # Your code here\n        pass",
        "java": "class Solution {\n    public boolean isValid(String s) {\n        // Your code here\n        return false;\n    }\n}",
        "cpp": "#include <bits/stdc++.h>\nusing namespace std;\n\nclass Solution {\npublic:\n    bool isValid(string s) {\n        // Your code here\n        return false;\n    }\n};",
        "javascript": "/**\n * @param {string} s\n * @return {boolean}\n */\nvar isValid = function(s) {\n    // Your code here\n};"
    }'
);

-- ── Problem 3: Maximum Subarray (Medium) ─────

INSERT INTO problems (
    title, slug, description, difficulty,
    examples, constraints, test_cases, function_signature
) VALUES (
    'Maximum Subarray',
    'maximum-subarray',
    'Given an integer array nums, find the subarray with the largest sum, and return its sum.

A subarray is a contiguous non-empty sequence of elements within an array.',
    'medium',
    '[
        {"input": "nums = [-2,1,-3,4,-1,2,1,-5,4]", "output": "6", "explanation": "The subarray [4,-1,2,1] has the largest sum 6."},
        {"input": "nums = [1]", "output": "1"},
        {"input": "nums = [5,4,-1,7,8]", "output": "23"}
    ]',
    '["1 <= nums.length <= 10^5", "-10^4 <= nums[i] <= 10^4"]',
    '[
        {"input": "9\n-2 1 -3 4 -1 2 1 -5 4", "expected_output": "6"},
        {"input": "1\n1", "expected_output": "1"},
        {"input": "5\n5 4 -1 7 8", "expected_output": "23"},
        {"input": "4\n-1 -2 -3 -4", "expected_output": "-1"},
        {"input": "6\n1 2 3 -2 5 -1", "expected_output": "9"},
        {"input": "3\n-2 -1 -3", "expected_output": "-1"}
    ]',
    '{
        "python3": "class Solution:\n    def maxSubArray(self, nums: list[int]) -> int:\n        # Your code here\n        pass",
        "java": "class Solution {\n    public int maxSubArray(int[] nums) {\n        // Your code here\n        return 0;\n    }\n}",
        "cpp": "#include <bits/stdc++.h>\nusing namespace std;\n\nclass Solution {\npublic:\n    int maxSubArray(vector<int>& nums) {\n        // Your code here\n        return 0;\n    }\n};",
        "javascript": "/**\n * @param {number[]} nums\n * @return {number}\n */\nvar maxSubArray = function(nums) {\n    // Your code here\n};"
    }'
);

-- Verify seed
SELECT id, title, difficulty FROM problems ORDER BY id;
