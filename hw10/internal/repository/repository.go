package repository

import (
	"context"
	"homeworks/hw10/api"
	"homeworks/hw10/internal/entity"
)

type Repository interface {
	CreateUser(ctx context.Context, u *entity.User) (string, error)
	GetUser(ctx context.Context, email string) (*entity.User, error)
	UpdateUser(ctx context.Context, id string, req *api.UpdateUserRequest) error
	DeleteUser(ctx context.Context, id string) error

	CreateBook(ctx context.Context, req *api.BookRequest) (string, error)
	GetAllBooks(ctx context.Context) ([]entity.Book, error)
	GetBookById(ctx context.Context, id string) (*entity.Book, error)
	DeleteBook(ctx context.Context, id string) error
	UpdateBook(ctx context.Context, id string, req *api.BookRequest) error

	CreateAuthor(ctx context.Context, req *api.AuthorRequest) (string, error)
	GetAllAuthors(ctx context.Context) ([]entity.Author, error)
	GetAuthorById(ctx context.Context, id string) (*entity.Author, error)
	DeleteAuthor(ctx context.Context, id string) error
	UpdateAuthor(ctx context.Context, id string, req *api.AuthorRequest) error

	CreateFilePath(ctx context.Context, req *api.FilePathRequest) (string, error)
	GetAllFilePaths(ctx context.Context) ([]entity.FilePath, error)
	GetFilePathById(ctx context.Context, id string) (*entity.FilePath, error)
	DeleteFilePath(ctx context.Context, id string) error
	UpdateFilePath(ctx context.Context, id string, req *api.FilePathRequest) error
}
