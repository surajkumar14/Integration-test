#!/bin/bash

# Define log files and target directory
TARGET_DIR="."  # Update this to your test directory
RESULT_LOG="./report/result.log"
COVERAGE_FILE="./report/coverage.out"
COVERAGE_HTML="./report/coverage.html"
BENCHMARK_LOG="./report/benchmark.log"



# Clean previous log files
rm -f $RESULT_LOG $COVERAGE_FILE $COVERAGE_HTML $MARKDOWN_REPORT

# Navigate to the target directory
cd $TARGET_DIR || { echo "Directory not found: $TARGET_DIR"; exit 1; }

# Run tests with coverage and verbose output, and log results
echo "Running integration tests in $TARGET_DIR..."
go test -v -coverprofile=$COVERAGE_FILE ./... > $RESULT_LOG 2>&1

# Check if the tests ran successfully
if grep -q "FAIL" $RESULT_LOG; then
    echo "Some tests failed. Check the log for details."
else
    echo "Tests completed successfully."
fi

# Generate HTML coverage report
echo "Generating coverage report..."
go tool cover -html=$COVERAGE_FILE -o $COVERAGE_HTML >> $RESULT_LOG 2>&1

# Check if coverage report was generated successfully
if [ $? -eq 0 ]; then
    echo "Coverage report generated: $COVERAGE_HTML"
else
    echo "Failed to generate coverage report."
fi

# Run benchmarks and log results
echo "Running benchmarks in $TARGET_DIR..."
go test -bench=. -benchmem ./... > $BENCHMARK_LOG 2>&1

echo "Log file: $RESULT_LOG"
