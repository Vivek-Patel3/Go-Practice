package models

import(
	"time"
	"errors"
)

type Note struct {
	title string
	content string
	createdAt time.Time
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("Invalid input.")
	}

	return Note{
		title,
		content,
		time.Now(),
	}, nil
}