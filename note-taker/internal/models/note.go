package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Note struct {
	Title string `json:"title"` 
	Content string `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}
// tags added to make the file's fields json-like

func (note Note) Save() error {
	fileName := strings.ReplaceAll(strings.ToLower(note.Title), " ", "_") + ".json"
	
	dataDir := os.Getenv("NOTESAPP_DATA_DIR")
	if dataDir == "" {
		dataDir = "data" // relative to the current working directory
	}

	err := os.MkdirAll(dataDir, 0755)

	if err != nil {
		return err
	}

	// final path to the file
	path := filepath.Join(dataDir, fileName)
	content, err := json.MarshalIndent(note, "", "  ")

	if err != nil {
		return err
	}

	return os.WriteFile(path, content, 0644)
}

func (note Note) Display() {
	fmt.Printf("Title %s\nCreated: %s\n\n%s\n", note.Title, note.CreatedAt.Format(time.RFC822), note.Content)
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