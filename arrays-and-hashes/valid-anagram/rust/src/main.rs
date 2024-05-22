use std::collections::HashMap;

fn main() {
    let sol = Solution::is_anagram("anagram".to_string(), "nagaram".to_string());
    println!("{}", sol);
    let sol2 = Solution::is_anagram("rat".to_string(), "car".to_string());
    println!("{}", sol2);
}

pub struct Solution;

impl Solution {
    pub fn is_anagram(s: String, t: String) -> bool {
        if s.len() != t.len() {
            return false;
        }
        let mut chars = HashMap::new();
        for c in s.chars() {
            let char_num = chars.get(&c).unwrap_or(&0);
            chars.insert(c, char_num + 1);
        }
        for c in t.chars() {
            let char_num = chars.get(&c).unwrap_or(&0);
            chars.insert(c, char_num - 1);
        }
        chars.values().all(|&v| v == 0)
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_is_anagram() {
        assert_eq!(
            Solution::is_anagram("anagram".to_string(), "nagaram".to_string()),
            true
        );
        assert_eq!(
            Solution::is_anagram("rat".to_string(), "car".to_string()),
            false
        );
    }
}
