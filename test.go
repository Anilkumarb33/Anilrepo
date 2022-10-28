package main

import "fmt"

func main() {
	even := []string{"0", "2", "4", "6", "8", "10", "12", "14", "16", "18", "20", "21", "22"}
	odd := []string{"1", "3", "5", "7", "9", "11", "13", "15", "17", "19"}

	fmt.Printf("%+v\n", weave(even, odd))
}

// TODO Finish this function to return the combined result of the two arrays (odd, even), where the values alternate.
// TODO Feel free to change and extend the  function as much as you want.
// TODO The result is supposed to look like this: [0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 22]
func weave(a []string, b []string) []string {
	arr3 := []string{}
	var i int
	for i < len(a) || i < len(b) {

		if i < len(a) {
			arr3 = append(arr3, a[i])
			//i = i + 1
		}
		if i < len(b) {
			arr3 = append(arr3, b[i])
			//i = i + 1
		}
		i++
	}
	return arr3
}
