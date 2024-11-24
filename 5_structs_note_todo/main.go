package main

import (
	"bufio"
	"fmt"
	"my_mod/note"
	"my_mod/todo"
	"os"
	"strings"
)

// defining interface called saver ; expects SaveJson function in every structs which signs to this interface.
// helps write more generic , flexible and reusable code.
type saver interface {
	SaveJson() error // SaveJson function with rturn type as error.
}

type outputtable interface {
	saver     // embedding saver interface.
	Display() // addding display func , which is present in both struct todo and note
}

func main() {
	todoText := getUserInput("Todo Text: ")
	todo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}
	//todo.Display()
	//err = saveData(todo)
	err = outputData(todo)
	if err != nil {
		fmt.Println(err)
		return
	}
	noteTitle, noteBody := getNoteData()
	userNote, err := note.New(noteTitle, noteBody)
	if err != nil {
		fmt.Println(err)
		return
	}
	//userNote.Display()
	//err = saveData(userNote)
	err = outputData(userNote)
	if err != nil {
		return
	}

	fmt.Println("\nPress Enter to exit...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	//fmt.Println("Title: ", userNote.Title)
	//fmt.Println("Body: ", userNote.Body)

}

func saveData(data saver) error {
	err := data.SaveJson()
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Saved!!!")
	return nil
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
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
