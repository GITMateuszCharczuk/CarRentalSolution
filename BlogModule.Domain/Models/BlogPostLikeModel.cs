﻿using BlogModule.Domain.Models.Ids;

namespace BlogModule.Domain.Models;

public record BlogPostLikeModel
{
    public BlogPostLikeId Id { get; init; }
    public Guid BlogPostId { get; set; }
    public Guid UserId { get; set; }
}