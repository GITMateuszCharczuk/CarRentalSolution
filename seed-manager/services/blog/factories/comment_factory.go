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

type CommentFactory struct {
	refStore *reference_store.InMemoryStore
}

func NewCommentFactory(refStore *reference_store.InMemoryStore) *CommentFactory {
	return &CommentFactory{
		refStore: refStore,
	}
}

func (f *CommentFactory) Create(postID string) (*models.BlogCommentModel, error) {
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

	return &models.BlogCommentModel{
		ID:        uuid.New().String(),
		Content:   faker.Paragraph(),
		AuthorID:  userID,
		PostID:    postID,
		CreatedAt: time.Now(),
	}, nil
}
