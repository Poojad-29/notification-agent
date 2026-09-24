package middleware

func ValidateToken(token string) bool {
	return token == "demo-token"
}
