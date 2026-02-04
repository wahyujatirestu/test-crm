package services

import (
	"context"
	"errors"

	"test_crm/dto"
	"test_crm/repository"
	"test_crm/utils"
)

type AuthService interface {
	Login(ctx context.Context, req dto.LoginRequest) (*dto.LoginResponse, error)
}

type authService struct {
	repo repository.MembershipRepository
	jwt  utils.JWTService
}

func NewAuthService(repo repository.MembershipRepository, jwt utils.JWTService) AuthService {
	return &authService{repo, jwt}
}

func (s *authService) Login(ctx context.Context, req dto.LoginRequest) (*dto.LoginResponse, error) {

	user, err := s.repo.FindByContactValue(ctx, req.Username)
	if err != nil {
		return nil, errors.New("invalid username or password")
	}

	hashed := utils.HashMD5(req.Password)
	if user.Password != hashed {
		return nil, errors.New("invalid username or password")
	}

	token, err := s.jwt.GenerateToken(user.MembershipID)
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{
		AccessToken: token,
	}, nil
}
