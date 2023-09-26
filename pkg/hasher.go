// Package hasher package is used to demonstrate the understanding of
// how packages work in go
package hasher

// HashPassword return the decrypted password if able to decrypt
// else return error
func HashPassword(password string) (string, error) {
	return "hello", nil
}

// CheckPasswordHash return bool based on condition
func CheckPasswordHash(password, hash string) bool {
	return true
}
