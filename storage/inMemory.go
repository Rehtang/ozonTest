package storage

import (
	"fmt"
	"sync"
)

type InMemoryStorage struct {
	urls map[string]string
	mu   sync.RWMutex
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		urls: make(map[string]string),
	}
}

func (s *InMemoryStorage) SaveURL(shortURL, originalURL string) error {
	exists, err := s.CheckURLExists(originalURL)
	if err != nil {
		return fmt.Errorf("failed to check URL existence: %w", err)
	}
	if exists {
		return fmt.Errorf("URL already exists in the storage")
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	s.urls[shortURL] = originalURL
	return nil
}

func (s *InMemoryStorage) GetURL(shortURL string) (string, bool, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	originalURL, ok := s.urls[shortURL]
	return originalURL, ok, nil
}

func (s *InMemoryStorage) Close() {
	//placeholder
}

func (s *InMemoryStorage) CheckURLExists(originalURL string) (bool, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, url := range s.urls {
		if url == originalURL {
			return true, nil
		}
	}
	return false, nil
}
