package main

import "fmt"

func main() {
	str := "";
	for i := 1; i < 100; i++ {
		if i%15 == 0 {
			str += "Fizz Buzz"
		} else if i%3 == 0 {
			str += "Fizz"
		} else if i%5 == 0 {
			str += "Buzz"
		} else {
			str += fmt.Sprint(i)
		}
		str += ", "
	}
	fmt.Println(str)
}
