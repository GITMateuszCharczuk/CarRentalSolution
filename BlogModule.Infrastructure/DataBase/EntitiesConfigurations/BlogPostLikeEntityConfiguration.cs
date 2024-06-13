using BlogModule.Domain.Models;
using BlogModule.Domain.Models.Ids;
using BlogModule.Infrastructure.DataBase.Entities;
using Microsoft.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore.Metadata.Builders;

namespace BlogModule.Infrastructure.DataBase.EntitiesConfigurations
{
    public class BlogPostLikeEntityConfiguration : IEntityTypeConfiguration<BlogPostLikeEntity>
    {
        public void Configure(EntityTypeBuilder<BlogPostLikeEntity> builder)
        {
            builder.HasKey(entity => entity.Id);

            // builder.Property(entity => entity.Id)
            //     .HasConversion(
            //         id => id.Value,
            //         value => new BlogPostLikeId(value))
            //     .ValueGeneratedOnAdd();
            
            builder.Property(entity => entity.Id)
                .HasConversion(
                    id => (Guid)id,
                    value => value)
                .ValueGeneratedOnAdd();

            builder.Property(entity => entity.BlogPostId)
                .IsRequired();

            builder.Property(entity => entity.UserId)
                .IsRequired();
        }
    }
}