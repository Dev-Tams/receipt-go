package main

import "fmt"

func main() {

	mybill := newBill("mario's bill")
	mybill.addTip(2)
	mybill.addItem("beans", 7)

	fmt.Println(mybill.format())
}