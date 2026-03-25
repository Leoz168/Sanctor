package search

// DefaultStopwords returns a basic English stopword set for MVP relevance tuning.
func DefaultStopwords() map[string]struct{} {
	words := []string{
		"a", "an", "and", "are", "as", "at",
		"be", "by",
		"for", "from",
		"in", "is", "it",
		"of", "on", "or",
		"that", "the", "this", "to",
		"was", "with",
	}

	set := make(map[string]struct{}, len(words))
	for _, word := range words {
		set[word] = struct{}{}
	}

	return set
}
