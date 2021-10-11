package internal

import "fmt"

// WrapError wraps a plain error into a custom error
// Use only for internal error
func WrapError(customErr string, originalErr error) error {
	err := fmt.Errorf("%s: %v", customErr, originalErr)
	return err
}
