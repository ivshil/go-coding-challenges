package main

// ref https://www.hackerrank.com/challenges/sock-merchant
func sockMerchant(n int32, ar []int32) {
	count := int32(0)
	sockMap := make(map[int32]bool)
	for i := int32(0); i < n; i++ {
		if val, ok := sockMap[ar[i]]; ok {
			if val == false {
				sockMap[ar[i]] = true
			} else {
				count++
				sockMap[ar[i]] = false
			}
		} else {
			sockMap[ar[i]] = true
		}
	}
}
