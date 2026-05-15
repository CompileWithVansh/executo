-- Migration 004: Seed NeetCode 150 problems
-- All 150 problems with LC numbers and links
-- No descriptions or test cases — users link out to LC for problem statements

TRUNCATE TABLE suggested_testcases RESTART IDENTITY CASCADE;
TRUNCATE TABLE submissions RESTART IDENTITY CASCADE;
TRUNCATE TABLE problems RESTART IDENTITY CASCADE;

-- Arrays & Hashing
INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Contains Duplicate', 'contains-duplicate', 'easy', 217, 'https://leetcode.com/problems/contains-duplicate/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Valid Anagram', 'valid-anagram', 'easy', 242, 'https://leetcode.com/problems/valid-anagram/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Two Sum', 'two-sum', 'easy', 1, 'https://leetcode.com/problems/two-sum/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Group Anagrams', 'group-anagrams', 'medium', 49, 'https://leetcode.com/problems/group-anagrams/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Top K Frequent Elements', 'top-k-frequent-elements', 'medium', 347, 'https://leetcode.com/problems/top-k-frequent-elements/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Product of Array Except Self', 'product-of-array-except-self', 'medium', 238, 'https://leetcode.com/problems/product-of-array-except-self/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Valid Sudoku', 'valid-sudoku', 'medium', 36, 'https://leetcode.com/problems/valid-sudoku/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Encode and Decode Strings', 'encode-and-decode-strings', 'medium', 659, 'https://leetcode.com/problems/encode-and-decode-strings/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Longest Consecutive Sequence', 'longest-consecutive-sequence', 'medium', 128, 'https://leetcode.com/problems/longest-consecutive-sequence/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

-- Two Pointers
INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Valid Palindrome', 'valid-palindrome', 'easy', 125, 'https://leetcode.com/problems/valid-palindrome/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Two Sum II Input Array Is Sorted', 'two-sum-ii-input-array-is-sorted', 'medium', 167, 'https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('3Sum', '3sum', 'medium', 15, 'https://leetcode.com/problems/3sum/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Container With Most Water', 'container-with-most-water', 'medium', 11, 'https://leetcode.com/problems/container-with-most-water/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Trapping Rain Water', 'trapping-rain-water', 'hard', 42, 'https://leetcode.com/problems/trapping-rain-water/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

-- Sliding Window
INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Best Time to Buy and Sell Stock', 'best-time-to-buy-and-sell-stock', 'easy', 121, 'https://leetcode.com/problems/best-time-to-buy-and-sell-stock/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Longest Substring Without Repeating Characters', 'longest-substring-without-repeating-characters', 'medium', 3, 'https://leetcode.com/problems/longest-substring-without-repeating-characters/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Longest Repeating Character Replacement', 'longest-repeating-character-replacement', 'medium', 424, 'https://leetcode.com/problems/longest-repeating-character-replacement/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Permutation in String', 'permutation-in-string', 'medium', 567, 'https://leetcode.com/problems/permutation-in-string/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Minimum Window Substring', 'minimum-window-substring', 'hard', 76, 'https://leetcode.com/problems/minimum-window-substring/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Sliding Window Maximum', 'sliding-window-maximum', 'hard', 239, 'https://leetcode.com/problems/sliding-window-maximum/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

-- Stack
INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Valid Parentheses', 'valid-parentheses', 'easy', 20, 'https://leetcode.com/problems/valid-parentheses/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Min Stack', 'min-stack', 'medium', 155, 'https://leetcode.com/problems/min-stack/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Evaluate Reverse Polish Notation', 'evaluate-reverse-polish-notation', 'medium', 150, 'https://leetcode.com/problems/evaluate-reverse-polish-notation/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Generate Parentheses', 'generate-parentheses', 'medium', 22, 'https://leetcode.com/problems/generate-parentheses/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Daily Temperatures', 'daily-temperatures', 'medium', 739, 'https://leetcode.com/problems/daily-temperatures/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Car Fleet', 'car-fleet', 'medium', 853, 'https://leetcode.com/problems/car-fleet/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Largest Rectangle in Histogram', 'largest-rectangle-in-histogram', 'hard', 84, 'https://leetcode.com/problems/largest-rectangle-in-histogram/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

