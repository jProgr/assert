package assert

import (
    "testing"
)

// Equals asserts whether expected and actual are equal.
func Equals[U comparable](test *testing.T, expected, actual U) {
    test.Helper()

    if expected != actual {
        test.Fatalf("Expected: %v. Actual: %v.", expected, actual)
    }
}

// Nil asserts whether value is nil.
func Nil(test *testing.T, value any) {
    test.Helper()

    if value != nil {
        test.Fatalf("Expected: nil. Actual: %v", value)
    }
}
