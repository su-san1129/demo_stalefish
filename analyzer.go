type Analyzer struct {
	charFilters []CharFilter
	tokenizer Tokenizer
	tokenFilters []TokenFilter
}

func (a Analyzer) Analyzer(s string) TokenStream {
	for _, c := range a.charFilters {
		s = c.Filter(s)
	}
	tokenStream := a.tokenizer.Tokenizer(s)
	for _, f := range a.tokenFilters {
		tokenStream _ f.Filters(tokenStream)
	}
	return tokenStream
}