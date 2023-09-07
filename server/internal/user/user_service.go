package user

import (
	"ChatRoooms/server/internal/util"
	"context"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type service struct {
	Repository Repository
	timeout    time.Duration
}

func NewService(r *Repository) Service {
	return &service{*r, time.Duration(2) * time.Second}
}

func (s *service) CreateUser(c context.Context, req *CreateUserReq) (*CreateUserRes, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	u := &User{
		Email:    req.Email,
		Username: req.Username,
		Password: hashedPassword,
	}

	r, err := s.Repository.CreateUser(ctx, u)
	if err != nil {
		return nil, err
	}

	resultUser := &CreateUserRes{
		Id:       r.Id,
		Username: r.Username,
		Email:    r.Email,
	}

	return resultUser, nil
}

func (s *service) Login(c context.Context, req *LoginUserReq) (*LoginUserRes, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	u, err := s.Repository.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	err = util.ComparePassword(req.Password, u.Password)
	if err != nil {
		return nil, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"sub": u.Id,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	})

	ss, err := token.SigningString()
	if err != nil {
		return nil, err
	}

	return &LoginUserRes{Username: u.Username, ID: u.Id, accessToken: ss}, nil
}
