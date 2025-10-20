package main

import (
	"fmt"
	"strings"
)

func RunTest[T any, S any](
	problemName string,
	testCases []T, 
	test func(testCase T, solution S) (bool, error),
	solution S,
	) {
	passed := 0
	failed := 0

	fmt.Println(fmt.Sprintf("Running Tests for: %s", problemName))
	fmt.Println(strings.Repeat("=", 70))

	for i, testCase := range testCases {
		result, err := test(testCase, solution)

		if err != nil {
			failed++
			fmt.Println(fmt.Sprintf("Test %d failed", i+1))
			fmt.Println(fmt.Sprintf("   Error: %s", err.Error()))
		}

		if result {
			passed++
			fmt.Println(fmt.Sprintf("Test %d passed", i+1))
		}
	}

	fmt.Println(strings.Repeat("=", 70))
	fmt.Println(fmt.Sprintf("Results: %d passed, %d failed, %d total", passed, failed, len(testCases)))

	if failed == 0 {
		fmt.Println("All tests passed!")
	}
}