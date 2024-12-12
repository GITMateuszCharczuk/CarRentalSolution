package contract

import (
	responses "identity-api/Domain/responses"
)

type BlogPostModel struct {
	Id               string   `json:"id" example:"123e4567-e89b-12d3-a456-426614174000" swaggertype:"string"`
	Heading          string   `json:"heading" example:"Blog Post Title" swaggertype:"string"`
	PageTitle        string   `json:"pageTitle" example:"Page Title" swaggertype:"string"`
	Content          string   `json:"content" example:"Blog post content..." swaggertype:"string"`
	ShortDescription string   `json:"shortDescription" example:"Short description" swaggertype:"string"`
	FeaturedImageUrl string   `json:"featuredImageUrl" example:"https://example.com/image.jpg" swaggertype:"string"`
	UrlHandle        string   `json:"urlHandle" example:"blog-post-title" swaggertype:"string"`
	PublishedDate    string   `json:"publishedDate" example:"2023-12-12" swaggertype:"string"`
	Author           string   `json:"author" example:"John Doe" swaggertype:"string"`
	Visible          bool     `json:"visible" example:"true" swaggertype:"boolean"`
	Tags             []string `json:"tags" example:"[\"tech\",\"programming\"]" swaggertype:"array,string"`
}

type GetBlogPostsResponse struct {
	responses.BaseResponse
	Page           *int                `json:"page" example:"1" swaggertype:"integer"`
	PageSize       *int                `json:"pageSize" example:"10" swaggertype:"integer"`
	TotalCount     int                 `json:"totalCount" example:"100" swaggertype:"integer"`
	OrderBy        *BlogPostSortColumn `json:"orderBy" example:"publishedDate" swaggertype:"string"`
	OrderDirection *SortOrder          `json:"orderDirection" example:"desc" swaggertype:"string"`
	Items          []BlogPostModel     `json:"items" swaggertype:"array,object"`
}

type GetBlogPostsResponse200 struct {
	Success        bool                `json:"success" example:"true" swaggertype:"boolean"`
	Message        string              `json:"message" example:"Blog posts retrieved successfully" swaggertype:"string"`
	Page           *int                `json:"page" example:"1" swaggertype:"integer"`
	PageSize       *int                `json:"pageSize" example:"10" swaggertype:"integer"`
	TotalCount     int                 `json:"totalCount" example:"100" swaggertype:"integer"`
	OrderBy        *BlogPostSortColumn `json:"orderBy" example:"publishedDate" swaggertype:"string"`
	OrderDirection *SortOrder          `json:"orderDirection" example:"desc" swaggertype:"string"`
	Items          []BlogPostModel     `json:"items" swaggertype:"array,object"`
}

type GetBlogPostsResponse400 struct {
	Success bool   `json:"success" example:"false" swaggertype:"boolean"`
	Message string `json:"message" example:"Invalid request parameters" swaggertype:"string"`
}

type GetBlogPostsResponse500 struct {
	Success bool   `json:"success" example:"false" swaggertype:"boolean"`
	Message string `json:"message" example:"Internal server error while retrieving blog posts" swaggertype:"string"`
}
