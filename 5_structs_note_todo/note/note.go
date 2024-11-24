package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string
	Body      string    // kept them in small letter as we wanted abstraction, later changed it to Caps as json.Marsshal function takes only publicly available data.
	CreatedAt time.Time // to capture timestamp
}

// constructor function to create note , arg: title and body ; return : Note instance or error
func New(title, body string) (Note, error) {
	if title == "" || body == "" {
		return Note{}, errors.New("invalid Input")
	}
	return Note{
		Title:     title,
		Body:      body,
		CreatedAt: time.Now(),
	}, nil // return nil as error since no error
}

func (n Note) Display() {
	fmt.Printf("\nYour note title %v has following content: \n\n %v", n.Title, n.Body)
}

// saving note as txt file
func (n Note) Save() {
	fileName := strings.ReplaceAll(n.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".txt"
	os.WriteFile(fileName, []byte(n.Body), 0644)
	fmt.Printf("\nCongrats! your note is saved as %v and ", fileName)
}

// saving note as json file
func (n Note) SaveJson() error {
	fileName := strings.ReplaceAll(n.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"
	json, err := json.Marshal(n)
	if err != nil {
		return err
	}

	fmt.Println(fileName)
	return os.WriteFile(fileName, json, 0644)
}
