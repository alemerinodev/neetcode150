package main

import (
	"fmt"
)

func main() {
    s := "anagram"
    t := "nagaram"
    fmt.Println(isAnagram(s, t)) // true

    s = "rat"
    t = "car"
    fmt.Println(isAnagram(s, t)) // false
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var chars = make(map[rune]int)
	for _, c := range s {
		chars[c] = chars[c] + 1
	}
	for _, c := range t {
		chars[c] = chars[c] - 1
	}

    for _, value := range chars {
        if value != 0 {
            return false
        }
    }
	return true 
}
