using BlogModule.Domain.Models;
using BlogModule.Domain.Models.Ids;
using BlogModule.Infrastructure.DataBase.Entities;
using Microsoft.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore.Metadata.Builders;

namespace BlogModule.Infrastructure.DataBase.EntitiesConfigurations
{
    public class TagEntityConfiguration : IEntityTypeConfiguration<TagEntity>
    {
        public void Configure(EntityTypeBuilder<TagEntity> builder)
        {
            builder.HasKey(entity => entity.Id);

            // builder.Property(entity => entity.Id)
            //     .HasConversion(
            //         id => id.Value,
            //         value => new TagId(value))
            //     .ValueGeneratedOnAdd();
            
            builder.Property(entity => entity.Id)
                .HasConversion(
                    id => (Guid)id,
                    value => value)
                .ValueGeneratedOnAdd();

            builder.Property(entity => entity.Name)
                .IsRequired()
                .HasMaxLength(100);

            builder.Property(entity => entity.BlogPostId)
                .IsRequired();
        }
    }
}