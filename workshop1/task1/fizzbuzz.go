package main

import "fmt"

func main() {
	str := ""
	const count = 100
	for i := 1; i <= count; i++ {
		if i%15 == 0 {
			str += "Fizz Buzz"
		} else if i%3 == 0 {
			str += "Fizz"
		} else if i%5 == 0 {
			str += "Buzz"
		} else {
			str += fmt.Sprint(i)
		}
		if i != count {
			str += ", "
		}
	}
	fmt.Println(str)
}
