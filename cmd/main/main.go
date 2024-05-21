package main

import (
	"fmt"
	"math"
	"os"
	"strings"
	"unicode"
)

func main() {
	answer := myAtoi("   -1123u3761867")
	fmt.Println(answer)
	// data := read()
}

func read() string {
	data, err := os.ReadFile("/home/fuad/projects/go-solve/cmd/main/input.txt")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return string(data)
}

func myAtoi(s string) int {
	s = strings.TrimSpace(s)
	negative := false
	if strings.HasPrefix(s, "--") || strings.HasPrefix(s, "++") || strings.HasPrefix(s, "-+") ||
		strings.HasPrefix(s, "+-") {
		return 0
	}
	if strings.HasPrefix(s, "-") {
		negative = true
	}
	s = strings.TrimFunc(s, removeSign)
	s = strings.TrimLeftFunc(s, removeZeros)
	return convert(s, negative)
}

func convert(s string, negative bool) int {
	answer := 0
	originalN := getLenNumbers(s)
	n := originalN
	for i := 0; i < n; i++ {
		if !unicode.IsDigit(rune(s[i])) {
			answer /= int(math.Pow10(n - i))
			break
		}
		curr := int(s[i] - '0')
		pw := int(math.Pow10(n - i - 1))
		if negative {
			pw *= -1
		}
		if !negative && pw < 0 {
			return math.MaxInt32
		}
		if negative && pw > 1 {
			return math.MinInt32
		}
		mul := curr * pw
		if mul > math.MaxInt32 || pw > math.MaxInt32 {
			return math.MaxInt32
		}
		if mul < math.MinInt32 || pw < math.MinInt32 {
			return math.MinInt32
		}
		answer += mul
	}

	if answer > math.MaxInt32 {
		return math.MaxInt32
	}
	if answer < math.MinInt32 {
		return math.MinInt32
	}

	return answer
}

func getLenNumbers(s string) int {
	counter := 0
	for i := 0; i < len(s); i++ {
		if unicode.IsDigit(rune(s[i])) {
			counter++
		} else {
			break
		}
	}
	return counter
}

func removeSign(r rune) bool {
	return r == '-' || r == '+'
}

func removeZeros(r rune) bool {
	return r == '0'
}
