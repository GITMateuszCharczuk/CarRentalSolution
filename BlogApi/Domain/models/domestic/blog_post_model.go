package models

type BlogPostRequestModel struct {
	Heading          string   `json:"heading" example:"Blog Post Title" swaggertype:"string"`
	PageTitle        string   `json:"pageTitle" example:"Page Title" swaggertype:"string"`
	Content          string   `json:"content" example:"Blog post content..." swaggertype:"string"`
	ShortDescription string   `json:"shortDescription" example:"Short description" swaggertype:"string"`
	FeaturedImageUrl string   `json:"featuredImageUrl" example:"https://example.com/image.jpg" swaggertype:"string"`
	UrlHandle        string   `json:"urlHandle" example:"blog-post-title" swaggertype:"string"`
	AuthorId         string   `json:"authorId" example:"123e4567-e89b-12d3-a456-426614174000" swaggertype:"string"`
	AuthorName       string   `json:"author" example:"John Doe" swaggertype:"string"`
	Visible          bool     `json:"visible" example:"true" swaggertype:"boolean"`
	CreatedAt        string   `json:"createdAt" example:"2023-12-12T10:00:00Z" swaggertype:"string"`
	Tags             []string `json:"tags" example:"[\"tech\",\"programming\"]" swaggertype:"array,string"`
}

type BlogPostResponseModel struct {
	Id               string `json:"id" example:"123e4567-e89b-12d3-a456-426614174000" swaggertype:"string"`
	Heading          string `json:"heading" example:"Blog Post Title" swaggertype:"string"`
	PageTitle        string `json:"pageTitle" example:"Page Title" swaggertype:"string"`
	Content          string `json:"content" example:"Blog post content..." swaggertype:"string"`
	ShortDescription string `json:"shortDescription" example:"Short description" swaggertype:"string"`
	FeaturedImageUrl string `json:"featuredImageUrl" example:"https://example.com/image.jpg" swaggertype:"string"`
	UrlHandle        string `json:"urlHandle" example:"blog-post-title" swaggertype:"string"`
	PublishedDate    string `json:"publishedDate" example:"2023-12-12" swaggertype:"string"`
	AuthorName       string `json:"author" example:"John Doe" swaggertype:"string"`
	CreatedAt        string `json:"createdAt" example:"2023-12-12T10:00:00Z" swaggertype:"string"`
}
