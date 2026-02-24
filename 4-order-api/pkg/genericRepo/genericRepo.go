package genericRepo

import "gorm.io/gorm"

type GenericRepository[T any] struct {
	Db *gorm.DB
}

// NewRepository создает новый экземпляр репозитория
func NewRepository[T any](db *gorm.DB) *GenericRepository[T] {
	return &GenericRepository[T]{Db: db}
}

// Create создает новую запись в базе данных
func (r *GenericRepository[T]) Create(entity *T) error {
	return r.Db.Create(entity).Error
}

// GetByID получает запись по ID
func (r *GenericRepository[T]) Get(id uint) (*T, error) {
	var entity T
	result := r.Db.First(&entity, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &entity, nil
}

// Update обновляет запись
func (r *GenericRepository[T]) Update(id uint, entity *T) error {
	return r.Db.Model(entity).Where("id = ?", id).Updates(entity).Error
}

// Delete удаляет запись по ID
func (r *GenericRepository[T]) Delete(id uint) error {
	return r.Db.Delete(new(T), id).Error
}

// GetAll получает все записи с пагинацией
func (r *GenericRepository[T]) GetAll() ([]T, error) {
	var entities []T
	result := r.Db.Find(&entities)
	if result.Error != nil {
		return nil, result.Error
	}
	return entities, nil
}
