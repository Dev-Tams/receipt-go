package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//standard input to read from terminal
var reader = bufio.NewReader(os.Stdin)

func getInput(prompt string, r *bufio.Reader) (string, error) {

	//read, trim and format prompt as they press enter for a new line
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func promptOptions(b bill, ){

	opt, _ :=getInput("Choose option (a - add item, t - add a tip, s - save bill):", reader)
	fmt.Println(opt)
}

func createBill() bill {

	//create reader that reads from standard input

	//read and request input from function get input
	name, _ := getInput("Kindly enter your name:", reader)

	b := newBill(name)
	fmt.Printf("Created %v's bill\n", b.name)

	return b

}
