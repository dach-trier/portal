package model

import (
	"html/template"
)

type LocalizedProjectWithThumbnail struct {
	ID          string
	Thumbnail   *Image
	Name        template.HTML
	Description template.HTML
}
