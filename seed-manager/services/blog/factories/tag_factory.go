package factories

import (
	"seeder-manager/models"

	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
)

type TagFactory struct{}

func NewTagFactory() *TagFactory {
	return &TagFactory{}
}

func (f *TagFactory) Create() (*models.BlogTagModel, error) {
	return &models.BlogTagModel{
		ID:   uuid.New().String(),
		Name: faker.Word(),
	}, nil
}

func (f *TagFactory) CreateMany(count int) ([]*models.BlogTagModel, error) {
	tags := make([]*models.BlogTagModel, 0, count)
	for i := 0; i < count; i++ {
		tag, err := f.Create()
		if err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}
	return tags, nil
}
