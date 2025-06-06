package repository

import (
    "context"
    "fmt"
    "user-service/internal/domain"
    "github.com/jackc/pgx/v5/pgxpool"
)

type UserRepo struct {
    db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) *UserRepo {
    return &UserRepo{db: db}
}
func (r *UserRepo) InitializeConnectionPool(minConns, maxConns int) {
    r.db.Config().MinConns = int32(minConns)
    r.db.Config().MaxConns = int32(maxConns)
}

func (r *UserRepo) CreateUser(ctx context.Context, user *domain.User) error {
    _, err := r.db.Exec(ctx, insertUserQuery,user.Username, user.Email, user.Password)
    return err
}

func (r *UserRepo) FindByCredentials(ctx context.Context, username string) (*domain.User, error) {
    row := r.db.QueryRow(ctx, findUserByUsernameQuery, username)

    var user domain.User
    err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Tokens)
    if err != nil {
        return nil, fmt.Errorf("user not found: %w", err)
    }

    return &user, nil
}

func (r *UserRepo) UpdateTokens(ctx context.Context, userID string, token string) error {
    _, err := r.db.Exec(ctx, updateUserTokensQuery,userID,token)
    return err
}

func (r *UserRepo) GetProfile(ctx context.Context, userID string) (*domain.User, error) {
    var user domain.User
    row := r.db.QueryRow(ctx, getProfile, userID)

    err := row.Scan(&user.ID, &user.Username, &user.Email)
    if err != nil {
        return nil, err
    }

    
    return &user, nil
}
