namespace Results.Contract;

public interface ISortable<TColumnEnum>
{
    public TColumnEnum? OrderBy  { get; init; }
    public SortOrderEnum? OrderDirection { get; init; }
}