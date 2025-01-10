package note

import (
	"errors"
	"time"
)

type Note struct {
	title     string
	content   string
	createdAt time.Time
}

func New(content, title string) (Note, error) {
	if content == "" || title == "" {
		return Note{}, errors.New(("empty input"))
	}

	return Note{
		content:   content,
		title:     title,
		createdAt: time.Now(),
	}, nil
}
