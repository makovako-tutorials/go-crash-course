package main

import ("fmt"
		"math"
		"github.com/makovako/go_crash_course/03_packages/strutil")

func main() {
	fmt.Println(math.Floor(2.8))
	fmt.Println(math.Ceil(2.3))
	fmt.Println(strutil.Reverse("elloH"))
}