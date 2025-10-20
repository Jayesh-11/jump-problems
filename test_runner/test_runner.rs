use std::io;

pub fn run_test<T,S: Copy>(problem_name: &str, test_cases: Vec<T>, test: fn(&T, S) -> Result<bool, io::Error>, solution: S) {
    let mut passed = 0;
    let mut failed = 0;

    println!("Running tests for: {}", problem_name);
    println!("{}", "=".repeat(70));

    for (i, test_case) in test_cases.iter().enumerate() {
        let result = test(test_case, solution);
        match result {
            Ok(true) => {
                passed += 1;
                println!("Test {} passed", i);
            }
            Ok(false) => {
                failed += 1;
                println!("Test {} failed", i);
            }   
            Err(err) => {
                failed += 1;
                println!("Test {} failed", i);
                println!("   Error: {}", err);
            }
        }
    }

    println!("{}", "=".repeat(70));
    println!("Results: {} passed, {} failed, {} total", passed, failed, test_cases.len());

    if failed == 0 {
        println!("All tests passed!");
    }

}