using BlogModule.Domain.Models;
using BlogModule.Domain.Models.Ids;
using BlogModule.Infrastructure.DataBase.Entities;
using Microsoft.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore.Metadata.Builders;

namespace BlogModule.Infrastructure.DataBase.EntitiesConfigurations
{
    public class BlogPostEntityConfiguration : IEntityTypeConfiguration<BlogPostEntity>
    {
        public void Configure(EntityTypeBuilder<BlogPostEntity> builder)
        {
            builder.HasKey(entity => entity.Id);
            
            builder.Property(entity => entity.Id)
                .HasConversion(
                    id => (Guid)id,
                    value => value)
                .ValueGeneratedOnAdd();

            builder.Property(entity => entity.Heading)
                .IsRequired()
                .HasMaxLength(200);

            builder.Property(entity => entity.PageTitle)
                .IsRequired()
                .HasMaxLength(100);

            builder.Property(entity => entity.Content)
                .IsRequired();

            builder.Property(entity => entity.ShortDescription)
                .IsRequired()
                .HasMaxLength(500);

            builder.Property(entity => entity.FeaturedImageUrl)
                .HasMaxLength(200);

            builder.Property(entity => entity.UrlHandle)
                .IsRequired()
                .HasMaxLength(100);

            builder.Property(entity => entity.PublishedDate)
                .IsRequired();

            builder.Property(entity => entity.Author)
                .IsRequired()
                .HasMaxLength(100);

            builder.Property(entity => entity.Visible)
                .IsRequired();

            builder.HasMany(entity => entity.Tags)
                .WithOne()
                .HasForeignKey(tag => tag.BlogPostId);

            builder.HasMany(entity => entity.Likes)
                .WithOne()
                .HasForeignKey(like => like.BlogPostId);

            builder.HasMany(entity => entity.Comments)
                .WithOne()
                .HasForeignKey(comment => comment.BlogPostId);
        }
    }


}
