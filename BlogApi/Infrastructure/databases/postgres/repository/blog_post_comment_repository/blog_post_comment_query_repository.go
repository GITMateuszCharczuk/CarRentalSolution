package repository

import (
	models "blog-api/Domain/models/domestic"
	"blog-api/Domain/pagination"
	repository_interfaces "blog-api/Domain/repository_interfaces/blog_post_comment_repository"
	"blog-api/Domain/sorting"
	postgres_db "blog-api/Infrastructure/databases/postgres/config"
	"blog-api/Infrastructure/databases/postgres/entities"
	mappers "blog-api/Infrastructure/databases/postgres/mappers/base"
	base "blog-api/Infrastructure/databases/postgres/repository/base"
	"blog-api/Infrastructure/databases/postgres/repository/base/helpers"
	unit_of_work "blog-api/Infrastructure/databases/postgres/repository/base/unit_of_work"
	"log"
)

type BlogPostCommentQueryRepositoryImpl struct {
	*base.QueryRepository[entities.BlogPostCommentEntity, string, models.BlogPostCommentModel]
	mapper mappers.PersistenceMapper[entities.BlogPostCommentEntity, models.BlogPostCommentModel]
}

func NewBlogPostCommentQueryRepositoryImpl(
	postgresDatabase *postgres_db.PostgresDatabase,
	mapper mappers.PersistenceMapper[entities.BlogPostCommentEntity, models.BlogPostCommentModel],
	uow unit_of_work.UnitOfWork,
) repository_interfaces.BlogPostCommentQueryRepository {
	return &BlogPostCommentQueryRepositoryImpl{
		QueryRepository: base.NewQueryRepository[entities.BlogPostCommentEntity, string, models.BlogPostCommentModel](postgresDatabase.DB, mapper, uow),
		mapper:          mapper,
	}
}

func (r *BlogPostCommentQueryRepositoryImpl) GetCommentByID(id string) (*models.BlogPostCommentModel, error) {
	queryRecord := helpers.NewQueryRecord[entities.BlogPostCommentEntity]("id", id)
	return r.GetFirstByQueryRecord(queryRecord)
}

func (r *BlogPostCommentQueryRepositoryImpl) GetCommentAuthorId(id string) (*string, error) {
	var blogPostComment entities.BlogPostCommentEntity
	if err := r.ConstructBaseQuery().Where("id = ?", id).First(&blogPostComment).Error; err != nil {
		return nil, err
	}
	res := blogPostComment.UserID.String()
	return &res, nil
}

func (r *BlogPostCommentQueryRepositoryImpl) GetComments(
	blogPostIDs []string,
	userIDs []string,
	dateTimeFrom string,
	dateTimeTo string,
	pagination *pagination.Pagination,
	sorting *sorting.Sortable,
) (*pagination.PaginatedResult[models.BlogPostCommentModel], error) {
	queryRecords := []helpers.QueryRecord[entities.BlogPostCommentEntity]{
		helpers.NewQueryRecord[entities.BlogPostCommentEntity]("blog_post_id", blogPostIDs),
		helpers.NewQueryRecord[entities.BlogPostCommentEntity]("user_id", userIDs),
	}
	log.Println(blogPostIDs)
	query := r.ConstructBaseQuery()
	query = r.ApplyWhereConditions(query, queryRecords...)

	if dateTimeFrom != "" && dateTimeTo != "" {
		query = query.Where("created_at BETWEEN ? AND ?", dateTimeFrom, dateTimeTo)
	} else if dateTimeFrom != "" {
		query = query.Where("created_at >= ?", dateTimeFrom)
	} else if dateTimeTo != "" {
		query = query.Where("created_at <= ?", dateTimeTo)
	}

	return r.ExecutePaginatedQuery(query, pagination, sorting)
}

func (r *BlogPostCommentQueryRepositoryImpl) GetCommentsCount(blogPostId string) (int, error) {
	var count int64
	if err := r.ConstructBaseQuery().Where("blog_post_id = ?", blogPostId).Count(&count).Error; err != nil {
		return 0, err
	}
	return int(count), nil
}
