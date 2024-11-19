package main

import (
	"bufio"
	"fmt"
	"my_mod/note"
	"os"
	"strings"
)

func main() {
	noteTitle, noteBody := getNoteData()
	userNote, err := note.New(noteTitle, noteBody)
	if err != nil {
		fmt.Println(err)
		return
	}
	userNote.Display()
	userNote.Save()
	userNote.SaveJson()
	//fmt.Println("Title: ", userNote.Title)
	//fmt.Println("Body: ", userNote.Body)

}

func getNoteData() (string, string) {
	fmt.Println("Welcome to Notepad app!")
	noteTitle := getUserInput("Enter note title: ")

	noteBody := getUserInput("Enter Note: ")

	return noteTitle, noteBody
}

func getUserInput(prompt string) string {
	//taking space separated input.
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(prompt)

	input, _ := reader.ReadString('\n')

	value := strings.Fields(input)
	temp := strings.Join(value, " ")

	return temp
}
