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
	Title     string    `json:"title"` // capitalize first letter to be able to accessed by Marshal
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

// func (receiver argument) methodName(argument) (return type, error)
func (note Note) Display() {
	fmt.Printf("Your note titled %v value has the following content: \n\n%v\n\n", note.Title, note.Content)
}

func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_") // replace spaces with _
	fileName = strings.ToLower(fileName) + ".json"       // convert to lowercase

	jsonText, err := json.Marshal(note) // convert struct to json

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, jsonText, 0644) // 0644: editable by any owner only
}

func New(content, title string) (Note, error) {
	if content == "" || title == "" {
		return Note{}, errors.New(("empty input"))
	}

	return Note{
		Content:   content,
		Title:     title,
		CreatedAt: time.Now(),
	}, nil
}
