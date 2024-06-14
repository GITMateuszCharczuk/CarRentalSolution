using Microsoft.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore.Metadata.Builders;
using RentalModule.Domain.Models;
using RentalModule.Domain.Models.Ids;
using RentalModule.Infrastructure.DataBase.Entities;

namespace RentalModule.Infrastructure.DataBase.EntitiesConfigurations;

public class CarTagEntityConfiguration : IEntityTypeConfiguration<CarTagEntity>
{
    public void Configure(EntityTypeBuilder<CarTagEntity> builder)
    {
        builder.HasKey(entity => entity.Id);

        builder.Property(entity => entity.Id)
            .HasConversion(
                id => (Guid)id,
                value => new CarTagId(value))
            .ValueGeneratedOnAdd();

        builder.Property(entity => entity.Name)
            .IsRequired()
            .HasMaxLength(100);

        builder.Property(entity => entity.CarOfferId)
            .IsRequired();
    }
}