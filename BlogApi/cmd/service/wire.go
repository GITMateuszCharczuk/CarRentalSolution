// main/wire.go

//go:build wireinject
// +build wireinject

package main

import (
	"identity-api/API/controllers"
	"identity-api/API/server"
	validators "identity-api/API/validators"
	comment_repository_interfaces "identity-api/Domain/repository_interfaces/blog_post_comment_repository"
	like_repository_interfaces "identity-api/Domain/repository_interfaces/blog_post_like_repository"
	blog_repository_interfaces "identity-api/Domain/repository_interfaces/blog_post_repository"
	tag_repository_interfaces "identity-api/Domain/repository_interfaces/blog_post_tag_repository"
	service_interfaces "identity-api/Domain/service_interfaces"
	config "identity-api/Infrastructure/config"
	data_fetcher "identity-api/Infrastructure/data_fetcher"
	postgres_db "identity-api/Infrastructure/databases/postgres/config"
	blog_mappers "identity-api/Infrastructure/databases/postgres/mappers"
	comment_repository "identity-api/Infrastructure/databases/postgres/repository/blog_post_comment_repository"
	like_repository "identity-api/Infrastructure/databases/postgres/repository/blog_post_like_repository"
	blog_repository "identity-api/Infrastructure/databases/postgres/repository/blog_post_repository"
	tag_repository "identity-api/Infrastructure/databases/postgres/repository/blog_post_tag_repository"

	"github.com/google/wire"
)

type InfrastructureComponents struct {
	Config             *config.Config
	BlogQueryRepo      blog_repository_interfaces.BlogPostQueryRepository
	BlogCommandRepo    blog_repository_interfaces.BlogPostCommandRepository
	CommentQueryRepo   comment_repository_interfaces.BlogPostCommentQueryRepository
	CommentCommandRepo comment_repository_interfaces.BlogPostCommentCommandRepository
	LikeQueryRepo      like_repository_interfaces.BlogPostLikeQueryRepository
	LikeCommandRepo    like_repository_interfaces.BlogPostLikeCommandRepository
	TagQueryRepo       tag_repository_interfaces.BlogPostTagQueryRepository
	DataFetcher        service_interfaces.DataFetcher
}

func InitializeInfrastructureComponents() (*InfrastructureComponents, error) {
	wire.Build(
		// Config
		config.WireSet,
		// Database
		postgres_db.WireSet,
		// Repositories
		blog_repository.WireSet,
		comment_repository.WireSet,
		like_repository.WireSet,
		tag_repository.WireSet,
		// Mappers
		blog_mappers.WireSet,
		// Services
		data_fetcher.WireSet,
		wire.Struct(new(InfrastructureComponents), "*"),
	)
	return &InfrastructureComponents{}, nil
}

func InitializeApi(
	blogQueryRepo blog_repository_interfaces.BlogPostQueryRepository,
	blogCommandRepo blog_repository_interfaces.BlogPostCommandRepository,
	commentQueryRepo comment_repository_interfaces.BlogPostCommentQueryRepository,
	commentCommandRepo comment_repository_interfaces.BlogPostCommentCommandRepository,
	likeQueryRepo like_repository_interfaces.BlogPostLikeQueryRepository,
	likeCommandRepo like_repository_interfaces.BlogPostLikeCommandRepository,
	tagQueryRepo tag_repository_interfaces.BlogPostTagQueryRepository,
	dataFetcher service_interfaces.DataFetcher,
	config *config.Config,
) (*server.Server, error) {
	wire.Build(
		validators.WireSet,
		controllers.WireSet,
		server.WireSet,
	)
	return &server.Server{}, nil
}
