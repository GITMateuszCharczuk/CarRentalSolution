package blog

import (
	"fmt"
	"log"
	"math/rand"
	"seeder-manager/api"
	"seeder-manager/config"
	"seeder-manager/models"
	"seeder-manager/reference_store"
	"seeder-manager/services/blog/factories"
	"strings"
)

type BlogSeeder struct {
	postFactory    *factories.BlogPostFactory
	commentFactory *factories.CommentFactory
	tagFactory     *factories.TagFactory
	apiClient      *api.APIClient
	cfg            *config.Config
}

func NewBlogSeeder(apiBaseURL string, store *reference_store.InMemoryStore) *BlogSeeder {
	return &BlogSeeder{
		postFactory:    factories.NewBlogPostFactory(store),
		commentFactory: factories.NewCommentFactory(store),
		tagFactory:     factories.NewTagFactory(),
		apiClient:      api.NewAPIClient(apiBaseURL),
		cfg:            config.GetConfig(),
	}
}

type CreateBlogPostRequest struct {
	Heading          string   `json:"heading" binding:"required" example:"Blog Post Title" swaggertype:"string" validate:"required"`
	PageTitle        string   `json:"pageTitle" binding:"required" example:"Page Title" swaggertype:"string" validate:"required"`
	Content          string   `json:"content" binding:"required" example:"Blog post content..." swaggertype:"string" validate:"required"`
	ShortDescription string   `json:"shortDescription" binding:"required" example:"Short description" swaggertype:"string" validate:"required"`
	FeaturedImageUrl string   `json:"featuredImageUrl" example:"https://example.com/image.jpg" swaggertype:"string"`
	UrlHandle        string   `json:"urlHandle" binding:"required" example:"blog-post-title" swaggertype:"string" validate:"required"`
	Visible          bool     `json:"visible" example:"true" swaggertype:"boolean"`
	Tags             []string `json:"tags" example:"[\"tech\",\"programming\"]" swaggertype:"array,string"`
}

type CreateBlogCommentRequest struct {
	Description string `json:"description"`
}

func (s *BlogSeeder) createBlogPost(post *models.BlogPostModel, token string) (string, error) {
	// Create a safe short description
	shortDesc := post.Content
	if len(shortDesc) > 100 {
		shortDesc = shortDesc[:100]
	}

	tags := []string{
		"general", "blog", "technology", "programming", "lifestyle", "travel", "food", "health", "fitness", "education",
		"finance", "business", "marketing", "design", "art", "music", "photography", "gaming", "sports", "news",
		"politics", "history", "science", "nature", "environment", "parenting", "relationships", "self-improvement", "motivation", "inspiration", "reviews",
	}

	pickRandomTags := func(tags []string, count int) []string {
		if count < 3 || count > 5 {
			count = 3
		}
		rand.Shuffle(len(tags), func(i, j int) {
			tags[i], tags[j] = tags[j], tags[i]
		})
		return tags[:count]
	}

	request := CreateBlogPostRequest{
		Heading:          post.Title,
		PageTitle:        post.Title,
		Content:          post.Content,
		ShortDescription: shortDesc,
		FeaturedImageUrl: post.ImageId,
		UrlHandle:        strings.ToLower(strings.ReplaceAll(post.Title, " ", "-")),
		Tags:             pickRandomTags(tags, rand.Intn(3)+3),
		Visible:          true,
	}

	resp, err := s.apiClient.Post("/blog-api/api/posts", request, token)
	if err != nil {
		return "", fmt.Errorf("error creating blog post: %w", err)
	}

	if resp.Id == "" {
		return "", fmt.Errorf("blog post created but no ID returned")
	}

	return resp.Id, nil
}

func (s *BlogSeeder) createComment(comment *models.BlogCommentModel, postID string, token string) error {
	request := CreateBlogCommentRequest{
		Description: comment.Content,
	}

	endpoint := fmt.Sprintf("/blog-api/api/posts/%s/comments", postID)
	_, err := s.apiClient.Post(endpoint, request, token)
	if err != nil {
		return fmt.Errorf("error creating comment: %w", err)
	}

	return nil
}

func (s *BlogSeeder) Seed(store *reference_store.InMemoryStore, token string) error {
	log.Printf("Starting to seed %d blog posts with %d comments each...", s.cfg.SeedCount.BlogPosts, s.cfg.SeedCount.CommentsPerPost)

	// Create blog posts
	posts, err := s.postFactory.CreateMany(s.cfg.SeedCount.BlogPosts)
	if err != nil {
		return err
	}

	// Create posts and their comments
	for _, post := range posts {
		// Create blog post
		postID, err := s.createBlogPost(post, token)
		if err != nil {
			return err
		}

		// Create comments for each post
		for i := 0; i < s.cfg.SeedCount.CommentsPerPost; i++ {
			comment, err := s.commentFactory.Create(postID)
			if err != nil {
				return err
			}

			err = s.createComment(comment, postID, token)
			if err != nil {
				return err
			}
		}
	}

	log.Printf("Completed seeding %d blog posts with %d comments each", s.cfg.SeedCount.BlogPosts, s.cfg.SeedCount.CommentsPerPost)
	return nil
}

func (s *BlogSeeder) Cleanup() error {
	// TODO: Implement cleanup logic if needed
	return nil
}
