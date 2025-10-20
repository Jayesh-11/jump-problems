mod test_runner;
use std::fs;
use std::io;
use serde::{Serialize, Deserialize};

#[derive(Serialize, Deserialize, Debug)]
struct TestCase {
    case: Vec<i32>,
    target: i32,
}

fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
    return vec![];
}

fn test(test_case_instance: &TestCase, solution: fn(nums: Vec<i32>, target: i32) -> Vec<i32>) -> Result<bool, io::Error>{
    let test_case = &test_case_instance.case;
    let target = test_case_instance.target;
    let result = solution(test_case.clone(), target);
    let result_sum = test_case[result[0] as usize] + test_case[result[1] as usize];

    if result[0] == result[1] {
        return Err(io::Error::new(io::ErrorKind::InvalidInput, "Result is not correct: expected two different indices but got the same index"));
    }

    if result_sum != target {
        return Err(io::Error::new(io::ErrorKind::InvalidInput, "Result is not correct: expected $target but got $resultSum"));
    }

    if result_sum == target {
        return Ok(true);
    }

    return Err(io::Error::new(io::ErrorKind::InvalidInput, "Result is not correct: expected $target but got $resultSum"));
}

fn main() {
    println!("Hello, world!");
    let file = fs::File::open("two_sum.test_cases.json").expect("file should open read only");
    let test_cases: Vec<TestCase> = serde_json::from_reader(file).expect("file should be proper JSON");

    test_runner::run_test("Two Sum", test_cases, test, two_sum)
    
}
