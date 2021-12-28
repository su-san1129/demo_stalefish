package main

import (
	"fmt"
	"github.com/kotaroooo0/stalefish"
)

func main() {
	db, _ := stalefish.NewDBClient(
		stalefish.NewDBConfig(
			"root",
			"password",
			"127.0.0.1",
			"3306",
			"stalefish",
		),
	)

	storage := stalefish.NewStorageRdbImpl(db)
	analyzer := stalefish.NewAnalyzer(
		[]stalefish.CharFilter{
			stalefish.NewMappingCharFilter(
				map[string]string{":(": "sad"},
			),
		},
		stalefish.NewStandardTokenizer(),
		[]stalefish.TokenFilter{
			stalefish.NewLowercaseFilter(),
			stalefish.NewStopWordFilter(
				[]string{"i", "my", "me", "the", "a", "for"},
			),
		},
	)

	fmt.Println(analyzer.Analyze("I feel TIRED :("))

	indexer := stalefish.NewIndexer(
		storage, analyzer, 1,
	)
	for _, body := range []string{
		"Ruby PHP JS",
		"Go Ruby",
		"Ruby Go PHP",
		"Go PHP",
	} {
		indexer.AddDocument(
			stalefish.NewDocument(body),
		)
	}

	sorter := stalefish.NewTfIdfSorter(storage)
	mq := stalefish.NewMatchQuery(
		"Go Ruby",
		stalefish.OR,
		analyzer,
		sorter,
	)

	mseacher := mq.Searcher(storage)
	result, _ := mseacher.Search()
	fmt.Println(result)

	pq := stalefish.NewPhraseQuery(
		"go Ruby", analyzer, nil,
	)
	psearcher := pq.Searcher(storage)
	result, _ = psearcher.Search()
	fmt.Println(result)
}