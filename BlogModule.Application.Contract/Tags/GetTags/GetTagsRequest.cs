using System.Text.Json.Serialization;
using BlogModule.Domain.Enums;
using BlogModule.Domain.Models.Ids;
using Results.Contract;

namespace BlogModule.Application.Contract.Tags.GetTags;

public class GetTagsRequest : ISortable<TagSortColumnEnum?>
{
    [JsonConverter(typeof(JsonStringEnumConverter))]
    public TagSortColumnEnum? OrderBy { get; init; } = TagSortColumnEnum.Name;
    [JsonConverter(typeof(JsonStringEnumConverter))]
    public SortOrderEnum? OrderDirection { get; init; } 
    public BlogPostId? BlogPostId { get; init; } 
}