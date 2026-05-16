-- ─────────────────────────────────────────────
--  Seed Data: Sample Problems for Executo
-- ─────────────────────────────────────────────

INSERT INTO problems (title, slug, description, difficulty, examples, constraints, test_cases, function_signature, lc_number, lc_url)
VALUES
(
    'Two Sum',
    'two-sum',
    'Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target. You may assume that each input would have exactly one solution, and you may not use the same element twice. You can return the answer in any order.',
    'easy',
    '[{"input": "nums = [2,7,11,15], target = 9", "output": "[0,1]", "explanation": "Because nums[0] + nums[1] == 9, we return [0, 1]."}, {"input": "nums = [3,2,4], target = 6", "output": "[1,2]"}, {"input": "nums = [3,3], target = 6", "output": "[0,1]"}]',
    '["2 <= nums.length <= 10^4", "-10^9 <= nums[i] <= 10^9", "-10^9 <= target <= 10^9", "Only one valid answer exists."]',
    '[{"input": "4\n2 7 11 15\n9", "expected_output": "0 1"}, {"input": "3\n3 2 4\n6", "expected_output": "1 2"}, {"input": "2\n3 3\n6", "expected_output": "0 1"}]',
    '{"python3": "class Solution:\n    def twoSum(self, nums, target):\n        pass", "java": "class Solution {\n    public int[] twoSum(int[] nums, int target) {\n    }\n}", "cpp": "class Solution {\npublic:\n    vector<int> twoSum(vector<int>& nums, int target) {\n    }\n};", "javascript": "var twoSum = function(nums, target) {\n};"}',
    1,
    'https://leetcode.com/problems/two-sum/'
),
(
    'Valid Parentheses',
    'valid-parentheses',
    'Given a string s containing just the characters ''('', '')'', ''{'', ''}'', ''['' and '']'', determine if the input string is valid. An input string is valid if: Open brackets must be closed by the same type of brackets. Open brackets must be closed in the correct order. Every close bracket has a corresponding open bracket of the same type.',
    'easy',
    '[{"input": "s = \"()\"", "output": "true"}, {"input": "s = \"()[]{}\"", "output": "true"}, {"input": "s = \"(]\"", "output": "false"}]',
    '["1 <= s.length <= 10^4", "s consists of parentheses only ''()[]{}''"]',
    '[{"input": "()", "expected_output": "true"}, {"input": "()[]{}", "expected_output": "true"}, {"input": "(]", "expected_output": "false"}, {"input": "([)]", "expected_output": "false"}, {"input": "{[]}", "expected_output": "true"}]',
    '{"python3": "class Solution:\n    def isValid(self, s: str) -> bool:\n        pass", "java": "class Solution {\n    public boolean isValid(String s) {\n    }\n}", "cpp": "class Solution {\npublic:\n    bool isValid(string s) {\n    }\n};", "javascript": "var isValid = function(s) {\n};"}',
    20,
    'https://leetcode.com/problems/valid-parentheses/'
),
(
    'Merge Two Sorted Lists',
    'merge-two-sorted-lists',
    'You are given the heads of two sorted linked lists list1 and list2. Merge the two lists into one sorted list. The list should be made by splicing together the nodes of the first two lists. Return the head of the merged linked list.',
    'easy',
    '[{"input": "list1 = [1,2,4], list2 = [1,3,4]", "output": "[1,1,2,3,4,4]"}, {"input": "list1 = [], list2 = []", "output": "[]"}, {"input": "list1 = [], list2 = [0]", "output": "[0]"}]',
    '["The number of nodes in both lists is in the range [0, 50]", "-100 <= Node.val <= 100", "Both list1 and list2 are sorted in non-decreasing order."]',
    '[{"input": "3\n1 2 4\n3\n1 3 4", "expected_output": "1 1 2 3 4 4"}, {"input": "0\n\n0\n", "expected_output": ""}, {"input": "0\n\n1\n0", "expected_output": "0"}]',
    '{"python3": "class Solution:\n    def mergeTwoLists(self, list1, list2):\n        pass", "java": "class Solution {\n    public ListNode mergeTwoLists(ListNode list1, ListNode list2) {\n    }\n}", "cpp": "class Solution {\npublic:\n    ListNode* mergeTwoLists(ListNode* list1, ListNode* list2) {\n    }\n};", "javascript": "var mergeTwoLists = function(list1, list2) {\n};"}',
    21,
    'https://leetcode.com/problems/merge-two-sorted-lists/'
),
(
    'Best Time to Buy and Sell Stock',
    'best-time-to-buy-and-sell-stock',
    'You are given an array prices where prices[i] is the price of a given stock on the ith day. You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock. Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.',
    'easy',
    '[{"input": "prices = [7,1,5,3,6,4]", "output": "5", "explanation": "Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5."}, {"input": "prices = [7,6,4,3,1]", "output": "0", "explanation": "No transactions are done and the max profit = 0."}]',
    '["1 <= prices.length <= 10^5", "0 <= prices[i] <= 10^4"]',
    '[{"input": "6\n7 1 5 3 6 4", "expected_output": "5"}, {"input": "5\n7 6 4 3 1", "expected_output": "0"}, {"input": "3\n2 4 1", "expected_output": "2"}]',
    '{"python3": "class Solution:\n    def maxProfit(self, prices) -> int:\n        pass", "java": "class Solution {\n    public int maxProfit(int[] prices) {\n    }\n}", "cpp": "class Solution {\npublic:\n    int maxProfit(vector<int>& prices) {\n    }\n};", "javascript": "var maxProfit = function(prices) {\n};"}',
    121,
    'https://leetcode.com/problems/best-time-to-buy-and-sell-stock/'
),
(
    'Contains Duplicate',
    'contains-duplicate',
    'Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.',
    'easy',
    '[{"input": "nums = [1,2,3,1]", "output": "true"}, {"input": "nums = [1,2,3,4]", "output": "false"}, {"input": "nums = [1,1,1,3,3,4,3,2,4,2]", "output": "true"}]',
    '["1 <= nums.length <= 10^5", "-10^9 <= nums[i] <= 10^9"]',
    '[{"input": "4\n1 2 3 1", "expected_output": "true"}, {"input": "4\n1 2 3 4", "expected_output": "false"}, {"input": "10\n1 1 1 3 3 4 3 2 4 2", "expected_output": "true"}]',
    '{"python3": "class Solution:\n    def containsDuplicate(self, nums) -> bool:\n        pass", "java": "class Solution {\n    public boolean containsDuplicate(int[] nums) {\n    }\n}", "cpp": "class Solution {\npublic:\n    bool containsDuplicate(vector<int>& nums) {\n    }\n};", "javascript": "var containsDuplicate = function(nums) {\n};"}',
    217,
    'https://leetcode.com/problems/contains-duplicate/'
)
ON CONFLICT (slug) DO NOTHING;
