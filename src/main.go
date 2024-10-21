package main

import (
	"fmt"
	"math/rand/v2"
)

type person struct {
	name   string
	age    int
	height float32
	weight float32 // double 32 bit
}

func main() {
	fmt.Println("Hello world")
	for i := 0; i < 10; i++ {
		//Type Inference  ===> Define variable without type (automatic type define like var in dart)
		x := rand.IntN(100)
		// Manual variable type define
		var message string
		if x%2 == 1 {
			message = fmt.Sprintf("%d is Odd", x)
			fmt.Println(message)
		} else {
			message = fmt.Sprintf("%d is Even", x)
			fmt.Println(message)
		}
	}
	fmt.Println("========================Array============================")
	// Array have to define size
	box := [3]string{"orange", "apple", "grape"}
	fmt.Println(box)
	fmt.Println("==========================================================")

	// =========================================
	// var box [3]string
	// box[0] = "orange"
	// box[1] = "apple"
	// box[2] = "grape"
	// ==========================================

	// Slice don't have to define array size (grawable array) use append to add value to array like (itemList.add(item1) ==== > itemList = append(itemList , item1))
	fmt.Println("========================Slice============================")
	var itemBox []string
	itemBox = append(itemBox, "sword")
	itemBox = append(itemBox, "shield")
	itemBox = append(itemBox, "dagger")
	fmt.Println(itemBox)
	fmt.Println("==========================================================")
	//
	//
	//
	printItemList(itemBox)
	even := isNumberEven()
	fmt.Println("isEven : ", even)
	random, output := isRandomNumberOdd()
	fmt.Println(fmt.Sprintf("%d"+" "+output, random))
	fmt.Println("=========================== Structure ===============================")
	person1 := person{name: "Tossaporn Meesiri", age: 26, height: 160, weight: 62.3}
	fmt.Println("My name is "+person1.name+"age : "+fmt.Sprintf("%d", person1.age), "Height : "+fmt.Sprintf("%.2f", person1.height), "Weihgt : "+fmt.Sprintf("%.1f", person1.weight))

}

// function without return value
func printItemList(itemList []string) {
	fmt.Println("========================For in range============================")
	for index, item := range itemList {
		fmt.Println(fmt.Sprintf("At Index %d", index) + "\t:\t" + item)
	}
}

// Function with return value
func isNumberEven() bool {
	var num int
	fmt.Print("Input Number :")
	// Assign value must pass pointer to value ('&num' instead of using 'num')
	fmt.Scanf("%d", &num)
	if num%2 == 1 {
		return false
	} else {
		return true
	}
}

// Function with multiple return value
func isRandomNumberOdd() (int, string) {
	randomNumber := rand.IntN(100)
	if randomNumber%2 == 0 {
		return randomNumber, "isEven"

	} else {
		return randomNumber, "isOdd"
	}

}
