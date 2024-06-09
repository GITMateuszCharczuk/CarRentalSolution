using BlogModule.Domain.Models;
using BlogModule.Domain.Models.Ids;
using Microsoft.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore.Metadata.Builders;

namespace BlogModule.Infrastructure.DataBase.EntitiesConfigurations
{
    public class BlogPostLikeEntityConfiguration : IEntityTypeConfiguration<BlogPostLikeModel>
    {
        public void Configure(EntityTypeBuilder<BlogPostLikeModel> builder)
        {
            builder.HasKey(entity => entity.Id);

            builder.Property(entity => entity.Id)
                .HasConversion(
                    id => (Guid)id,
                    value => new BlogPostLikeId(value))
                .ValueGeneratedOnAdd();

            builder.Property(entity => entity.BlogPostId)
                .IsRequired();

            builder.Property(entity => entity.UserId)
                .IsRequired();
        }
    }
}