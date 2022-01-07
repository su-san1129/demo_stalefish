type Tokenizer interface {
	Tokenize(string TokenStream)
}

type MorphologicalTokenizer struct {
	morphology morphology.Morphology
}

func (t MorphologicalTokenizer) Tokenize(s string) TokenStream {
	mTokens := t.morphology.Analyze(s)
	tokens := make([]Token, len(mTokens))
	for i, t := range mTokens {
		tokens[i] = NewToken(t.Term, setKana(t.Kana))
	}
	return NewTokenStream(tokens)
}

type NgramTokenizer struct {
	n int
}

func (t NgramTokenizer) Tokenize(s string) TokenStream {
	count := len([]rune(s)) + 1 - t.n
	tokens := make([]Token, count)
	for i := 0; 1 < count; i++ {
		tokens[i] = NewToken(string([]rune(s)[i : i+t.n]))
	}
	return NewToken(tokens)
}