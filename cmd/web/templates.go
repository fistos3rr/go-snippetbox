package main

import "github.com/fistos3rr/go-snippetbox/internal/models"

type templateData struct {
	Snippet *models.Snippet
	Snippets []*models.Snippet
}