-- Binary Search
INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Binary Search', 'binary-search', 'easy', 704, 'https://leetcode.com/problems/binary-search/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Search a 2D Matrix', 'search-a-2d-matrix', 'medium', 74, 'https://leetcode.com/problems/search-a-2d-matrix/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Koko Eating Bananas', 'koko-eating-bananas', 'medium', 875, 'https://leetcode.com/problems/koko-eating-bananas/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Find Minimum in Rotated Sorted Array', 'find-minimum-in-rotated-sorted-array', 'medium', 153, 'https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Search in Rotated Sorted Array', 'search-in-rotated-sorted-array', 'medium', 33, 'https://leetcode.com/problems/search-in-rotated-sorted-array/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Time Based Key Value Store', 'time-based-key-value-store', 'medium', 981, 'https://leetcode.com/problems/time-based-key-value-store/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Median of Two Sorted Arrays', 'median-of-two-sorted-arrays', 'hard', 4, 'https://leetcode.com/problems/median-of-two-sorted-arrays/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

-- Linked List
INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Reverse Linked List', 'reverse-linked-list', 'easy', 206, 'https://leetcode.com/problems/reverse-linked-list/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Merge Two Sorted Lists', 'merge-two-sorted-lists', 'easy', 21, 'https://leetcode.com/problems/merge-two-sorted-lists/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Reorder List', 'reorder-list', 'medium', 143, 'https://leetcode.com/problems/reorder-list/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Remove Nth Node From End of List', 'remove-nth-node-from-end-of-list', 'medium', 19, 'https://leetcode.com/problems/remove-nth-node-from-end-of-list/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Copy List With Random Pointer', 'copy-list-with-random-pointer', 'medium', 138, 'https://leetcode.com/problems/copy-list-with-random-pointer/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Add Two Numbers', 'add-two-numbers', 'medium', 2, 'https://leetcode.com/problems/add-two-numbers/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Linked List Cycle', 'linked-list-cycle', 'easy', 141, 'https://leetcode.com/problems/linked-list-cycle/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Find the Duplicate Number', 'find-the-duplicate-number', 'medium', 287, 'https://leetcode.com/problems/find-the-duplicate-number/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('LRU Cache', 'lru-cache', 'medium', 146, 'https://leetcode.com/problems/lru-cache/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Merge K Sorted Lists', 'merge-k-sorted-lists', 'hard', 23, 'https://leetcode.com/problems/merge-k-sorted-lists/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Reverse Nodes in K Group', 'reverse-nodes-in-k-group', 'hard', 25, 'https://leetcode.com/problems/reverse-nodes-in-k-group/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

-- Trees
INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Invert Binary Tree', 'invert-binary-tree', 'easy', 226, 'https://leetcode.com/problems/invert-binary-tree/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Maximum Depth of Binary Tree', 'maximum-depth-of-binary-tree', 'easy', 104, 'https://leetcode.com/problems/maximum-depth-of-binary-tree/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Diameter of Binary Tree', 'diameter-of-binary-tree', 'easy', 543, 'https://leetcode.com/problems/diameter-of-binary-tree/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Balanced Binary Tree', 'balanced-binary-tree', 'easy', 110, 'https://leetcode.com/problems/balanced-binary-tree/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Same Tree', 'same-tree', 'easy', 100, 'https://leetcode.com/problems/same-tree/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Subtree of Another Tree', 'subtree-of-another-tree', 'easy', 572, 'https://leetcode.com/problems/subtree-of-another-tree/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Lowest Common Ancestor of a Binary Search Tree', 'lowest-common-ancestor-of-a-binary-search-tree', 'medium', 235, 'https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Binary Tree Level Order Traversal', 'binary-tree-level-order-traversal', 'medium', 102, 'https://leetcode.com/problems/binary-tree-level-order-traversal/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Binary Tree Right Side View', 'binary-tree-right-side-view', 'medium', 199, 'https://leetcode.com/problems/binary-tree-right-side-view/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Count Good Nodes in Binary Tree', 'count-good-nodes-in-binary-tree', 'medium', 1448, 'https://leetcode.com/problems/count-good-nodes-in-binary-tree/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Validate Binary Search Tree', 'validate-binary-search-tree', 'medium', 98, 'https://leetcode.com/problems/validate-binary-search-tree/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Kth Smallest Element in a BST', 'kth-smallest-element-in-a-bst', 'medium', 230, 'https://leetcode.com/problems/kth-smallest-element-in-a-bst/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Construct Binary Tree From Preorder and Inorder Traversal', 'construct-binary-tree-from-preorder-and-inorder-traversal', 'medium', 105, 'https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Binary Tree Maximum Path Sum', 'binary-tree-maximum-path-sum', 'hard', 124, 'https://leetcode.com/problems/binary-tree-maximum-path-sum/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Serialize and Deserialize Binary Tree', 'serialize-and-deserialize-binary-tree', 'hard', 297, 'https://leetcode.com/problems/serialize-and-deserialize-binary-tree/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

