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

	var menuId int
menuloop:

	for {
		printMenu()
		fmt.Scanf("%d", &menuId)

		switch menuId {
		case 0:
			break menuloop
		case 1:
			simpleSyntax()
		case 2:
			functionArray()
		case 3:
			functionSlice()
		case 4:
			usingFunc()
		case 5:
			functionStructure()
		case 6:
			fuctionPointer()

		default:
			fmt.Println("Your Select menu is not correct")
		}
		menuId = 0

	}

	//
	//
	//

}
func printMenu() {
	fmt.Println("Select Menu")
	fmt.Println("0 : exit program")
	fmt.Println("1 : Simple Go Syntax and loop")
	fmt.Println("2 : Using Array")
	fmt.Println("3 : Using Slice")
	fmt.Println("4 : Using Function")
	fmt.Println("5 : Using Structure")
	fmt.Println("6 : Experiment 0")

	fmt.Printf(" Input Number to select menu : ")
}

func simpleSyntax() {
	fmt.Println("========= Simple Go Syntax and loop ============")
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
	fmt.Println("==========================================================")
}

func fuctionPointer() {
	fmt.Println("=========================== Pointer ===============================")
	name := "Tossaporn"
	pointer := &name
	name2 := name
	pointer2 := &name2
	fmt.Println(name)
	fmt.Println(pointer)
	fmt.Println(name2)
	fmt.Println(pointer2)

	fmt.Println("==========================================================")
}

func functionArray() {
	fmt.Println("========================Array============================")
	// Array have to define size
	box := [3]string{"orange", "apple", "grape"}
	fmt.Println(box)
	fmt.Println("==========================================================")
}

func functionSlice() {
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
	printItemList(itemBox)

	fmt.Println("==========================================================")
}

func usingFunc() {
	fmt.Println("======================== Function ============================")
	even := isNumberEven()
	fmt.Println("isEven : ", even)
	random, output := isRandomNumberOdd()
	fmt.Println(fmt.Sprintf("random number :: %d"+" "+output, random))
	fmt.Println("==========================================================")
}

func functionStructure() {
	fmt.Println("=========================== Structure ===============================")
	person1 := person{name: "Tossaporn Meesiri", age: 26, height: 160, weight: 62.3}
	fmt.Println("My name is "+person1.name+"age : "+fmt.Sprintf("%d", person1.age), "Height : "+fmt.Sprintf("%.2f", person1.height), "Weihgt : "+fmt.Sprintf("%.1f", person1.weight))
	fmt.Println("==========================================================")
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
