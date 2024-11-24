package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string
}

// constructor function to create note , arg: title and body ; return : Todo instance or error
func New(text string) (Todo, error) {
	if text == "" {
		return Todo{}, errors.New("invalid Input")
	}
	return Todo{
		Text: text,
	}, nil // return nil as error since no error
}

func (t Todo) Display() {
	fmt.Printf("\nYour Todo has following content: \n\n %v", t.Text)
}

// saving todo as json file
func (t Todo) SaveJson() error {
	fileName := "todo.json"
	json, err := json.Marshal(t)
	if err != nil {
		return err
	}

	fmt.Println(fileName)
	return os.WriteFile(fileName, json, 0644)
}
