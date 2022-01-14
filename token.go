package main

type TokenID uint64

type Token struct {
	ID   TokenID
	Term string
	Kana string
}

type TokenOption func(*Token)

func NewToken(term string, options ...TokenOption) Token {
	token := Token{Term: term}
	for _, option := range options {
		option(&token)
	}
	return token
}

func setKana(kana string) TokenOption {
	return func(s *Token) {
		s.Kana = kana
	}
}

type TokenStream struct {
	Tokens []Token
}

func NewTokenStream(tokens []Token) TokenStream {
	return TokenStream{
		Tokens: tokens,
	}
}

func (ts TokenStream) Size() int {
	return len(ts.Tokens)
}

func (ts TokenStream) Terms() []string {
	terms := make([]string, ts.Size())
	for i, t := range ts.Tokens {
		terms[i] = t.Term
	}
	return terms
}
