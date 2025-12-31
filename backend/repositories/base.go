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

func (r *BaseRepository[T]) DB() *gorm.DB {
    return r.db
}
// common query initializer
func (r *BaseRepository[T]) Query(ctx context.Context) *gorm.DB {
    return r.db.WithContext(ctx)
}

func (r *BaseRepository[T]) QueryWithPreload(ctx context.Context, preloads ...string) *gorm.DB {
    query := r.db.WithContext(ctx)
    // apply preloads
    for _, p := range preloads {
        query = query.Preload(p)
    }

    return query
}

/**
 * Execute an SQL statement.
 */
func (r *BaseRepository[T]) exec (db *gorm.DB) error {
    return db.Error
}

func (r *BaseRepository[T]) Create(ctx context.Context, model *T) error {
    return r.exec(r.Query(ctx).Create(model))
}

func (r *BaseRepository[T]) Update(ctx context.Context, model *T) error {
    return r.exec(r.Query(ctx).Save(model))
}

func (r *BaseRepository[T]) Delete(ctx context.Context, model *T) error {
    return r.exec(r.Query(ctx).Delete(model))
}

func (r *BaseRepository[T]) FindByID(ctx context.Context, model *T, id any) error {
    return r.exec(r.Query(ctx).First(model, id))
}

func (r *BaseRepository[T]) FindByIDWithPreload(ctx context.Context, model *T, id any, preloads ...string) error {
    query := r.QueryWithPreload(ctx, preloads...)

    return r.exec(query.First(model, id))
}

func (r *BaseRepository[T]) FindAll(ctx context.Context, models *[]T) error {
    return r.exec(r.Query(ctx).Find(models))
}

func (r *BaseRepository[T]) FindAllWithPreload(ctx context.Context, models *[]*T, preloads ...string) error {
    query := r.QueryWithPreload(ctx, preloads...)

    return r.exec(query.Find(models))
}

func (r *BaseRepository[T]) Paginate(ctx context.Context, models *[]T, page, pageSize int) (int64, error) {
    var total int64
    var model T

    query := r.Query(ctx)
    // count on model type
    query.Model(&model).Count(&total)

    offset := (page - 1) * pageSize
    err := r.exec(query.Offset(offset).Limit(pageSize).Find(models))

    return total, err
}

func (r *BaseRepository[T]) Transaction(fc func(tx *gorm.DB) error) error {
    return r.db.Transaction(fc)
}

