using System.Collections.Immutable;
using System.Text.Json.Serialization;
using BlogModule.Domain.Enums;
using BlogModule.Domain.Models.Ids;
using Results.Contract;

namespace BlogModule.Application.Contract.BlogPosts.GetBlogPosts;

public record GetBlogPostsRequest : IPageableRequest, ISortable<BlogPostSortColumnEnum?>
{
    public int? Page { get; init; }
    
    public int? PageSize { get; init; }
    
    [JsonConverter(typeof(JsonStringEnumConverter))]
    public BlogPostSortColumnEnum? OrderBy { get; init; }
    
    [JsonConverter(typeof(JsonStringEnumConverter))]
    public SortOrderEnum? OrderDirection { get; init; }

    public BlogPostId[]? Ids { get; init; }
    
    public DateTime[]? PublishedDates { get; init; }
    
    public string[]? Authors { get; init; }
    
}