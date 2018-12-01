package fizzbuzz

import "testing"

func TestFizzBuzz(t *testing.T) {
	cases := []struct {
		count    int
		expected string
	}{
		{-5, ""},
		{0, ""},
		{1, "1"},
		{3, "1, 2, Fizz"},
		{5, "1, 2, Fizz, 4, Buzz"},
		{15, "1, 2, Fizz, 4, Buzz, Fizz, 7, 8, Fizz, Buzz, 11, Fizz, 13, 14, Fizz Buzz"},
		{36, "1, 2, Fizz, 4, Buzz, Fizz, 7, 8, Fizz, Buzz, 11, Fizz, 13, 14, Fizz Buzz, 16, 17, Fizz, 19, Buzz, Fizz, 22, 23, Fizz, Buzz, 26, Fizz, 28, 29, Fizz Buzz, 31, 32, Fizz, 34, Buzz, Fizz"},
	}
	for _, c := range (cases) {
		result := FizzBuzz(c.count)
		if (result != c.expected) {
			t.Errorf("Unexpected result: expected %q, got %q", c.expected, result)
		}
	}
}
