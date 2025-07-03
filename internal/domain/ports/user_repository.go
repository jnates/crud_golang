package ports

import "github.com/jnates/crud_golang/internal/domain/model"

type UserRepository interface {
	GetByID(id int64) (*model.User, error)
	Create(user *model.User) (int64, error)
	Update(user *model.User) error
	Delete(id int64) error
	List(offset, limit int, filter map[string]interface{}) ([]*model.User, error)
}
