package application

import (
	"github.com/jnates/crud_golang/internal/domain/model"
	"github.com/jnates/crud_golang/internal/domain/ports"
)

type UserService struct {
	repo ports.UserRepository
}

func NewUserService(repo ports.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Get(id int64) (*model.User, error) {
	return s.repo.GetByID(id)
}

func (s *UserService) Create(user *model.User) (int64, error) {
	return s.repo.Create(user)
}

func (s *UserService) Update(user *model.User) error {
	return s.repo.Update(user)
}

func (s *UserService) Delete(id int64) error {
	return s.repo.Delete(id)
}

func (s *UserService) List(offset, limit int, filter map[string]interface{}) ([]*model.User, error) {
	return s.repo.List(offset, limit, filter)
}
