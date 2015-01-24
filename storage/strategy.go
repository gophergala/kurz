package storage

import (
	"database/sql"
	"errors"
	"log"
)

type Strategy interface {
	Name() string
	Alias(url LongUrl) (ShortUrl, error)
	UseCount(storage Storage) int
}

type baseStrategy struct{}

func (y baseStrategy) Name() string {
	return "base"
}

func (y baseStrategy) Alias(long LongUrl) (ShortUrl, error) {
	var ret ShortUrl
	var err error = errors.New("Base strategy is abstract")

	return ret, err
}

/**
Any nonzero result is likely an error.
*/
func (y baseStrategy) UseCount(s Storage) int {
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

/*
Legacy strategy : hex dump for crc32 hash of source URL.

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

func (s HexCrc32Strategy) Alias(url LongUrl) (ShortUrl, error) {
	var ret ShortUrl

	return ret, errors.New("HexCrc32 not implemented yet")
}

func (y HexCrc32Strategy) UseCount(s Storage) int {
	return y.base.UseCount(s)
}

type ManualStrategy struct {
	base baseStrategy
}

func (y ManualStrategy) Name() string {
	return "manual"
}

func (y ManualStrategy) Alias(long LongUrl) (ShortUrl, error) {
	var ret ShortUrl

	return ret, errors.New("Manual not implemented yet")
}

func (y ManualStrategy) UseCount(s Storage) int {
	return y.base.UseCount(s)
}

var Strategies map[string]Strategy

func init() {
	var strategyImplementations []Strategy = []Strategy{
		baseStrategy{},
		HexCrc32Strategy{},
		ManualStrategy{},
	}

	Strategies = make(map[string]Strategy, len(strategyImplementations))
	for _, s := range strategyImplementations {
		Strategies[s.Name()] = s
	}
}

func StrategyStatistics(s Storage) map[string]int64 {
	var ret map[string]int64 = make(map[string]int64, len(Strategies))
	var err error
	var strategyResult sql.NullString
	var countResult sql.NullInt64

	sql := `
SELECT strategy, COUNT(*)
FROM shorturl
GROUP BY strategy
	`

	rows, err := s.DB.Query(sql)
	if err != nil {
		log.Printf("Failed querying database for strategy statistics: %v\n", err)
	}
	defer rows.Close()
	for rows.Next() {
		if err = rows.Scan(&strategyResult, &countResult); err != nil {
			log.Fatal(err)
		}
		validStrategy, ok := Strategies[strategyResult.String]
		if !ok {
			log.Fatalf("'%s' is not a valid strategy\n", strategyResult)
		}
		ret[validStrategy.Name()] = countResult.Int64
	}

	for name, _ := range Strategies {
		_, ok := ret[name]
		if !ok {
			ret[name] = 0
		}
	}
	return ret
}
