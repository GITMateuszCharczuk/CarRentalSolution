package contract

import (
	models "identity-api/Domain/models/domestic"
	"identity-api/Domain/pagination"
	responses "identity-api/Domain/responses"
)

type GetBlogPostsResponse struct {
	responses.BaseResponse
	pagination.PaginatedResult[models.BlogPostResponseModel] `json:",inline"`
}

type GetBlogPostsResponse200 struct {
	StatusCode int    `json:"status_code" example:"200"`
	Message    string `json:"message" example:"Users retrieved successfully"`
	Data       struct {
		Items []struct {
			Id               string   `json:"id" example:"123e4567-e89b-12d3-a456-426614174000" swaggertype:"string"`
			Heading          string   `json:"heading" example:"Blog Post Title" swaggertype:"string"`
			PageTitle        string   `json:"pageTitle" example:"Page Title" swaggertype:"string"`
			Content          string   `json:"content" example:"Blog post content..." swaggertype:"string"`
			ShortDescription string   `json:"shortDescription" example:"Short description" swaggertype:"string"`
			FeaturedImageUrl string   `json:"featuredImageUrl" example:"https://example.com/image.jpg" swaggertype:"string"`
			UrlHandle        string   `json:"urlHandle" example:"blog-post-title" swaggertype:"string"`
			PublishedDate    string   `json:"publishedDate" example:"2023-12-12" swaggertype:"string"`
			AuthorName       string   `json:"author" example:"John Doe" swaggertype:"string"`
			Visible          bool     `json:"visible" example:"true" swaggertype:"boolean"`
			Tags             []string `json:"tags" example:"[\"tech\",\"programming\"]" swaggertype:"array,string"`
		} `json:"items"`
		TotalItems  int `json:"total_items" example:"100"`
		CurrentPage int `json:"current_page" example:"1"`
		PageSize    int `json:"page_size" example:"10"`
		TotalPages  int `json:"total_pages" example:"10"`
	} `json:"data"`
}

type GetBlogPostsResponse400 struct {
	Success bool               `json:"success" example:"false" swaggertype:"boolean"`
	Message string             `json:"message" example:"Invalid request parameters" swaggertype:"string"`
	Data    []dummyArrayObject `json:"data" swaggertype:"array,object"`
}

type GetBlogPostsResponse500 struct {
	Success bool               `json:"success" example:"false" swaggertype:"boolean"`
	Message string             `json:"message" example:"Internal server error while retrieving blog posts" swaggertype:"string"`
	Data    []dummyArrayObject `json:"data" swaggertype:"array,object"`
}

type dummyArrayObject struct{}
