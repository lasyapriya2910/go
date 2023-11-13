package main

import "fmt"

func removeDuplicates(s []string) []string {
	bucket := make(map[string]bool)
	var result []string
	for _, str := range s {
		if _, ok := bucket[str]; !ok {
			bucket[str] = true
			result = append(result, str)
		}
	}
	return result
}
func main() {
	array := []string{"hai", "my", "csd", "batch", "section", "c", "2020", "and", "hai", "csm", "batch", "section"}
	fmt.Println("The given array of string is:", array)
	fmt.Println()
	result := removeDuplicates(array)
	fmt.Println("The array obtained after removing the duplicate entries is:", result)
}
