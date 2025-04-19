package main

import (
	"fmt"

)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func newBill(name string) bill {
	b := bill{
		name: name,
		items: map[string]float64{
			"rice": 3.99,
			"yam":  2.66,
		},
		tip: 0,
	}
	return b
}

//creating a reciever func to format bill

func (b bill) format() string {

	fb := "Bill Items: \n"
	var total float64 = 0.0
	var tax float64 = total + 0.15

	//cycle round the struct
	for k, v := range b.items {
		fb += fmt.Sprintf("%-20v....$%v\n", k+":", v)
		total += tax + v
	}

	fb += fmt.Sprintf("%-20v....$%0.2f\n", "tax:", tax)
	fb += fmt.Sprintf("%-20v....$%0.2f", "total:", total)

	return fb
}
