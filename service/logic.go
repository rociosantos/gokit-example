package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type service struct {
	repository Repository
	logger *logrus.Logger
}

func NewService(repo Repository, logger *logrus.Logger) Service {
	return &service{
		repository: repo,
		logger: logger,
	}
}

func (s service) CreateUser(ctx context.Context, email string, password string) (string, error){
	s.logger.Info("method", "CreateUser")

	uuid, errUUID := uuid.NewRandom()
	if errUUID != nil {
		return "", errUUID
	}

	user := User{
		ID: uuid.String(),
		Email: email,
		Password: password,
	}

	err := s.repository.CreateUser(ctx, user); 
	if err != nil {
		s.logger.Debug("err", err)
		return "", err
	}

	s.logger.Debug("created user with id: ", uuid)

	return uuid.String(), nil
}

func (s service) GetUser(ctx context.Context, id string) (string, error){
	s.logger.Info("method", "GetUser")

	email, err := s.repository.GetUser(ctx, id)

	if err != nil {
		s.logger.Debug("err", err)
		return "", err
	}

	s.logger.Info("Get user", id)
	
	return email, nil
}