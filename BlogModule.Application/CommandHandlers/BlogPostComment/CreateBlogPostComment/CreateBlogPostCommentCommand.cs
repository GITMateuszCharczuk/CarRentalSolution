using System.ComponentModel.DataAnnotations;
using BlogModule.Application.Contract.BlogPostComments.CreateBlogPostComment;
using BlogModule.Application.ValidationAttributes;
using BlogModule.Domain.Models.Ids;
using Results.Domain;
using Shared.CQRS;
using Shared.CQRS.CommandHandlers;

namespace BlogModule.Application.CommandHandlers.BlogPostComment.CreateBlogPostComment;

public class CreateBlogPostCommentCommand : ICommand<HandlerResult<CreateBlogPostCommentResponse, IErrorResult>>
{
    [NoSwearWords]
    [StringLength(200)]
    [Required]
    public string Description { get; set; } = string.Empty;
    public BlogPostId BlogPostId { get; set; }
    public Guid UserId { get; set; }
}