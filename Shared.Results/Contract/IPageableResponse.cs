namespace Results.Contract;

public interface IPageableResponse
{
    public int? Page { get; init; }
    public int? PageSize { get; init; }
    public int TotalCount { get; init; }
}