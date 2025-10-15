/**
 * Two Sum - https://leetcode.com/problems/two-sum
 */

const testCases = require("./tests/two_sum.test_cases.json");

/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
const twoSum = function (nums, target) {
  return [];
};

class TestRunner {
  constructor() {
    this.passed = 0;
    this.failed = 0;
    this.testCases = testCases;
  }

  test(testCase, target) {
    const result = twoSum(testCase, target);
    const resultSum = result[0] + result[1];

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
  }

  run() {
    console.log("\nRunning Two Sum Tests...\n");
    console.log("=".repeat(70));

    this.testCases.forEach((test, idx) => {
      try {
        this.test(test.case, test.target);
        this.passed++;
        console.log(`Test ${idx + 1} passed`);
      } catch (error) {
        this.failed++;
        console.log(`Test ${idx + 1} failed`);
        console.log(`   Error: ${error.message}\n`);
      }
    });

    console.log("=".repeat(70));
    console.log(
      `\nResults: ${this.passed} passed, ${this.failed} failed, ${this.testCases.length} total\n`
    );

    if (this.failed === 0) {
      console.log("All tests passed!\n");
    }
  }
}

const runner = new TestRunner();

runner.run();
