package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/Vivek-Patel3/note-taker/internal/models"
)

func main() {
	title, content := getNoteDetails()

	userNote, err := models.New(title,content)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.Display()
	userNote.Save()

	fmt.Println("File saved successfully")
}

func getNoteDetails() (string, string) {
	title := getInputFromUser("Title:")		
	content := getInputFromUser("Content:")

	return title,content
}

func getInputFromUser(prompt string) string {
	fmt.Printf("%v ",prompt)
	
	reader := bufio.NewReader(os.Stdin)

	// now extract the string from reader
	text, err := reader.ReadString('\n')

	if(err != nil) {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
