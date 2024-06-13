using BlogModule.Infrastructure.DataBase.Entities;
using BlogModule.Infrastructure.DataBase.EntitiesConfigurations;
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
        // modelBuilder.ApplyConfiguration(new TagEntityConfiguration());
        // modelBuilder.ApplyConfiguration(new BlogPostEntityConfiguration());
        // modelBuilder.ApplyConfiguration(new BlogPostCommentEntityConfiguration());
        // modelBuilder.ApplyConfiguration(new BlogPostLikeEntityConfiguration());
        modelBuilder.ApplyConfigurationsFromAssembly(typeof(BlogDbContext).Assembly);
        //base.OnModelCreating(modelBuilder);
    }
}