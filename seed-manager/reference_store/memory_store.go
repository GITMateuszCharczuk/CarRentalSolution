package reference_store

import (
	"math/rand"
)

type InMemoryStore struct {
	userIDs          map[string]string // email -> userID
	carFileIDs       []string          // list of car image IDs
	blogPostImageIDs []string          // list of blog post image IDs
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		userIDs:          make(map[string]string),
		carFileIDs:       make([]string, 0),
		blogPostImageIDs: make([]string, 0),
	}
}

func (s *InMemoryStore) StoreUserID(email string, userID string) error {
	s.userIDs[email] = userID
	return nil
}

func (s *InMemoryStore) GetUserID(email string) (string, error) {
	if id, ok := s.userIDs[email]; ok {
		return id, nil
	}
	return "", ErrNotFound
}

func (s *InMemoryStore) GetAllUserEmails() []string {
	emails := make([]string, 0, len(s.userIDs))
	for email := range s.userIDs {
		emails = append(emails, email)
	}
	return emails
}

func (s *InMemoryStore) AddCarFileID(fileID string) {
	s.carFileIDs = append(s.carFileIDs, fileID)
}

func (s *InMemoryStore) AddBlogPostImageID(fileID string) {
	s.blogPostImageIDs = append(s.blogPostImageIDs, fileID)
}

func (s *InMemoryStore) GetRandomCarImageIDs(count int) []string {
	if len(s.carFileIDs) == 0 {
		return []string{}
	}

	if count > len(s.carFileIDs) {
		count = len(s.carFileIDs)
	}

	// Create a copy of the IDs to shuffle
	ids := make([]string, len(s.carFileIDs))
	copy(ids, s.carFileIDs)

	// Fisher-Yates shuffle
	for i := len(ids) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		ids[i], ids[j] = ids[j], ids[i]
	}

	return ids[:count]
}

func (s *InMemoryStore) GetRandomBlogPostImageID() string {
	if len(s.blogPostImageIDs) == 0 {
		return ""
	}
	return s.blogPostImageIDs[rand.Intn(len(s.blogPostImageIDs))]
}

func (s *InMemoryStore) Clear() error {
	s.userIDs = make(map[string]string)
	s.carFileIDs = make([]string, 0)
	s.blogPostImageIDs = make([]string, 0)
	return nil
}
