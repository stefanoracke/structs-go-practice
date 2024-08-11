package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/practice-structs/note"
)

func main() {
	title, content := getNoteData()

	userNote, err := note.New(title, content)
	if err != nil {

		return
	}
	userNote.Display()
	err = userNote.Save()

	if err != nil {
		fmt.Println(("Saving the note failed. "))
		return
	}
	fmt.Println("Saving the note succeeded!")
}

func getNoteData() (string, string) {
	title := getUserInput("Note title: ")

	content := getUserInput("note content: ")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
