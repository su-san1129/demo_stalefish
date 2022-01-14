package main

import "github.com/kotaroooo0/stalefish"

type InvertedIndex map[TokenID]stalefish.PostingList
type PostingList struct {
	Postings *Postings
}

type Postings struct {
	DocumentId stalefish.DocumentID
	Positions  []uint64
	Next       *Postings
}
