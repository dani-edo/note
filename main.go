package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

func main() {
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

	todo.Display()
	err = todo.Save()
	if err != nil {
		fmt.Println("Saving todo failed:", err)
		return
	}

	fmt.Println("Todo saved successfully")

	userNote.Display()
	err = userNote.Save()
	if err != nil {
		fmt.Println("Saving note failed:", err)
		return
	}

	fmt.Println("Note saved successfully")
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
