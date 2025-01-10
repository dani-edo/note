package note

import (
	"errors"
	"fmt"
	"time"
)

type Note struct {
	title     string
	content   string
	createdAt time.Time
}

// func (receiver argument) methodName(argument) (return type, error)
func (note Note) Display() {
	fmt.Printf("Your note titled %v value has the following content: \n\n%v\n", note.title, note.content)
}

func New(content, title string) (Note, error) {
	if content == "" || title == "" {
		return Note{}, errors.New(("empty input"))
	}

	return Note{
		content:   content,
		title:     title,
		createdAt: time.Now(),
	}, nil
}
