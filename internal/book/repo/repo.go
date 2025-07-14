package repo

import (
	"context"

	"github.com/natbabo1/sample-gin-api/internal/book"
	"gorm.io/gorm"
)

type Repository interface {
	Create(ctx context.Context, book *book.Book) error
	FindByID(ctx context.Context, id int64) (*book.Book, error)
	FindAll(ctx context.Context) ([]*book.Book, error)
	// Update(ctx context.Context, book *book.Book) error
	// Delete(ctx context.Context, id int64) error
}

type repo struct {
	db *gorm.DB
}

func New(db *gorm.DB) Repository {
	return &repo{db}
}

func (r *repo) Create(ctx context.Context, book *book.Book) error {
	return r.db.WithContext(ctx).Create(book).Error
}

func (r *repo) FindByID(ctx context.Context, id int64) (out *book.Book, err error) {
	err = r.db.WithContext(ctx).First(&out, id).Error
	return
}

func (r *repo) FindAll(ctx context.Context) (out []*book.Book, err error) {
	err = r.db.WithContext(ctx).Find(&out).Error
	return
}
