package factories

import (
	"fmt"
	"math/rand"
	"time"

	"seeder-manager/models"
	"seeder-manager/reference_store"

	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
)

type BlogPostFactory struct {
	refStore *reference_store.InMemoryStore
}

func NewBlogPostFactory(refStore *reference_store.InMemoryStore) *BlogPostFactory {
	return &BlogPostFactory{
		refStore: refStore,
	}
}

func (f *BlogPostFactory) Create() (*models.BlogPostModel, error) {
	// Get all user emails from the store
	userEmails := f.refStore.GetAllUserEmails()
	if len(userEmails) == 0 {
		return nil, fmt.Errorf("no users available in reference store")
	}

	// Pick a random user email
	randomEmail := userEmails[rand.Intn(len(userEmails))]
	userID, err := f.refStore.GetUserID(randomEmail)
	if err != nil {
		return nil, fmt.Errorf("error getting user ID: %w", err)
	}

	return &models.BlogPostModel{
		ID:        uuid.New().String(),
		Title:     faker.Sentence(),
		Content:   faker.Paragraph(),
		AuthorID:  userID,
		ImageId:   f.refStore.GetRandomBlogPostImageID(),
		CreatedAt: time.Now(),
	}, nil
}

// CreateMany creates multiple blog posts
func (f *BlogPostFactory) CreateMany(count int) ([]*models.BlogPostModel, error) {
	posts := make([]*models.BlogPostModel, 0, count)

	for i := 0; i < count; i++ {
		post, err := f.Create()
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}
