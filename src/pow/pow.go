package pow

type Pow interface {
	// Verify can be used to verify the token
	Verify(string) bool

	// Generate creates a new token putting the given resource into it.
	Generate(string) string

	// Parse should unpack the token and return payload.
	Parse(string) []string
}
