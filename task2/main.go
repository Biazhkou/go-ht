package main

func MapTo(inputData []int, fn func(item, i int) string) []string {
	var result []string
	for index, value := range inputData {
		result = append(result, fn(value, index))
	}
	return result
}

func Convert(arr []int) []string {
	fn := func(item, i int) string {
		numbers := map[int]string{1: "one", 2: "two", 3: "three", 4: "four", 5: "five", 6: "six", 7: "seven", 8: "eight", 9: "nine"}
		value, ok := numbers[item]
		if ok {
			return value
		}
		return "unknown"
	}
	return MapTo(arr, fn)
}

func main() {
}
