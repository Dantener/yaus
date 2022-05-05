package service

import (
	"math/rand"
	"yaus/internal/entity"
	"yaus/internal/repository"
	"yaus/pkg/base62"
)

type ConvertationService struct {
	repo repository.Convertation
}

func NewConvertationService(repo repository.Convertation) *ConvertationService {
	return &ConvertationService{repo: repo}
}

func (s *ConvertationService) CreateShortURL(item entity.Item) (string, error) {
	shortURL, err := s.repo.FindShortURLByOriginURL(item)
	if err != nil {
		return "", err
	}

	if shortURL != "" {
		return shortURL, nil
	}
	item.Id = generateId()
	item.ShortURL += base62.Encode(item.Id)
	return s.repo.CreateShortURL(item)
}

func (s *ConvertationService) GetOriginURLByShortURL(item entity.Item) (string, error) {
	originURL, err := s.repo.FindOriginUrlByShortURL(item)
	if err != nil {
		return "", err
	}

	return originURL, nil
}

func generateId() (id uint32) {
	return rand.Uint32()
}
