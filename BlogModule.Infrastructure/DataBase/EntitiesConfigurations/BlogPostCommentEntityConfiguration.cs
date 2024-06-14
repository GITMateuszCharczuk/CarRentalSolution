using BlogModule.Domain.Models;
using BlogModule.Domain.Models.Ids;
using BlogModule.Infrastructure.DataBase.Entities;
using Microsoft.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore.Metadata.Builders;

namespace BlogModule.Infrastructure.DataBase.EntitiesConfigurations
{
    public class BlogPostCommentEntityConfiguration : IEntityTypeConfiguration<BlogPostCommentEntity>
    {
        public void Configure(EntityTypeBuilder<BlogPostCommentEntity> builder)
        {
            builder.HasKey(entity => entity.Id);

            // builder.Property(entity => entity.Id)
            //     .HasConversion(
            //         id => id.Value,
            //         value => new CarOfferId(value))
            //     .ValueGeneratedOnAdd();
            
            builder.Property(entity => entity.Id)
                .HasConversion(
                    id => (Guid)id,
                    value => value)
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