using BlogModule.Domain.Enums;
using RentalModule.Application.Contract.Tags.GetTags;
using RentalModule.Domain.Models.Ids;
using Results.Contract;
using Results.Domain;
using Shared.CQRS;
using Shared.CQRS.QueryHandlers;

namespace RentalModule.Application.QueryHandlers.Tag.GetTags;

public record GetTagsQuery : ISortable<CarTagSortColumnEnum?>, IQuery<HandlerResult<GetTagsResponse, IErrorResult>>

{
    public CarTagSortColumnEnum? OrderBy { get; init; }
    public SortOrderEnum? OrderDirection { get; init; }
    public CarOfferId? CarOfferId { get; init; } 
}