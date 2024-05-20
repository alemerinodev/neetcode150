package main

import "testing"

func TestContainsDuplicate(t *testing.T) {
    input := []int{1, 2, 3, 1}
    output := containsDuplicate(input)
    if output != true {
        t.Errorf("Test failed, expected: %v, got: %v", true, output)
    }
}
