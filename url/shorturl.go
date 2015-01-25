package url

import (
	"github.com/FGM/kurz/storage"
)

type ShortUrl struct {
	Value       string
	ShortFor    LongUrl
	Domain      string
	Strategy    string // name of AliasingStrategy
	SubmittedBy storage.User
	SubmittedOn int64 // UNIX timestamp
	IsEnabled   bool
}
