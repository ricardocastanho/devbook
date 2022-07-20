package models

import (
	"errors"
	"strings"
)

type Post struct {
	ID        string `json:"id,omitempty"`
	Title     string `json:"title,omitempty"`
	Content   string `json:"content,omitempty"`
	Likes     uint64 `json:"likes"`
	Author    User   `json:"author,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	DeletedAt string `json:"deleted_at,omitempty"`
}

func (p *Post) Validate() error {
	err := p.checkEmptyFields()

	if err != nil {
		return err
	}

	p.trimFields()
	return nil
}

func (p *Post) checkEmptyFields() error {
	if p.Title == "" {
		return errors.New("title is required")
	}
	if p.Content == "" {
		return errors.New("content is required")
	}
	return nil
}

func (p *Post) trimFields() {
	p.Title = strings.TrimSpace(p.Title)
	p.Content = strings.TrimSpace(p.Content)
}
