using System.Collections.Immutable;
using Blog.API.Mappers;
using BlogModule.Application.CommandHandlers.BlogPost.CreateBlogPost;
using BlogModule.Application.CommandHandlers.BlogPost.UpdateBlogPost;
using BlogModule.Application.CommandHandlers.BlogPostComment.CreateBlogPostComment;
using BlogModule.Application.CommandHandlers.BlogPostLike.CreateLikeForBlogPost;
using BlogModule.Application.Contract.BlogPostComments.CreateBlogPostComment;
using BlogModule.Application.Contract.BlogPostComments.GetBlogPostComments;
using BlogModule.Application.Contract.BlogPostLikes.CreateLikeForBlogPost;
using BlogModule.Application.Contract.BlogPosts.CreateBlogPost;
using BlogModule.Application.Contract.BlogPosts.GetBlogPosts;
using BlogModule.Application.Contract.BlogPosts.UpdateBlogPost;
using BlogModule.Application.Contract.Tags.GetTags;
using BlogModule.Application.QueryHandlers.BlogPost.GetBlogPosts;
using BlogModule.Application.QueryHandlers.BlogPostComment.GetBlogPostComments;
using BlogModule.Application.QueryHandlers.Tag.GetTags;

namespace BlogModule.API.Mappers;

public class BlogPostApiMapper : IBlogPostApiMapper
{
    public GetBlogPostsQuery MapToMessage(GetBlogPostsRequest request) => new()
    {
        Page = request.Page,
        PageSize = request.PageSize,
        OrderBy = request.OrderBy,
        OrderDirection = request.OrderDirection,
        Ids = request.Ids is null ? null : ImmutableArray.Create(request.Ids),
        PublishedDates = request.PublishedDates is null ? null : ImmutableArray.Create(request.PublishedDates),
        Authors = request.Authors is null ? null : ImmutableArray.Create(request.Authors)
    };

    public CreateBlogPostCommand MapToMessage(CreateBlogPostRequest request) => new()
    {
        Heading = request.Heading,
        PageTitle = request.PageTitle,
        Content = request.Content,
        ShortDescription = request.ShortDescription,
        FeaturedImageUrl = request.FeaturedImageUrl,
        UrlHandle = request.UrlHandle,
        PublishedDate = request.PublishedDate,
        Author = request.Author,
        Visible = request.Visible,
        Tags = request.Tags is null ? null : ImmutableArray.Create(request.Tags)
    };

    public UpdateBlogPostCommand MapToMessage(UpdateBlogPostRequest request) => new()
    {
        Heading = request.Heading,
        PageTitle = request.PageTitle,
        Content = request.Content,
        ShortDescription = request.ShortDescription,
        FeaturedImageUrl = request.FeaturedImageUrl,
        UrlHandle = request.UrlHandle,
        PublishedDate = request.PublishedDate,
        Author = request.Author,
        Visible = request.Visible,
        Tags = request.Tags is null ? null : ImmutableArray.Create(request.Tags)
    };

    public GetBlogPostCommentsQuery MapToMessage(GetBlogPostCommentsRequest request) => new()
    {
        Page = request.Page,
        PageSize = request.PageSize,
        OrderBy = request.OrderBy,
        OrderDirection = request.OrderDirection,
        BlogPostId = request.BlogPostId,
        Ids = request.Ids is null ? null : ImmutableArray.Create(request.Ids),
        DateTimes = request.DateTimes is null ? null : ImmutableArray.Create(request.DateTimes),
        UserIds = request.UserIds is null ? null : ImmutableArray.Create(request.UserIds)
    };

    public CreateBlogPostCommentCommand MapToMessage(CreateBlogPostCommentRequest request) => new()
    {
        Description = request.Description,
        BlogPostId = request.BlogPostId,
        UserId = request.UserId
    };

    public CreateLikeForBlogPostCommand MapToMessage(CreateLikeForBlogPostRequest request) => new()
    {
        BlogPostId = request.BlogPostId,
        UserId = request.UserId
    };

    public GetTagsQuery MapToMessage(GetTagsRequest request) => new()
    {
        OrderBy = request.OrderBy,
        OrderDirection = request.OrderDirection,
        BlogPostId = request.BlogPostId
    };
}