-- Tries
INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Implement Trie Prefix Tree', 'implement-trie-prefix-tree', 'medium', 208, 'https://leetcode.com/problems/implement-trie-prefix-tree/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Design Add and Search Words Data Structure', 'design-add-and-search-words-data-structure', 'medium', 211, 'https://leetcode.com/problems/design-add-and-search-words-data-structure/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Word Search II', 'word-search-ii', 'hard', 212, 'https://leetcode.com/problems/word-search-ii/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

-- Heap / Priority Queue
INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Kth Largest Element in a Stream', 'kth-largest-element-in-a-stream', 'easy', 703, 'https://leetcode.com/problems/kth-largest-element-in-a-stream/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Last Stone Weight', 'last-stone-weight', 'easy', 1046, 'https://leetcode.com/problems/last-stone-weight/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('K Closest Points to Origin', 'k-closest-points-to-origin', 'medium', 973, 'https://leetcode.com/problems/k-closest-points-to-origin/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Kth Largest Element in an Array', 'kth-largest-element-in-an-array', 'medium', 215, 'https://leetcode.com/problems/kth-largest-element-in-an-array/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Task Scheduler', 'task-scheduler', 'medium', 621, 'https://leetcode.com/problems/task-scheduler/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Design Twitter', 'design-twitter', 'medium', 355, 'https://leetcode.com/problems/design-twitter/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Find Median From Data Stream', 'find-median-from-data-stream', 'hard', 295, 'https://leetcode.com/problems/find-median-from-data-stream/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

-- Backtracking
INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Subsets', 'subsets', 'medium', 78, 'https://leetcode.com/problems/subsets/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Combination Sum', 'combination-sum', 'medium', 39, 'https://leetcode.com/problems/combination-sum/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Combination Sum II', 'combination-sum-ii', 'medium', 40, 'https://leetcode.com/problems/combination-sum-ii/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Permutations', 'permutations', 'medium', 46, 'https://leetcode.com/problems/permutations/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Subsets II', 'subsets-ii', 'medium', 90, 'https://leetcode.com/problems/subsets-ii/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Word Search', 'word-search', 'medium', 79, 'https://leetcode.com/problems/word-search/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Palindrome Partitioning', 'palindrome-partitioning', 'medium', 131, 'https://leetcode.com/problems/palindrome-partitioning/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Letter Combinations of a Phone Number', 'letter-combinations-of-a-phone-number', 'medium', 17, 'https://leetcode.com/problems/letter-combinations-of-a-phone-number/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('N Queens', 'n-queens', 'hard', 51, 'https://leetcode.com/problems/n-queens/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

-- Graphs
INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Number of Islands', 'number-of-islands', 'medium', 200, 'https://leetcode.com/problems/number-of-islands/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Clone Graph', 'clone-graph', 'medium', 133, 'https://leetcode.com/problems/clone-graph/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Max Area of Island', 'max-area-of-island', 'medium', 695, 'https://leetcode.com/problems/max-area-of-island/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Pacific Atlantic Water Flow', 'pacific-atlantic-water-flow', 'medium', 417, 'https://leetcode.com/problems/pacific-atlantic-water-flow/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Surrounded Regions', 'surrounded-regions', 'medium', 130, 'https://leetcode.com/problems/surrounded-regions/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Rotting Oranges', 'rotting-oranges', 'medium', 994, 'https://leetcode.com/problems/rotting-oranges/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Walls and Gates', 'walls-and-gates', 'medium', 286, 'https://leetcode.com/problems/walls-and-gates/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Course Schedule', 'course-schedule', 'medium', 207, 'https://leetcode.com/problems/course-schedule/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Course Schedule II', 'course-schedule-ii', 'medium', 210, 'https://leetcode.com/problems/course-schedule-ii/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Redundant Connection', 'redundant-connection', 'medium', 684, 'https://leetcode.com/problems/redundant-connection/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Number of Connected Components in an Undirected Graph', 'number-of-connected-components-in-an-undirected-graph', 'medium', 323, 'https://leetcode.com/problems/number-of-connected-components-in-an-undirected-graph/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Graph Valid Tree', 'graph-valid-tree', 'medium', 261, 'https://leetcode.com/problems/graph-valid-tree/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Word Ladder', 'word-ladder', 'hard', 127, 'https://leetcode.com/problems/word-ladder/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

