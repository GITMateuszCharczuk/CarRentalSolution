using BlogModule.Domain.Models;
using BlogModule.Domain.Models.Ids;
using Microsoft.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore.Metadata.Builders;

namespace BlogModule.Infrastructure.DataBase.EntitiesConfigurations
{
    public class BlogPostCommentEntityConfiguration : IEntityTypeConfiguration<BlogPostCommentModel>
    {
        public void Configure(EntityTypeBuilder<BlogPostCommentModel> builder)
        {
            builder.HasKey(entity => entity.Id);

            builder.Property(entity => entity.Id)
                .HasConversion(
                    id => (Guid)id,
                    value => new BlogPostCommentId(value))
                .ValueGeneratedOnAdd();

            builder.Property(entity => entity.Description)
                .IsRequired()
                .HasMaxLength(500);

            builder.Property(entity => entity.BlogPostId)
                .IsRequired();

            builder.Property(entity => entity.UserId)
                .IsRequired();

            builder.Property(entity => entity.DateAdded)
                .IsRequired();
        }
    }
}