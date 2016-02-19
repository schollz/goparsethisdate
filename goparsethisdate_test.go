package main

import (
	"fmt"
	"testing"
)

// func Example1() {
// 	date, _ := ParseDate("02/17/2016 3:05pm")
// 	fmt.Println(date)
// 	// Output: 2016-02-17 15:05:00 +0000 UTC
// }

// func Example2() {
// 	date, _ := ParseDate("02/17/2016 3pm")
// 	fmt.Println(date)
// 	// Output: 2016-02-17 15:00:00 +0000 UTC
// }

// func Example3() {
// 	date, _ := ParseDate("3pm 02/17/2016")
// 	fmt.Println(date)
// 	// Output: 2016-02-17 15:00:00 +0000 UTC
// }

// func Example4() {
// 	date, _ := ParseDate("3pm, 02/17/2016")
// 	fmt.Println(date)
// 	// Output: 2016-02-17 15:00:00 +0000 UTC
// }

// func Example5() {
// 	date, _ := ParseDate("February 17th, 2016")
// 	fmt.Println(date)
// 	// Output: 2016-02-17 00:00:00 +0000 UTC
// }

// func Example6() {
// 	date, _ := ParseDate("Feb 1st 2016")
// 	fmt.Println(date)
// 	// Output: 2016-02-01 00:00:00 +0000 UTC
// }

func Example1() {
	date, _ := ParseDate("Feb. 2nd 2016")
	fmt.Println(date)
	// Output: 2016-02-02 00:00:00 +0000 UTC
}

func BenchmarkParseDate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParseDate("02/17/2016 3:05pm")
	}
}

func Example2() {
	date, _ := ParseDateChannel("Feb. 2nd 2016")
	fmt.Println(date)
	// Output: 2016-02-02 00:00:00 +0000 UTC
}

func BenchmarkParseDateChannel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParseDateChannel("02/17/2016 3:05pm")
	}
}
