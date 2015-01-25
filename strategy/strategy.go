/*
The "strategy" package in Kurz provides the aliasing strategies.

Files
	- strategy.go contains the interface and base implementation
	- strategies.go contains the strategy instances and utilities
	- manual.go contains the "manual" strategy
	- hexrcr32.go contains the "hexcrc32" strategy
*/
package strategy

import (
	"errors"
	"github.com/FGM/kurz/storage"
	"github.com/FGM/kurz/url"
	"log"
)

/*
AliasingStrategy defines the operations provided by the various aliasing implementations:

The options parameter for Alias() MAY be used by some strategies, in which case they
have to define their expectations about it.
*/
type AliasingStrategy interface {
	Name() string                                                        // Return the name of the strategy object
	Alias(url url.LongUrl, options ...interface{}) (url.ShortUrl, error) // Return the short URL (alias) for a given long (source) URL
	UseCount(storage storage.Storage) int                                // Return the number of short URLs (aliases) using this strategy.
}

type baseStrategy struct{}

func (y baseStrategy) Name() string {
	return "base"
}

func (y baseStrategy) Alias(long url.LongUrl, options ...interface{}) (url.ShortUrl, error) {
	var ret url.ShortUrl
	var err error = errors.New("Base strategy is abstract")

	return ret, err
}

/**
Any nonzero result is likely an error.
*/
func (y baseStrategy) UseCount(s storage.Storage) int {
	sql := `
SELECT COUNT(*)
FROM shorturl
WHERE strategy = ?
	`
	var count int
	err := s.DB.QueryRow(sql, y.Name()).Scan(&count)
	if err != nil {
		count = 0
		log.Printf("Failed querying database for base strategy use count: %v\n", err)
	}

	return count
}
