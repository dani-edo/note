package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"` // capitalize first letter to be able to accessed by Marshal
}

// func (receiver argument) methodName(argument) (return type, error)
func (todo Todo) Display() {
	fmt.Printf("%v\n", todo.Text)
}

func (todo Todo) Save() error {
	fileName := "todo.json"

	jsonText, err := json.Marshal(todo) // convert struct to json

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, jsonText, 0644) // 0644: editable by any owner only
}

func New(content string) (Todo, error) {
	if content == "" {
		return Todo{}, errors.New(("empty input"))
	}

	return Todo{
		Text: content,
	}, nil
}
