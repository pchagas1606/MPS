package security

import "golang.org/x/crypto/bcrypt"

// SecurityProvider define a interface para operações de segurança
type SecurityProvider interface {
    Hash(password string) ([]byte, error)
    VerifyPassword(password, passwordHash string) error
}

// BcryptProvider implementa a segurança usando bcrypt
type BcryptProvider struct{}

// NewBcryptProvider cria uma nova instância do provedor bcrypt
func NewBcryptProvider() *BcryptProvider {
    return &BcryptProvider{}
}

// Hash recebe uma string com a senha do usuario e retorna o hash dessa senha para ser armazenado pelo sistema
func (b *BcryptProvider) Hash(password string) ([]byte, error) {
    return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// VerifyPassword recebe a senha e o hash e compara se ambos são equivalentes
func (b *BcryptProvider) VerifyPassword(password, passwordHash string) error {
    return bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
}