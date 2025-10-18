//  Two Sum - https://leetcode.com/problems/two-sum

const testCases = require("./two_sum.test_cases.json");
const TestRunner = require("./test_runner.js");

/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
const twoSum = function (nums, target) {
  return [];
};

/**
 * Test function for two sum
 * @param {Object} test - Test case
 * @param {number[]} test.case - Test case array
 * @param {number} test.target - Test case target
 * @returns {boolean|void} - True if test passed, false otherwise
 */
const test = (test) => {
  const testCase = test.case;
  const target = test.target;
  const result = twoSum(testCase, target);
  const resultSum = testCase[result[0]] + testCase[result[1]];

  if (result[0] === result[1]) {
    throw new Error(
      `Result is not correct: expected two different indices but got the same index`
    );
  }

  if (resultSum !== target) {
    throw new Error(
      `Result is not correct: expected ${target} but got ${resultSum}`
    );
  }

  if (resultSum === target) {
    return true;
  }

  throw new Error(
    `Result is not correct: expected ${target} but got ${resultSum}`
  );
};

const runner = new TestRunner(testCases, test);

runner.run();
