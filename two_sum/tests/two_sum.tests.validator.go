package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type TestCase struct {
	Case []int `json:"case"`
	Target int `json:"target"`
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

for i:=range testCases{
	testCase := testCases[i].Case
	testTarget := testCases[i].Target
	testCaseLength := len(testCase)
	if testCaseLength < 2 {
		fmt.Println("Invalid Test case", i+1, "failed: Array must contain at least 2 elements")
		continue
	}
	if testCaseLength > 10000 {
		fmt.Println("Invalid Test case", i+1, "failed: Array must contain at most 10000 elements")
		continue
	}
	if testTarget < -1000000000 || testTarget > 1000000000 {
		fmt.Println("Invalid Test case", i+1, "failed: Target must be between -1000000000 and 1000000000")
		continue
	}
	isValidTestCaseElements := true
	invalidElement := 0
	for j:=range testCase{

		if testCase[j] < -1000000000 || testCase[j] > 1000000000 {
			isValidTestCaseElements = false
			invalidElement = testCase[j]
			break
		}
		
	}
	if !isValidTestCaseElements {
		fmt.Println("Invalid Test case", i+1, "failed: Array must contain numbers between -1000000000 and 1000000000 but contains", invalidElement)
		continue
	}
	umap := make(map[int]int)
	numberOfSolutions := 0
	for j:=range testCase{
		if _, ok := umap[testTarget - testCase[j]];ok {
			numberOfSolutions++
		}
		umap[testCase[j]] = j
	}

	if numberOfSolutions != 1 {
		fmt.Println("Invalid Test case", i+1, "failed: Array must contain exactly one solution")
		continue
	}


	fmt.Println("Valid Test case ", i+1, " passed")
}

}