package main

import "fmt"

//function eg1
// func add(x, y int) int {
// 	return x + y
// }

func main() {

	// //variable eg.
	// var message string
	// message = "Hello, Go!"
	// fmt.Println(message)

	// //constant eg
	// const pi = 3.14
	// fmt.Println(pi)

	// //function eg1
	// result := add(3, 4)
	// fmt.Println(result)

	// // if statement
	// x := 10
	// if x > 5 {
	// 	fmt.Println("x is greater than 5")
	// } else {
	// 	fmt.Println("x is not greater than 5")
	// }

	// // for loop
	// for i := 0; i < 5; i++ {
	// 	fmt.Println(i)
	// }

	// // switch statement
	// day := "Monday"
	// switch day {
	// case "Monday":
	// 	fmt.Println("It's Monday!")
	// case "Tuesday":
	// 	fmt.Println("It's Tuesday!")
	// default:
	// 	fmt.Println("It's some other day.")
	// }

	// range
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

}
