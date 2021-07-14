# GoLang Problems Challenge

- [Two Sum](https://github.com/alibugrat/go-problems-challenge/blob/main/README.md#two-sum) - Easy
- [Reverse Integer](https://github.com/alibugrat/go-problems-challenge/blob/main/README.md#reverse-integer) - Easy
- [Add Two Number](https://github.com/alibugrat/go-problems-challenge/blob/main/README.md#add-two-number) - Medium
- [Length Of Longest Substring](https://github.com/alibugrat/go-problems-challenge/blob/main/README.md#length-of-longest-substring) - Medium



## [Two Sum](https://github.com/alibugrat/go-problems-challenge/blob/main/challenges/twoSum.go)

Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target. You may assume that each input would have exactly one solution, and you may not use the same element twice. You can return the answer in any order.

```sh
Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Output: Because nums[0] + nums[1] == 9, we return [0, 1].
```

## [Reverse Integer](https://github.com/alibugrat/go-problems-challenge/blob/main/challenges/reverseInt.go)

Given a signed 32-bit integer x, return x with its digits reversed. If reversing x causes the value to go outside the signed 32-bit integer range [-231, 231 - 1], then return 0.

Assume the environment does not allow you to store 64-bit integers (signed or unsigned).

```sh
Input: x = 123
Output: 321

Input: x = -123
Output: -321
```

## [Add Two Number](https://github.com/alibugrat/go-problems-challenge/blob/main/challenges/addTwoNumbers.go)

You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

```sh
Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Explanation: 342 + 465 = 807.
```



## [Length Of Longest Substring](https://github.com/alibugrat/go-problems-challenge/blob/main/challenges/LengthOfLongestSubstring.go)

Given a string s, find the length of the longest substring without repeating characters.

```sh
Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.

Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.
```
