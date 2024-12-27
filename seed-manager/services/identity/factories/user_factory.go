package factories

import (
	"time"

	"seeder-manager/models"

	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
)

type UserFactory struct{}

func NewUserFactory() *UserFactory {
	return &UserFactory{}
}

func (f *UserFactory) Create() (*models.UserModel, error) {
	return &models.UserModel{
		ID:        uuid.New().String(),
		Email:     faker.Email(),
		Username:  faker.Username(),
		FirstName: faker.FirstName(),
		LastName:  faker.LastName(),
		CreatedAt: time.Now(),
	}, nil
}

func (f *UserFactory) CreateMany(count int) ([]*models.UserModel, error) {
	users := make([]*models.UserModel, 0, count)
	for i := 0; i < count; i++ {
		user, err := f.Create()
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
