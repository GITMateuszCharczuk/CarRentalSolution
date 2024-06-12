using System.Text.Json.Serialization;
using BlogModule.Domain.Enums;
using BlogModule.Domain.Models.Ids;
using Results.Contract;

namespace BlogModule.Application.Contract.BlogPostComments.GetBlogPostComments;

public class GetBlogPostCommentsRequest : IPageableRequest, ISortable<BlogPostCommentSortColumnEnum?>
{
    public int? Page { get; init; }
    public int? PageSize { get; init; }
    [JsonConverter(typeof(JsonStringEnumConverter))]
    public BlogPostCommentSortColumnEnum? OrderBy { get; init; }
    [JsonConverter(typeof(JsonStringEnumConverter))]
    public SortOrderEnum? OrderDirection { get; init; }
    public BlogPostId? BlogPostId { get; init; }
    public BlogPostCommentId[]? Ids{ get; init; }
    public DateTime[]? DateTimes{ get; init; }
    public Guid[]? UserIds{ get; init; }
}