package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
)

// take data, write it to json
func main() {
	title, content := getNoteData()

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}
	userNote.Display()
	err = userNote.Save()

	if err != nil {
		fmt.Println("Cannot save note")
		return
	}
	fmt.Println("Save note succeeded")
}

func getNoteData() (string, string) {
	title := getUserInput("Note tile:")

	content := getUserInput("Note content:")
	return title, content
}
func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	//takes multiple word input vs. single w/scan
	reader := bufio.NewReader(os.Stdin)
	// defines where to stop reading
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
