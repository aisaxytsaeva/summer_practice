package main

import "fmt"

func LongestString(strings ...string) string {
	if len(strings) == 0 {
		return ""
	}

	longest := strings[0]
	for _, s := range strings[1:] {
		if len(s) > len(longest) {
			longest = s
		}
	}
	return longest
}

func main() {

	fmt.Println("Самая длинная строка:",
		LongestString("яблоко", "апельсин", "банан", "гранат"))

}
