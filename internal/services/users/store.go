package users

import (
	"context"
	repo "ecom-local/internal/adapters/postgresql/sqlc"
)

type Store struct {
	queries *repo.Queries
}

func NewStore(db repo.DBTX) *Store {
	return &Store{
		queries: repo.New(db),
	}
}

func (s *Store) CreateUser(ctx context.Context, u *User) error {
	dbUser, err := s.queries.CreateUser(ctx, repo.CreateUserParams{
		Email:        u.Email,
		PasswordHash: u.PasswordHash,
		FirstName:    u.FirstName,
		LastName:     u.LastName,
	})

	if err != nil {
		return err
	}

	u.ID = dbUser.ID
	u.CreatedAt = dbUser.CreatedAt.Time
	return nil
}

func (s *Store) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	dbUser, err := s.queries.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:           dbUser.ID,
		Email:        dbUser.Email,
		PasswordHash: dbUser.PasswordHash,
		FirstName:    dbUser.FirstName,
		LastName:     dbUser.LastName,
		CreatedAt:    dbUser.CreatedAt.Time,
	}, nil
}
