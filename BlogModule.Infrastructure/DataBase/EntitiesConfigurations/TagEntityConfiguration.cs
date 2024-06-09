using BlogModule.Domain.Models;
using BlogModule.Domain.Models.Ids;
using CarRental.Web.Models.Domain.Blog;
using Microsoft.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore.Metadata.Builders;

namespace BlogModule.Infrastructure.DataBase.EntitiesConfigurations
{
    public class TagEntityConfiguration : IEntityTypeConfiguration<TagModel>
    {
        public void Configure(EntityTypeBuilder<TagModel> builder)
        {
            builder.HasKey(entity => entity.Id);

            builder.Property(entity => entity.Id)
                .HasConversion(
                    id => (Guid)id,
                    value => new TagId(value))
                .ValueGeneratedOnAdd();

            builder.Property(entity => entity.Name)
                .IsRequired()
                .HasMaxLength(100);

            builder.Property(entity => entity.BlogPostId)
                .IsRequired();
        }
    }
}