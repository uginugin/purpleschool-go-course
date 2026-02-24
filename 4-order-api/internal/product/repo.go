package product

import (
	"4-order-api/pkg/genericRepo"

	"gorm.io/gorm"
)

type ProductRepo struct {
	*genericRepo.GenericRepository[Product]
	Db *gorm.DB
}

type ProductRepoDeps struct {
	Db *gorm.DB
}

func NewProductRepo(deps *ProductRepoDeps) *ProductRepo {
	return &ProductRepo{
		Db:                deps.Db,
		GenericRepository: genericRepo.NewRepository[Product](deps.Db),
	}
}
