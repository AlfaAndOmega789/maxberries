package usecase

import "maxberries/catalog_service/internal/domain/entity"

type CategoryUsecase interface {
	GetAll() ([]entity.Category, error)
	GetByID(id string) (entity.Category, error)
	Create(p entity.Category, err error) (string, error)
	Update(id string, p entity.Category) error
	Delete(id string) error
}
