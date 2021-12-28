type Token struct {
	ID TokenID
	Term string
	Kana string
}

type TokenStream struct {
	Tokens []Token
}