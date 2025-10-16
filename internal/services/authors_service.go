package services

import (
	"context"
	"errors"

	"github.com/Luiz-Hen-Reis/go-book-library/internal/store/pgstore"
	"github.com/Luiz-Hen-Reis/go-book-library/internal/usecases/authors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	ErrDuplicatedName = errors.New("this author already exists")
	ErrUnexpectedError = errors.New("something unexpected happened")
	ErrAuthorNotFound = errors.New("author not found for this ID")
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

func (as *AuthorService) CreateAuthor(ctx context.Context, name, bio string) (authors.DefaultAuthorRes, error) {
	args := pgstore.CreateAuthorParams{
		Name: name,
		Bio:  pgtype.Text{String: bio, Valid: true},
	}
	
	author, err := as.queries.CreateAuthor(ctx, args)

	if err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return authors.DefaultAuthorRes{}, ErrDuplicatedName
		}

		return authors.DefaultAuthorRes{}, err
	}

	return authors.DefaultAuthorRes{
		ID:   author.ID.String(),
		Name: author.Name,
		Bio:  author.Bio.String,
	}, nil
}

func (as *AuthorService) ListAuthors(ctx context.Context) ([]authors.DefaultAuthorRes, error) {
	authorsList, err := as.queries.ListAuthors(ctx)

	if err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) {
			return []authors.DefaultAuthorRes{}, ErrUnexpectedError
		}

		return []authors.DefaultAuthorRes{}, err
	}

	var authorResList []authors.DefaultAuthorRes

	for _, v := range authorsList {
		authorRes := authors.DefaultAuthorRes{
			ID:   v.ID.String(),
			Name: v.Name,
			Bio:  v.Bio.String,
		}

		authorResList = append(authorResList, authorRes)
	}

	return authorResList, nil
}

func (as *AuthorService) GetAuthorByID (ctx context.Context, id string) (authors.DefaultAuthorRes, error) {
	uuidUID, err := uuid.Parse(id)

	if err != nil {
		return authors.DefaultAuthorRes{}, err
	}

	author, err := as.queries.GetAuthorByID(ctx, uuidUID)

	if err != nil {
		return authors.DefaultAuthorRes{}, ErrAuthorNotFound
	}

	return authors.DefaultAuthorRes{
		ID:   id,
		Name: author.Name,
		Bio:  author.Bio.String,
	}, nil
}

func (as *AuthorService) DeleteAuthorByID (ctx context.Context, id string) error  {
	uuidUID, err := uuid.Parse(id)

	if err != nil {
		 return err
	}

	rows, err := as.queries.DeleteAuthor(ctx, uuidUID)

	if err != nil {
		return ErrUnexpectedError
	}

	if rows == 0 {
		return ErrAuthorNotFound
	}
	
	return nil
}