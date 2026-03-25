package search

// Token represents one normalized searchable term and its position in the stream.
type Token struct {
	Term string
	Pos  int
}

// Tokenizer defines a stable contract for indexing and query parsing.
type Tokenizer interface {
	Tokenize(input string) []Token
}

// TokenizerConfig controls tokenizer behavior.
type TokenizerConfig struct {
	MinLen       int
	Stopwords    map[string]struct{}
	KeepHashtags bool
	KeepMentions bool
}
