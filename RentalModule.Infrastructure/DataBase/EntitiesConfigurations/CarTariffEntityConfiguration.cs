using Microsoft.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore.Metadata.Builders;
using RentalModule.Domain.Models;
using RentalModule.Domain.Models.Ids;
using RentalModule.Infrastructure.DataBase.Entities;

namespace RentalModule.Infrastructure.DataBase.EntitiesConfigurations;

public class CarTariffEntityConfiguration : IEntityTypeConfiguration<CarTariffEntity>
{
    public void Configure(EntityTypeBuilder<CarTariffEntity> builder)
    {
        builder.HasKey(entity => entity.Id);

        builder.Property(entity => entity.Id)
            .HasConversion(
                id => (Guid)id,
                value => new CarTariffId(value))
            .ValueGeneratedOnAdd();

        builder.Property(entity => entity.CarOfferId)
            .IsRequired();

        builder.Property(entity => entity.OneNormalDayPrice)
            .IsRequired();

        builder.Property(entity => entity.OneWeekendDayPrice)
            .IsRequired();

        builder.Property(entity => entity.FullWeekendPrice)
            .IsRequired();

        builder.Property(entity => entity.OneWeekPrice)
            .IsRequired();

        builder.Property(entity => entity.OneMonthPrice)
            .IsRequired();
    }
}