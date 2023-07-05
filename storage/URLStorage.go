package storage

// интерфейс для хранилища
type URLStorage interface {
	SaveURL(shortURL, originalURL string) error
	GetURL(shortURL string) (string, bool, error)
	CheckURLExists(originalURL string) (bool, error)
	Close()
}
