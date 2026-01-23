package main

import (
	"fmt"
	"github.com/Vivek-Patel3/note-taker/internal/models"
)

func main() {
	title, content := getNoteDetails()

	userNote, err := models.New(title,content)

	if err != nil {
		fmt.Println(err)
		return
	}

}

func getNoteDetails() (string, string) {
	title := getInputFromUser("Title:")		
	content := getInputFromUser("Content:")

	return title,content
}

func getInputFromUser(prompt string) string {
	fmt.Print(prompt)
	var value string
	fmt.Scan(&value)

	return value
}
