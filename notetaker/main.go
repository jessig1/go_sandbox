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

// type displayer interface {
// 	Display() error
// }

// interface that embedds other interfaces and method
type outputable interface {
	saver
	Display()
}

//	type outputable interface {
//		Save() error
//		Display()
//	}
//
// take data, write it to json
func main() {
	printSomething(1)
	printSomething(1.6)
	printSomething("test")
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

	err = outputData(todo)

	if err != nil {
		return
	}

	outputData(userNote)

}

// allows function to pass in any value type i.e. string, float, int
func printSomething(value interface{}) {
	intVal, ok := value.(int)
	// checks type of value and outputs if type matches
	if ok {
		fmt.Println("Integer: ", intVal)
		return
	}

	floatVal, ok := value.(float64)
	if ok {
		fmt.Println("float: ", floatVal)
		return
	}

	// switch value.(type) {
	// case int:
	// 	fmt.Println("integer: ", value)
	// case float64:
	// 	fmt.Println("float: ", value)
	// case string:
	// 	fmt.Println("string: ", value)

	// }
	// fmt.Println(value)
}
func outputData(data outputable) error {
	data.Display()
	return saveData(data)
}
func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println(" Cannot save note")
		return err
	}
	fmt.Println(" Save note succeeded")
	return nil
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
