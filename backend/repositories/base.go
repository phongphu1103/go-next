package repositories

import (
    "context"
    "gorm.io/gorm"
)

type BaseRepository[T any] struct {
    db *gorm.DB
}

func NewBaseRepository[T any](db *gorm.DB) *BaseRepository[T] {
    return &BaseRepository[T]{db: db}
}

func (b *BaseRepository[T]) DB() *gorm.DB {
    return b.db
}

func (b *BaseRepository[T]) Create(ctx context.Context, model *T) error {
    return b.db.WithContext(ctx).Create(model).Error
}

func (b *BaseRepository[T]) Update(ctx context.Context, model *T) error {
    return b.db.WithContext(ctx).Save(model).Error
}

func (b *BaseRepository[T]) Delete(ctx context.Context, model *T) error {
    return b.db.WithContext(ctx).Delete(model).Error
}

func (b *BaseRepository[T]) FindByID(ctx context.Context, model *T, id any) error {
    return b.db.WithContext(ctx).First(model, id).Error
}

func (b *BaseRepository[T]) FindAll(ctx context.Context, models *[]T) error {
    return b.db.WithContext(ctx).Find(models).Error
}

func (b *BaseRepository[T]) Paginate(ctx context.Context, models *[]T, page, pageSize int) (int64, error) {
    var total int64
    var model T

    query := b.db.WithContext(ctx)
    // count on model type
    query.Model(&model).Count(&total)

    offset := (page - 1) * pageSize
    err := query.Offset(offset).Limit(pageSize).Find(models).Error

    return total, err
}

func (b *BaseRepository[T]) Transaction(fc func(tx *gorm.DB) error) error {
    return b.db.Transaction(fc)
}

