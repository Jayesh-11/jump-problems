package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func twoSum(nums []int, target int) []int {
    return []int{}
}

type TestCase struct {
    Case []int `json:"case"`
    Target int `json:"target"`
}
func test(testCaseInstance TestCase, solution func(nums []int, target int) []int) (bool, error) {
    testCase := testCaseInstance.Case;
    target := testCaseInstance.Target;
    result := solution(testCase, target);
    resultSum := testCase[result[0]] + testCase[result[1]];

    if result[0] == result[1] {
        return false, fmt.Errorf("Result is not correct: expected two different indices but got the same index")
    }

    if resultSum != target {
        return false, fmt.Errorf("Result is not correct: expected %d but got %d", target, resultSum)
    }

    if resultSum == target {
        return true, nil
    }

    return false, fmt.Errorf("Result is not correct: expected %d but got %d", target, resultSum)
}

func main() {
    jsonFile, err := os.Open("two_sum.test_cases.json")
    if err != nil {
       fmt.Println("Error opening file:", err)
       panic(err)
    }
    defer jsonFile.Close()
    var testCases []TestCase
    byteValue, _ := io.ReadAll(jsonFile)
    json.Unmarshal(byteValue, &testCases)
    RunTest("Two Sum", testCases, test, twoSum)
}