package config

// Config armazena as configurações da aplicação
type Config struct {
	Port           int
	AllowedOrigins []string
}

// New retorna uma nova instância de configuração com valores padrão
func New() *Config {
	return &Config{
		Port:           3000,
		AllowedOrigins: []string{"http://localhost:3001"}, // CORS, embora não tenha mais necessidade acredito eu
	}
}