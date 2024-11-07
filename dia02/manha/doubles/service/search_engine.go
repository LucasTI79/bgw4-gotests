package service

import (
	"errors"
	"fmt"
)

// nossa service que vai chamar a repository e fazer algumas validacoes
type SearchEngine struct {
	// nossa repository que contera os metodos para manipular os dados
	searchEngine searchEngine
}

type searchEngine interface {
	SearchByName(name string) string
	SearchByPhone(phone string) string
	AddEntry(name, phone string) error
}

func (s *SearchEngine) SearchByName(name string) string {
	return s.searchEngine.SearchByName(name)
}

func (s *SearchEngine) SearchByPhone(phone string) (string, error) {
	if len(phone) < 11 {
		return "", errors.New("phone must have 11 characters")
	}
	return s.searchEngine.SearchByPhone(phone), nil
}

func (s *SearchEngine) AddEntry(name, phone string) error {
	return s.searchEngine.AddEntry(name, phone)
}

func (s *SearchEngine) GetVersion() string {
	return fmt.Sprintf("version: %s", "1.0.0")
}

func NewEngine(searchEngine searchEngine) *SearchEngine {
	return &SearchEngine{searchEngine}
}
