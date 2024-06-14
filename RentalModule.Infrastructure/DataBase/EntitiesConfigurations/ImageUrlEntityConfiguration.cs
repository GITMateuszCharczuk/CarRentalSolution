using Microsoft.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore.Metadata.Builders;
using RentalModule.Domain.Models;
using RentalModule.Domain.Models.Ids;
using RentalModule.Infrastructure.DataBase.Entities;

namespace RentalModule.Infrastructure.DataBase.EntitiesConfigurations;

public class ImageUrlEntityConfiguration : IEntityTypeConfiguration<ImageUrlEntity>
{
    public void Configure(EntityTypeBuilder<ImageUrlEntity> builder)
    {
        builder.HasKey(entity => entity.Id);

        builder.Property(entity => entity.Id)
            .HasConversion(
                id => (Guid)id,
                value => new ImageUrlId(value))
            .ValueGeneratedOnAdd();

        builder.Property(entity => entity.CarOfferId)
            .IsRequired();

        builder.Property(entity => entity.Url)
            .IsRequired()
            .HasMaxLength(500);
    }
}