package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("We will input text to a file, enter text to add in file:")
	reader := bufio.NewReader(os.Stdin)
	data, _ := reader.ReadString('\n') // Read input until newline (Enter)
	user_input := []byte(data)
	input_to_file(user_input)
	fmt.Println("Now we will read from that file :)")
	read_from_file()

}
func input_to_file(user_input []byte) {
	err := os.WriteFile("sample.txt", user_input, 0644)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Scan(&user_input) // will only take single word from input
	//data := []byte("Hello Aditya")
	//err := os.WriteFile("sample.txt", data, 0644)
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
func read_from_file() {
	data, _ := os.ReadFile("sample.txt")
	fmt.Println(string(data))
	fmt.Println(data) // this will print bytes , not actual string.
}
