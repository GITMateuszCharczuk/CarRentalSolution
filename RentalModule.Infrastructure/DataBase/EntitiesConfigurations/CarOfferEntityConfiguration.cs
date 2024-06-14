using RentalModule.Domain.Models.Ids;
using RentalModule.Infrastructure.DataBase.Entities;

namespace RentalModule.Infrastructure.DataBase.EntitiesConfigurations;

using Microsoft.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore.Metadata.Builders;
using RentalModule.Domain.Models;

public class CarOfferEntityConfiguration : IEntityTypeConfiguration<CarOfferEntity>
{
    public void Configure(EntityTypeBuilder<CarOfferEntity> builder)
    {
        builder.HasKey(entity => entity.Id);
        
        builder.Property(entity => entity.Id)
            .HasConversion(
                id => (Guid)id,
                value => new CarOfferId(value))
            .ValueGeneratedOnAdd();

        builder.Property(entity => entity.Heading)
            .IsRequired()
            .HasMaxLength(200);

        builder.Property(entity => entity.ShortDescription)
            .IsRequired()
            .HasMaxLength(500);

        builder.Property(entity => entity.FeaturedImageUrl)
            .HasMaxLength(200);

        builder.Property(entity => entity.UrlHandle)
            .IsRequired()
            .HasMaxLength(100);

        builder.Property(entity => entity.Horsepower)
            .IsRequired()
            .HasMaxLength(50);

        builder.Property(entity => entity.YearOfProduction)
            .IsRequired();

        builder.Property(entity => entity.EngineDetails)
            .IsRequired()
            .HasMaxLength(100);

        builder.Property(entity => entity.DriveDetails)
            .IsRequired()
            .HasMaxLength(100);

        builder.Property(entity => entity.GearboxDetails)
            .IsRequired()
            .HasMaxLength(100);

        builder.Property(entity => entity.CarDeliverylocation)
            .IsRequired()
            .HasMaxLength(200);

        builder.Property(entity => entity.CarReturnLocation)
            .IsRequired()
            .HasMaxLength(200);

        builder.Property(entity => entity.PublishedDate)
            .IsRequired();

        builder.Property(entity => entity.Visible)
            .IsRequired();

        builder.HasOne(entity => entity.Tarrif)
            .WithOne()
            .HasForeignKey<CarTariffModel>(t => t.CarOfferId);

        builder.HasMany(entity => entity.Tags)
            .WithOne()
            .HasForeignKey(tag => tag.CarOfferId);

        builder.HasMany(entity => entity.ImageUrls)
            .WithOne()
            .HasForeignKey(image => image.CarOfferId);

        builder.HasMany(entity => entity.UnavailableDates)
            .WithOne()
            .HasForeignKey(period => period.CarOfferId);
    }
}
