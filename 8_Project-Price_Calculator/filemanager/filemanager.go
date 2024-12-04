package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

//constructor

func New(inputPath, outputPath string) FileManager {
	return FileManager{
		InputFilePath:  inputPath,
		OutputFilePath: outputPath,
	}
}

func (fm FileManager) ReadLines() ([]string, error) { // arg for accepting file path as arg
	file, err := os.Open(fm.InputFilePath)
	if err != nil {
		fmt.Println("File not found!")
		fmt.Println(err)
	}
	// NewScanner() is used to create a Scanner, which is a convenient tool for reading input from a variety of sources
	//(like files, standard input, or strings) token by token.
	scanner := bufio.NewScanner(file)
	var lines []string // to store scanned prices.
	//scanner.Scan() // read one line at a time from file.
	for scanner.Scan() {
		lines = append(lines, scanner.Text()) // text() to get the content of scanned line by scan()
	}
	err = scanner.Err()
	if err != nil {
		fmt.Println("Reading the file content failed!!")
		fmt.Println(err)
		file.Close()
		return nil, err
	}
	file.Close()
	return lines, nil
}

func (fm FileManager) WriteResult(data any) error { //or data interface{}
	file, err := os.Create(fm.OutputFilePath)
	if err != nil {
		return errors.New("Failed to create Json file")
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		file.Close()
		return errors.New("Failed to convert data to Json")
	}
	return nil
}
