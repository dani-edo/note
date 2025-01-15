package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

type saver interface { // type: syntax for declaring a new struct or interface
	Save() error
}

type outputable interface {
	saver
	Display()
}

func main() {
	printSomething(1)
	printSomething(1.5)
	printSomething("test")

	title, content := getNoteData()
	todoText := getUserInput("Todo text: ")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(content, title)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)
	if err != nil {
		fmt.Println(err)
		return
	}

	outputData(userNote)
}

func printSomething(value interface{}) {
	switch value.(type) {
	case int:
		fmt.Println("integer: ", value)
	case float64:
		fmt.Println("float64: ", value)
	case string:
		fmt.Println("string: ", value)
	}
}

func outputData(data outputable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Saving data failed:", err)
		return err
	}

	fmt.Println("Data saved successfully")
	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note title: ")
	content := getUserInput("Note content: ")
	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)  // os.Stdin: OS standard input
	text, err := reader.ReadString('\n') // \n: stop reading when enter is pressed

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n") // remove \n from text (in the end of the string)
	text = strings.TrimSuffix(text, "\r") // in case of Windows, the line break is \r\n

	return text
}
