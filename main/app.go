package main

import (
	"fmt"
	"errors"
)

func main() {
	title, content, err := getNoteDetails()

	if(err != nil) {
		fmt.Println(err.Error());
	}
}

func getNoteDetails() (string, string, error) {
	title, err := getInputFromUser("Title:")		

	if(err != nil) {
		fmt.Println(err.Error())
		return "","",err
	}

	content, err := getInputFromUser("Content:")

	if(err != nil) {
		fmt.Println(err.Error())
		return  title,"",err
	}

	return title,content,nil
}

func getInputFromUser(prompt string) (string, error) {
	fmt.Print(prompt)
	var value string
	fmt.Scan(&value)

	if(value == "") {
		return "",errors.New("Cannot be empty value")
	}

	return value,nil
}