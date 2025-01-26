package main

func main() {
	indraReverse := ReverseString("indra")
	println(indraReverse)
}

func ReverseString(s string) string {
	var result string
	for _, v := range s {
		result = string(v) + result
	}
	return result
}
