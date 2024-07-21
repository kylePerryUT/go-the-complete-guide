package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
)

func main() {
	title, content := getNoteData()

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println("Error creating note:", err)
		return
	}

	userNote.Display()

	err = userNote.Save()

	if err != nil {
		fmt.Println("Error saving note:", err)
		return
	}

	fmt.Println("Note saved successfully!")
}

func getNoteData() (string, string) {
	noteTitle := getUserInput("Please enter the note title: ")
	noteContent := getUserInput("Please enter the note content: ")

	return noteTitle, noteContent
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)
	var input string

	text, err := bufio.NewReader(os.Stdin).ReadString('\n')

	if err != nil {
		fmt.Println("Error reading input:", err)
		return ""
	}

	input = strings.TrimSuffix(text, "\n")
	input = strings.TrimSuffix(input, "\r")

	return input
}
