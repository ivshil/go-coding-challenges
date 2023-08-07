package main

// ref https://www.hackerrank.com/challenges/counting-valleys
func countingValleys(steps int32, path string) int32 {
	current := int32(0)
	count := int32(0)
	runes := []rune(path)
	for _, char := range runes {
		if char == 'U' {
			current++
			if current == 0 {
				count++
			}
		} else {
			current--
		}
	}
	return count
}
