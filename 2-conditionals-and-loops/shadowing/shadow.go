package main

func main() {
	var int int = 12
	var append string = "minutes"
	var fmt = "formatter"

	var count int // error: int is not a type
	var languages = append([]string{}, "Español") // error: cannot call non-function append (variable of type string)
	fmt.Println(int, append, "on", fmt, languages) // fmt.Println undefined
}