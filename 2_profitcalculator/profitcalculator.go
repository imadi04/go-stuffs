package main

import "fmt"

func main() {
	var revenue int
	var expenses int
	var tax float64
	var EBT float64 //earning before tax
	var EAT float64 // earning after tax
	fmt.Println("Enter Revenue:")
	fmt.Scan(&revenue)
	fmt.Println("Enter Expenses:")
	fmt.Scan(&expenses)
	fmt.Println("Enter tax rate:")
	fmt.Scan(&tax)
	EBT = float64(revenue) - float64(expenses)
	fmt.Println("Earning before tax=", EBT)
	EAT = float64(revenue) - float64((float64(revenue) * (float64(tax) / 100)))
	fmt.Println("Earning After ta=", EAT-float64(expenses))

}
