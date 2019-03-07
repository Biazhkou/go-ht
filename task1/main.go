package main

// Specify Filter function here

func main() {

}

func Filter(inputData []int, fn func(int, int) bool) []int {
	var result []int
	for index, value := range inputData {
		if fn(value, index) {
			result = append(result, value)
		}
	}
	return result
}
