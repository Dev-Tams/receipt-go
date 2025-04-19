package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// standard input to read from terminal
var reader = bufio.NewReader(os.Stdin)

func getInput(prompt string, r *bufio.Reader) (string, error) {

	//read, trim and format prompt as they press enter for a new line
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func promptOptions(b bill) {

	opt, _ := getInput("Choose option (a - add item, t - add a tip, s - save bill):", reader)
	fmt.Println(opt)

	switch opt {

	//add item
	case "a":
		name, _ := getInput("Enter the name of item:", reader)
		price, _ := getInput("Enter price:", reader)

		//check price input and parse price string to float
		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("The price ought to be a number")
			promptOptions(b)
		}

		//print out
		fmt.Println(name, price)
		b.addItem(name, p)
		fmt.Println("items added", name, price)

	case "t":
		tip, _ := getInput("Kindly enter tip(optional):", reader)
		
		
		//check tip input and parse price string to float
		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("The tip ought to be a number")
			promptOptions(b)
		}
		fmt.Println(tip)
		b.addTip(t)


	case "s":
		fmt.Println("Saved bill", b)


	default:
		fmt.Println("try again werey")
		promptOptions(b)
	}
}

func createBill() bill {

	//create reader that reads from standard input

	//read and request input from function get input
	name, _ := getInput("Kindly enter your name:", reader)

	b := newBill(name)
	fmt.Printf("Created %v's bill\n", b.name)

	return b

}
