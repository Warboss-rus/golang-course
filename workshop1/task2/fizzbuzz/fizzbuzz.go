package fizzbuzz

import (
	"bytes"
	"strconv"
)

func FizzBuzz(count int) string {
	var buffer bytes.Buffer
	for i := 1; i <= count; i++ {
		if i%15 == 0 {
			buffer.WriteString("Fizz Buzz")
		} else if i%3 == 0 {
			buffer.WriteString("Fizz")
		} else if i%5 == 0 {
			buffer.WriteString("Buzz")
		} else {
			buffer.WriteString(strconv.Itoa(i))
		}
		if i != count {
			buffer.WriteString(", ")
		}
	}
	return buffer.String()
}
