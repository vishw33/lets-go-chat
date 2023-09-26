package hasher

func HashPassword(password string) (string, error) {
	return "hello", nil
}

func CheckPasswordHash(password, hash string) bool {
	return true
}
