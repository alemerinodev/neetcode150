package main

import "testing"

func TestIsAnagram(t *testing.T) {
    s := "anagram"
    v := "nagaram"
    if !isAnagram(s, v) {
        t.Errorf("Expected true, got false")
    }
    s = "rat"
    v = "car"
    if isAnagram(s, v) {
        t.Errorf("Expected false, got true")
    }
}
