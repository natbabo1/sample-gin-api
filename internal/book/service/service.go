package service

import (
	"context"
	"errors"

	"github.com/natbabo1/sample-gin-api/internal/book"
	"github.com/natbabo1/sample-gin-api/internal/book/repo"
)

type Service interface {
	Create(ctx context.Context, book book.Book) (*book.Book, error)
	List(ctx context.Context) ([]*book.Book, error)
	FindByID(ctx context.Context, id int64) (*book.Book, error)
}

type service struct {
	repo repo.Repository
}

func New(repo repo.Repository) Service {
	return &service{repo}
}

var ErrInvalidBook = errors.New("invalid book: title and author are required")

func (s *service) Create(ctx context.Context, book book.Book) (*book.Book, error) {
	if book.Title == "" || book.Author == "" {
		return nil, ErrInvalidBook
	}
	if err := s.repo.Create(ctx, &book); err != nil {
		return nil, err
	}

	return &book, nil
}

func (s *service) List(ctx context.Context) ([]*book.Book, error) {
	books, err := s.repo.FindAll(ctx)

	if err != nil {
		return nil, err
	}
	return books, nil
}

func (s *service) FindByID(ctx context.Context, bookId int64) (*book.Book, error) {
	book, err := s.repo.FindByID(ctx, bookId)

	if err != nil {
		return nil, err
	}

	return book, nil
}
