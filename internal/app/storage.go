package app

import "errors"

type Storage interface {
	SaveURL(originURL string) string
	GetOriginURL(shortID string) (string, error)
}

type MemStorage struct {
	urls map[string]string
}

func NewMemStorage() *MemStorage {
	return &MemStorage{
		urls: make(map[string]string),
	}
}

func (s *MemStorage) SaveURL(originURL string) string {
	shortURL := GenerateShortURL(originURL)
	s.urls[shortURL] = originURL
	return shortURL
}

func (s *MemStorage) GetOriginURL(shortID string) (string, error) {
	if originURL, ok := s.urls[shortID]; ok {
		return originURL, nil
	}
	return "", errors.New("URL Not Found")
}
