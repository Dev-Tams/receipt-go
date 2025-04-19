package main

import "fmt"

func main() {

	// mybill := newBill("mario's bill")
	// mybill.addTip(2)
	// mybill.addItem("beans", 7)

	mybill := createBill()
	promptOptions(mybill)

	fmt.Println(mybill.format())
}