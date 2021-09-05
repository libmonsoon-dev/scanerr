package source

type StringMatcher interface {
	// FilterMatched should exclude not relevant to originalError Strings and sort remaining by matched fragment position
	FilterMatched(originalError string, strings []String) ([]String, [][2]int)
}
