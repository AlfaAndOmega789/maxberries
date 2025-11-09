package usecase

import "maxberries/catalog_service/internal/domain/entity"

type ProductUsecase interface {
	GetAll() ([]entity.Product, error)
	GetByID(id string) (entity.Product, error)
	Create(p entity.Product) (string, error)
	Update(id string, p entity.Product) error
	Delete(id string) error
}
