package memory

// Token struct for token repo
type Token struct {
	secretKey string
}

// TokenManager interface
type TokenManager interface {
	SaveToken(token string) error
	ClearToken() error
	GetToken() (string, error)
}

// NewManager create manager
func NewManager(secretKey string) TokenManager {
	return &Token{secretKey: secretKey}
}
