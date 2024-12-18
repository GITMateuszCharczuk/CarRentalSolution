package contract

import (
	models "blog-api/Domain/models/domestic"
	responses "blog-api/Domain/responses"
)

type GetBlogPostResponse struct {
	responses.BaseResponse
	BlogPost models.BlogPostResponseModel `json:"blog_post" swaggertype:"object"`
}

type GetBlogPostResponse200 struct {
	Success  bool   `json:"success" example:"true" swaggertype:"boolean"`
	Message  string `json:"message" example:"Blog post retrieved successfully" swaggertype:"string"`
	BlogPost struct {
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
	} `json:"blog_post"`
}

type GetBlogPostResponse400 struct {
	Success  bool                         `json:"success" example:"false" swaggertype:"boolean"`
	Message  string                       `json:"message" example:"Invalid request parameters" swaggertype:"string"`
	BlogPost models.BlogPostResponseModel `json:"blog_post" swaggertype:"object"`
}

type GetBlogPostResponse401 struct {
	Success  bool                         `json:"success" example:"false" swaggertype:"boolean"`
	Message  string                       `json:"message" example:"Unauthorized" swaggertype:"string"`
	BlogPost models.BlogPostResponseModel `json:"blog_post" swaggertype:"object"`
}

type GetBlogPostResponse404 struct {
	Success  bool                         `json:"success" example:"false" swaggertype:"boolean"`
	Message  string                       `json:"message" example:"Blog post not found" swaggertype:"string"`
	BlogPost models.BlogPostResponseModel `json:"blog_post" swaggertype:"object"`
}

type GetBlogPostResponse500 struct {
	Success  bool                         `json:"success" example:"false" swaggertype:"boolean"`
	Message  string                       `json:"message" example:"Internal server error while retrieving blog post" swaggertype:"string"`
	BlogPost models.BlogPostResponseModel `json:"blog_post" swaggertype:"object"`
}
