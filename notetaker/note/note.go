package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

// “ define a struct tag, which allows metadata for a value
type Note struct {
	Title     string    `json:"tile"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func (note Note) Display() {
	fmt.Printf("Your note tile %v \nhas the following content:\n\n%v", note.Title, note.Content)
}

// writes output to json
func (note Note) Save() error {
	//replaces the spaces in the title with underscores
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	// creates json file
	fileName = strings.ToLower(fileName) + ".json"
	json, err := json.Marshal(note)

	if err != nil {
		return err
	}
	return os.WriteFile(fileName, json, 0644)
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("invalid input")
	}
	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}
