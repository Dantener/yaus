package service

import (
	"yaus/internal/entity"
	"yaus/internal/repository"
)

type Convertation interface {
	CreateShortURL(url entity.Item) (string, error)
	GetOriginURLByShortURL(url entity.Item) (string, error)
}

type Service struct {
	Convertation
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Convertation: NewConvertationService(repos.Convertation),
	}
}
