package model

type TranslatedInitiativeWithThumbnail struct {
	ID           string
	Kind         string
	Thumbnail    *Image
	Name         string
	Description  string
}
