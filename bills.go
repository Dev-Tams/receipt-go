package main

import (
	"fmt"

)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
	tax  float64
}

func newBill(name string) bill {
	b := bill{
		name: name,
		items: map[string]float64{},
		tip: 0,
	}
	return b
}

//creating a reciever func to format bill

func (b *bill) format() string {

	fb := "Bill Items: \n"
	/**/
	var total float64 = 0.0
	var tax float64 =  0.10

	//cycle round the struct
	for k, v := range b.items {
		fb += fmt.Sprintf("%-20v....$%v\n", k+":", v)
		total += tax + v 
	}

	fb += fmt.Sprintf("%-20v....$%0.2f\n", "tax:", tax)
	fb += fmt.Sprintf("%-20v....$%0.2f\n", "tip:", b.tip)
	fb += fmt.Sprintf("%-20v....$%0.2f", "total:", total+b.tip)

	return fb
}

//add tip
func (b *bill) addTip(tip float64){
	b.tip = tip
}

//add item to the bill

func (b *bill) addItem(name string, price float64){
	b.items[name] = price
}
