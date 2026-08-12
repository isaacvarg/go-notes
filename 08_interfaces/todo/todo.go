package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Todo struct {
	Content string `json:"content"`
}

func (todo Todo) Display() {
	fmt.Printf("Your todo: %v\n", todo.Content)
}

func (todo Todo) Save() error {
	fileName := "todo"
	fileName = strings.ToLower(fileName) + ".json"

	json, err := json.Marshal(todo)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0o644)
}

func New(title, content string) (Todo, error) {
	if content == "" {
		return Todo{}, errors.New("invalid input")
	}

	return Todo{
		Content: content,
	}, nil
}
