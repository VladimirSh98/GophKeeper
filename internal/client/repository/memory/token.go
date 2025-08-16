package memory

import "sync"

type TokenManager struct {
	mu    sync.Mutex
	token string
}

func (t *TokenManager) AddToken(token string) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.token = token
}
