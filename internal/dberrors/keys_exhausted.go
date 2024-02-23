package dberrors

import "fmt"

type KeysExhausted struct{}

func (e *KeysExhausted) Error() string {
	return fmt.Sprintf("All Keys Exhausted! Please add new keys")
}
