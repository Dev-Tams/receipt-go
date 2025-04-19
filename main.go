package main

import "fmt"

func main() {

	// mybill := newBill("mario's bill")
	
	mybill := createBill()
	// mybill.addItem("beans", 7)
	// mybill.addTip(2)
	promptOptions(mybill)

	fmt.Println(mybill)
}