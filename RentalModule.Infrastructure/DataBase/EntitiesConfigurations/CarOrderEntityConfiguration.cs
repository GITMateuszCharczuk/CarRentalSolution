using Microsoft.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore.Metadata.Builders;
using RentalModule.Domain.Models;
using RentalModule.Domain.Models.Ids;
using RentalModule.Infrastructure.DataBase.Entities;

namespace RentalModule.Infrastructure.DataBase.EntitiesConfigurations;

public class CarOrderEntityConfiguration : IEntityTypeConfiguration<CarOrderEntity>
{
    public void Configure(EntityTypeBuilder<CarOrderEntity> builder)
    {
        builder.HasKey(entity => entity.Id);

        builder.Property(entity => entity.Id)
            .HasConversion(
                id => (Guid)id,
                value => new CarOrderId(value))
            .ValueGeneratedOnAdd();

        builder.Property(entity => entity.UserId)
            .IsRequired();

        builder.Property(entity => entity.CarOfferId)
            .HasConversion(
                id => (Guid)id,
                value => new CarOfferId(value))
            .IsRequired();

        builder.Property(entity => entity.StartDate)
            .IsRequired();

        builder.Property(entity => entity.EndDate)
            .IsRequired();

        builder.Property(entity => entity.Notes)
            .HasMaxLength(1000);

        builder.Property(entity => entity.TotalCost)
            .IsRequired();
    }
}