-- Advanced Graphs
INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Reconstruct Itinerary', 'reconstruct-itinerary', 'hard', 332, 'https://leetcode.com/problems/reconstruct-itinerary/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Min Cost to Connect All Points', 'min-cost-to-connect-all-points', 'medium', 1584, 'https://leetcode.com/problems/min-cost-to-connect-all-points/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Network Delay Time', 'network-delay-time', 'medium', 743, 'https://leetcode.com/problems/network-delay-time/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Swim in Rising Water', 'swim-in-rising-water', 'hard', 778, 'https://leetcode.com/problems/swim-in-rising-water/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Alien Dictionary', 'alien-dictionary', 'hard', 269, 'https://leetcode.com/problems/alien-dictionary/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Cheapest Flights Within K Stops', 'cheapest-flights-within-k-stops', 'medium', 787, 'https://leetcode.com/problems/cheapest-flights-within-k-stops/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

-- 1-D Dynamic Programming
INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Climbing Stairs', 'climbing-stairs', 'easy', 70, 'https://leetcode.com/problems/climbing-stairs/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Min Cost Climbing Stairs', 'min-cost-climbing-stairs', 'easy', 746, 'https://leetcode.com/problems/min-cost-climbing-stairs/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('House Robber', 'house-robber', 'medium', 198, 'https://leetcode.com/problems/house-robber/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('House Robber II', 'house-robber-ii', 'medium', 213, 'https://leetcode.com/problems/house-robber-ii/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Longest Palindromic Substring', 'longest-palindromic-substring', 'medium', 5, 'https://leetcode.com/problems/longest-palindromic-substring/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Palindromic Substrings', 'palindromic-substrings', 'medium', 647, 'https://leetcode.com/problems/palindromic-substrings/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Decode Ways', 'decode-ways', 'medium', 91, 'https://leetcode.com/problems/decode-ways/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Coin Change', 'coin-change', 'medium', 322, 'https://leetcode.com/problems/coin-change/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Maximum Product Subarray', 'maximum-product-subarray', 'medium', 152, 'https://leetcode.com/problems/maximum-product-subarray/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Word Break', 'word-break', 'medium', 139, 'https://leetcode.com/problems/word-break/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Longest Increasing Subsequence', 'longest-increasing-subsequence', 'medium', 300, 'https://leetcode.com/problems/longest-increasing-subsequence/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Partition Equal Subset Sum', 'partition-equal-subset-sum', 'medium', 416, 'https://leetcode.com/problems/partition-equal-subset-sum/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

-- 2-D Dynamic Programming
INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Unique Paths', 'unique-paths', 'medium', 62, 'https://leetcode.com/problems/unique-paths/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Longest Common Subsequence', 'longest-common-subsequence', 'medium', 1143, 'https://leetcode.com/problems/longest-common-subsequence/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Best Time to Buy and Sell Stock With Cooldown', 'best-time-to-buy-and-sell-stock-with-cooldown', 'medium', 309, 'https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Coin Change II', 'coin-change-ii', 'medium', 518, 'https://leetcode.com/problems/coin-change-ii/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Target Sum', 'target-sum', 'medium', 494, 'https://leetcode.com/problems/target-sum/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Interleaving String', 'interleaving-string', 'medium', 97, 'https://leetcode.com/problems/interleaving-string/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Longest Increasing Path in a Matrix', 'longest-increasing-path-in-a-matrix', 'hard', 329, 'https://leetcode.com/problems/longest-increasing-path-in-a-matrix/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Distinct Subsequences', 'distinct-subsequences', 'hard', 115, 'https://leetcode.com/problems/distinct-subsequences/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Edit Distance', 'edit-distance', 'medium', 72, 'https://leetcode.com/problems/edit-distance/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Burst Balloons', 'burst-balloons', 'hard', 312, 'https://leetcode.com/problems/burst-balloons/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Regular Expression Matching', 'regular-expression-matching', 'hard', 10, 'https://leetcode.com/problems/regular-expression-matching/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

