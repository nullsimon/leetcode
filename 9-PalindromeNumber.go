package main

import "fmt"

func main() {
	ints := []int{
		123,
		434,
		889989,
		0,
	}
	for i := 0; i < len(ints); i++ {
		fmt.Println(ints[i], " is palindrome :", isPalindrome(ints[i]))
	}

}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	var reverse int
	var temp int = x
	for {
		if temp == 0 {
			break
		}
		reverse = reverse*10 + temp%10
		temp = temp / 10
	}
	if reverse == x {
		return true
	}
	return false
}
