package strrev

import "errors"

var (
	// ErrInputEmptyStr error input empty string.
	ErrInputEmptyStr = errors.New("strrev: input string is empty")
)

// Reverse returns a reversed string.
func Reverse(s string) (string, error) {
	if s == "" {
		return "", ErrInputEmptyStr
	}
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r), nil
}