-- Greedy
INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Maximum Subarray', 'maximum-subarray', 'medium', 53, 'https://leetcode.com/problems/maximum-subarray/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Jump Game', 'jump-game', 'medium', 55, 'https://leetcode.com/problems/jump-game/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Jump Game II', 'jump-game-ii', 'medium', 45, 'https://leetcode.com/problems/jump-game-ii/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Gas Station', 'gas-station', 'medium', 134, 'https://leetcode.com/problems/gas-station/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Hand of Straights', 'hand-of-straights', 'medium', 846, 'https://leetcode.com/problems/hand-of-straights/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Merge Triplets to Form Target Triplet', 'merge-triplets-to-form-target-triplet', 'medium', 1899, 'https://leetcode.com/problems/merge-triplets-to-form-target-triplet/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Partition Labels', 'partition-labels', 'medium', 763, 'https://leetcode.com/problems/partition-labels/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Valid Parenthesis String', 'valid-parenthesis-string', 'medium', 678, 'https://leetcode.com/problems/valid-parenthesis-string/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

-- Intervals
INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Insert Interval', 'insert-interval', 'medium', 57, 'https://leetcode.com/problems/insert-interval/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Merge Intervals', 'merge-intervals', 'medium', 56, 'https://leetcode.com/problems/merge-intervals/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Non Overlapping Intervals', 'non-overlapping-intervals', 'medium', 435, 'https://leetcode.com/problems/non-overlapping-intervals/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Meeting Rooms', 'meeting-rooms', 'easy', 252, 'https://leetcode.com/problems/meeting-rooms/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Meeting Rooms II', 'meeting-rooms-ii', 'medium', 253, 'https://leetcode.com/problems/meeting-rooms-ii/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Minimum Interval to Include Each Query', 'minimum-interval-to-include-each-query', 'hard', 1851, 'https://leetcode.com/problems/minimum-interval-to-include-each-query/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

-- Math & Geometry
INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Rotate Image', 'rotate-image', 'medium', 48, 'https://leetcode.com/problems/rotate-image/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Spiral Matrix', 'spiral-matrix', 'medium', 54, 'https://leetcode.com/problems/spiral-matrix/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Set Matrix Zeroes', 'set-matrix-zeroes', 'medium', 73, 'https://leetcode.com/problems/set-matrix-zeroes/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Happy Number', 'happy-number', 'easy', 202, 'https://leetcode.com/problems/happy-number/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Plus One', 'plus-one', 'easy', 66, 'https://leetcode.com/problems/plus-one/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Pow(x, n)', 'powx-n', 'medium', 50, 'https://leetcode.com/problems/powx-n/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Multiply Strings', 'multiply-strings', 'medium', 43, 'https://leetcode.com/problems/multiply-strings/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Detect Squares', 'detect-squares', 'medium', 2013, 'https://leetcode.com/problems/detect-squares/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

-- Bit Manipulation
INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Single Number', 'single-number', 'easy', 136, 'https://leetcode.com/problems/single-number/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Number of 1 Bits', 'number-of-1-bits', 'easy', 191, 'https://leetcode.com/problems/number-of-1-bits/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Counting Bits', 'counting-bits', 'easy', 338, 'https://leetcode.com/problems/counting-bits/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Reverse Bits', 'reverse-bits', 'easy', 190, 'https://leetcode.com/problems/reverse-bits/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Missing Number', 'missing-number', 'easy', 268, 'https://leetcode.com/problems/missing-number/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Sum of Two Integers', 'sum-of-two-integers', 'medium', 371, 'https://leetcode.com/problems/sum-of-two-integers/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;

INSERT INTO problems (title, slug, difficulty, lc_number, lc_url, description, examples, constraints, test_cases, function_signature)
VALUES ('Reverse Integer', 'reverse-integer', 'medium', 7, 'https://leetcode.com/problems/reverse-integer/', '', '[]', '[]', '[]', '{}')
ON CONFLICT (slug) DO NOTHING;
