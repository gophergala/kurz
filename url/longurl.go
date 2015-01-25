package url

import "github.com/FGM/kurz/storage"

type LongUrl struct {
	Value string
}

func (l LongUrl) Domain() string {
	return "default"
}

type LongMeta struct {
	Url       LongUrl
	MimeType  string
	Language  storage.Language
	ImagePath string
	Origin    storage.EventInfo
}
