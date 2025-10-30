package usecase

import "maxberries/catalog_service/internal/domain/entity"

type CategoryUsecase interface {
	GetAll() ([]entity.Category, error)
	Create(p entity.Category, err error)
	Update(id int, p entity.Category) error
	Delete(id int) error
}
