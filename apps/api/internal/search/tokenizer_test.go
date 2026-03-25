package search

import "testing"

func TestSimpleTokenizer_DefaultBehavior(t *testing.T) {
	tokenizer := NewSimpleTokenizer(TokenizerConfig{})
	tokens := tokenizer.Tokenize("The quick brown fox jumps in the park")

	want := []Token{
		{Term: "quick", Pos: 0},
		{Term: "brown", Pos: 1},
		{Term: "fox", Pos: 2},
		{Term: "jumps", Pos: 3},
		{Term: "park", Pos: 4},
	}

	assertTokensEqual(t, tokens, want)
}

func TestSimpleTokenizer_HashtagsAndMentions(t *testing.T) {
	tokenizer := NewSimpleTokenizer(TokenizerConfig{
		KeepHashtags: true,
		KeepMentions: true,
		Stopwords:    map[string]struct{}{},
	})

	tokens := tokenizer.Tokenize("Find #Sanctor updates from @tobias")

	want := []Token{
		{Term: "find", Pos: 0},
		{Term: "#sanctor", Pos: 1},
		{Term: "updates", Pos: 2},
		{Term: "from", Pos: 3},
		{Term: "@tobias", Pos: 4},
	}

	assertTokensEqual(t, tokens, want)
}

func TestSimpleTokenizer_MinLenAndStopwords(t *testing.T) {
	tokenizer := NewSimpleTokenizer(TokenizerConfig{
		MinLen: 3,
		Stopwords: map[string]struct{}{
			"group": {},
		},
	})

	tokens := tokenizer.Tokenize("Go group ranking by ai")

	want := []Token{
		{Term: "ranking", Pos: 0},
	}

	assertTokensEqual(t, tokens, want)
}

func assertTokensEqual(t *testing.T, got, want []Token) {
	t.Helper()

	if len(got) != len(want) {
		t.Fatalf("token length mismatch: got=%d want=%d", len(got), len(want))
	}

	for i := range got {
		if got[i] != want[i] {
			t.Fatalf("token mismatch at index %d: got=%+v want=%+v", i, got[i], want[i])
		}
	}
}
