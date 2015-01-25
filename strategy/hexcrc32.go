package strategy

import (
	"errors"
	"github.com/FGM/kurz/storage"
	"github.com/FGM/kurz/url"
)

/*
HexCrc32Strategy is a legacy AliasingStrategy : hex dump for crc32 hash of source URL.

  - Pros :
    - URLs are easy to communicate over the phone, especially to programmers, even in poor sound conditions.
  - Cons :
    - They are rather long
    - Does not handle collisions: first come, first serve. Later entries are simply rejected.
*/
type HexCrc32Strategy struct {
	base baseStrategy
}

func (s HexCrc32Strategy) Name() string {
	return "hexCrc32"
}

func (s HexCrc32Strategy) Alias(short url.LongUrl, options ...interface{}) (url.ShortUrl, error) {
	var ret url.ShortUrl

	return ret, errors.New("HexCrc32 not implemented yet")
}

func (y HexCrc32Strategy) UseCount(s storage.Storage) int {
	return y.base.UseCount(s)
}
