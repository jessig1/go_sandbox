package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	todo "example.com/note/interfaces"
	"example.com/note/note"
)

type saver interface {
	Save() error
}

// take data, write it to json
func main() {
	title, content := getNoteData()
	todoText := getUserInput("Todo text: ")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	todo.Display()
	err = todo.Save()

	if err != nil {
		fmt.Println(" Cannot save todo")
		return
	}
	fmt.Println(" Save todo succeeded")

	userNote.Display()
	err = userNote.Save()

	if err != nil {
		fmt.Println(" Cannot save note")
		return
	}
	fmt.Println(" Save note succeeded")
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
