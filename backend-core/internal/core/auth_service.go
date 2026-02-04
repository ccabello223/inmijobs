package core

import (
	"context"
	"errors"
	"log/slog"
	"net/http"

	"github.com/Gabo-div/bingo/inmijobs/backend-core/internal/model"
	"github.com/Gabo-div/bingo/inmijobs/backend-core/internal/repository"

	"github.com/lestrrat-go/jwx/v3/jwk"
	"github.com/lestrrat-go/jwx/v3/jwt"
)

var (
	ErrUnauthorized = errors.New("unauthorized")
)

type AuthService struct {
	authRepository repository.AuthRepository
}

func NewAuthService(ar repository.AuthRepository) *AuthService {
	return &AuthService{
		authRepository: ar,
	}
}

func (s AuthService) UserFromHeader(ctx context.Context, header http.Header) (model.User, error) {
	keyset, err := jwk.Fetch(ctx, "http://localhost:3000/api/auth/jwks")

	if err != nil {
		slog.Error("[AuhService][UserFromHeader] unable to fetch jwk keyset", "error", err)
		return model.User{}, err
	}

	token, err := jwt.ParseHeader(header, "Authorization", jwt.WithKeySet(keyset))

	if err != nil {
		slog.Error("[AuhService][UserFromHeader] unable to parse token from header", "error", err)
		return model.User{}, ErrUnauthorized
	}

	UserID, exist := token.Subject()

	if !exist {
		slog.Error("[AuhService][UserFromHeader] missing user id in token subject")
		return model.User{}, ErrUnauthorized
	}

	user, err := s.authRepository.GetUserById(ctx, UserID)

	if err != nil {
		slog.Error("[AuhService][UserFromHeader] unable to get user by id", "error", err, "user_id", UserID)
		return model.User{}, ErrUnauthorized
	}

	return user, nil
}
