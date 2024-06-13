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

namespace Blog.API.Mappers;

public interface IBlogPostApiMapper
{
    GetBlogPostsQuery MapToMessage(GetBlogPostsRequest request);
    CreateBlogPostCommand MapToMessage(CreateBlogPostRequest request);
    UpdateBlogPostCommand MapToMessage(UpdateBlogPostRequest request);
    GetBlogPostCommentsQuery MapToMessage(GetBlogPostCommentsRequest request);
    CreateBlogPostCommentCommand MapToMessage(CreateBlogPostCommentRequest request);
    CreateLikeForBlogPostCommand MapToMessage(CreateLikeForBlogPostRequest request);
    GetTagsQuery MapToMessage(GetTagsRequest request);
}