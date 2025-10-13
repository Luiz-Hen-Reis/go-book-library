package services

import (
	"context"
	"errors"

	"github.com/Luiz-Hen-Reis/go-book-library/internal/store/pgstore"
	"github.com/Luiz-Hen-Reis/go-book-library/internal/usecases/authors"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	ErrUnexpectedError = errors.New("unexpected error when trying to create a new author")
)

type AuthorService struct {
	pool *pgxpool.Pool
	queries *pgstore.Queries
}

func NewAuthorService(pool *pgxpool.Pool) AuthorService {
	return AuthorService{
		pool:    pool,
		queries: pgstore.New(pool),
	}
}

func (as *AuthorService) CreateAuthor(ctx context.Context, name, bio string) (authors.CreateAuthorRes, error) {
	args := pgstore.CreateAuthorParams{
		Name: name,
		Bio:  pgtype.Text{String: bio, Valid: true},
	}

	author, err := as.queries.CreateAuthor(ctx, args)

	if err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) {
			// criar um logger para logar a mensagem de erro
			return authors.CreateAuthorRes{}, ErrUnexpectedError
		}

		return authors.CreateAuthorRes{}, err
	}

	return authors.CreateAuthorRes{
		ID:   author.ID.String(),
		Name: author.Name,
		Bio:  author.Bio.String,
	}, nil
}