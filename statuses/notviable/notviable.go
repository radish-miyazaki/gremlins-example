// Package notviable demonstrates the NOT VIABLE status.
// Mutating the `+` of a string concatenation into `-` produces a type error and fails to compile.
package notviable

// Concat returns a + b (string concatenation). The `+`->`-` mutation cannot
// compile for strings, so gremlins classifies it as NOT VIABLE.
func Concat(a, b string) string {
	return a + b
}
