package model

import (
	"html/template"
)

type TranslatedInitiativeWithThumbnail struct {
	ID          string
	Thumbnail   *Image
	Name        template.HTML
	Description template.HTML
}
