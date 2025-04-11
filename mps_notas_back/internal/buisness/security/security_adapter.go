package security

// SecurityAdapter adapta as operações de segurança para o sistema existente
type SecurityAdapter struct {
    provider SecurityProvider
}

// NewSecurityAdapter cria uma nova instância do adapter de segurança
func NewSecurityAdapter(provider SecurityProvider) *SecurityAdapter {
    return &SecurityAdapter{
        provider: provider,
    }
}

// Hash adapta a função de hash para o formato existente
func Hash(password string) ([]byte, error) {
    provider := NewBcryptProvider()
    return provider.Hash(password)
}

// VerifyPassword adapta a função de verificação para o formato existente
func VerifyPassword(password, passwordHash string) error {
    provider := NewBcryptProvider()
    return provider.VerifyPassword(password, passwordHash)
}