using Microsoft.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore.Metadata.Builders;
using RentalModule.Domain.Models.Ids;
using RentalModule.Infrastructure.DataBase.Entities;

namespace RentalModule.Infrastructure.DataBase.EntitiesConfigurations;

public class TimePeriodEntityConfiguration : IEntityTypeConfiguration<TimePeriodEntity>
{
    public void Configure(EntityTypeBuilder<TimePeriodEntity> builder)
    {
        builder.HasKey(entity => entity.Id);

        builder.Property(entity => entity.Id)
            .HasConversion(
                id => (Guid)id,
                value => new TimePeriodId(value))
            .ValueGeneratedOnAdd();

        builder.Property(entity => entity.CarOfferId)
            .IsRequired();

        builder.Property(entity => entity.StartDate)
            .IsRequired();

        builder.Property(entity => entity.EndDate)
            .IsRequired();
    }
}