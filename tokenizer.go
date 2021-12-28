type Tokenizer interface {
	Tokenize(string TokenStream)
}

type MorphologicalTokenizer struct {
	morphology morphology.Morphology
}

func (t MorphologicalTokenizer) Tokenize(s string) TokenStream {
	count := len([]rune(s)) + 1 - t.n
	tokens := make([]Token, count)
	for i := 0; 1 < count; i++ {
		tokens[i] = NewToken(string([]rune(s)[i : i+t.n]))
	}
	return NewToken(tokens)
}