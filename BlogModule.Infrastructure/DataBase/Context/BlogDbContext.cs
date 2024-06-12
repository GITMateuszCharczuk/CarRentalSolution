using BlogModule.Infrastructure.DataBase.Entities;
using CarRental.Web.Models.Domain.Blog;
using Microsoft.EntityFrameworkCore;

namespace BlogModule.Infrastructure.DataBase.Context;

public class BlogDbContext : DbContext
{
    public DbSet<BlogPostEntity> BlogPosts { get; set; } = null!;
    public DbSet<TagEntity> Tags { get; set; } = null!;
    public DbSet<BlogPostLikeEntity> BlogPostLikes { get; set; } = null!;
    public DbSet<BlogPostCommentEntity> BlogPostComments { get; set; } = null!;
    
    public BlogDbContext()
    {
    }
    

    public BlogDbContext(DbContextOptions<BlogDbContext> options) : base(options)
    {
    }
    
    protected override void OnModelCreating(ModelBuilder modelBuilder)
    {
        modelBuilder.ApplyConfigurationsFromAssembly(typeof(BlogDbContext).Assembly);
    }
}