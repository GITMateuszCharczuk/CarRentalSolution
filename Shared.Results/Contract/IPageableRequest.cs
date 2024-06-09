namespace Results.Contract;

public interface IPageableRequest
{
    public int? Page { get; init; }
    public int? PageSize { get; init; }
}