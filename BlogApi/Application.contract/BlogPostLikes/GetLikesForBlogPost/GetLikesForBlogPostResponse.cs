using System.Collections.Immutable;
using BlogModule.Domain.Models;
using Results.Contract;
using Results.Domain;

namespace BlogModule.Application.Contract.BlogPostLikes.GetLikesForBlogPost;

public class GetLikesForBlogPostResponse : IApiCollectionResult<BlogPostLikeModel>,ISuccessResult
{
    public ImmutableArray<BlogPostLikeModel> Items { get; init; }
}