package model

import (
	"html/template"
)

type TranslatedProjectWithThumbnail struct {
	ID        int64
	Thumbnail *Asset
	Name      template.HTML
	Body      template.HTML
}
