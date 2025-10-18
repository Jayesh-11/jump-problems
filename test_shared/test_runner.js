class TestRunner {
  constructor(testCases, test, solution) {
    this.passed = 0;
    this.failed = 0;
    this.testCases = testCases;
    this.test = test;
    this.solution = solution;
  }

  run() {
    console.log("\nRunning Two Sum Tests...\n");
    console.log("=".repeat(70));

    this.testCases.forEach((testCase, idx) => {
      try {
        const result = this.test(testCase, this.solution);
        if (result === true) {
          this.passed++;
          console.log(`Test ${idx + 1} passed`);
        }
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

module.exports = TestRunner;
