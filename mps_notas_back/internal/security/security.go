package security

import "golang.org/x/crypto/bcrypt"

// Hash recebe uma string com a senha do usuario e retorna o hash dessa senha para ser armazenado pelo sistema
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// VerifyPassword recebe a senha e o hash e compara se ambos são equivalentes
func VerifyPassword(password, passwordHash string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
}
