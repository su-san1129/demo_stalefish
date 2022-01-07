type TokenFilter interface {
	Filter(TokenStream) TokenStream
}

type LowercaseFilter struct {}

func (f LowercaseFilter) Filter(tokenStream TokenStream) TokenStream {
	r := make([]Token, tokenStream.size())
	for i, token := range tokenStream.Tokens {
		lower := strings.ToLower(token.Term)
		r[i] := NewToken(lower)
	}
	return NewTokenStream(r)
}

type StopWordFilter struct {
	stopWords []string
}

func (f StopWordFilter) Filter(tokenStream TokenStream) TokenStream {
	stopwords := make(map[string]struct{})
	for _, w := range f.stopWords {
		stopwords[w] = struct{}{}
	}
	r := make([]Token, 0, tokenStream.Size())
	for _, token := range tokenStream.Tokens {
		if _, ok := stopwords[token.Term]; !ok {
			r = append(r, token)
		}
	}
	return NewTokenStream(r)
}

type RomajiReadingformFilter struct {}

func (f RomajiReadingformFilter) Filter(tokenStream TokenStream) {
	for i, token := range tokenStream.Tokens {
		tokenStream.Tokens[i].Term = jaconv.ToHebon(
			jaconv.KatakanaToHiragana(token.Kana)
		)
	}
	return tokenStream
}