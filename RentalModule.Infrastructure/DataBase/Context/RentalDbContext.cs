using Microsoft.EntityFrameworkCore;
using RentalModule.Infrastructure.DataBase.Entities;

namespace RentalModule.Infrastructure.DataBase.Context;

public class RentalDbContext : DbContext
{
    public DbSet<CarTariffEntity> CarTariff { get; set; } = null!;
    public DbSet<ImageUrlEntity> ImageUrls { get; set; } = null!;
    public DbSet<CarTagEntity> CarTags { get; set; } = null!;
    public DbSet<CarOfferEntity> CarOffers { get; set; } = null!;
    public DbSet<CarOrderEntity> CarOrders { get; set; } = null!;
    
    public RentalDbContext()
    {
    }
    

    public RentalDbContext(DbContextOptions<RentalDbContext> options) : base(options)
    {
    }
    
    protected override void OnModelCreating(ModelBuilder modelBuilder)
    {
        modelBuilder.ApplyConfigurationsFromAssembly(typeof(RentalDbContext).Assembly);
    }
}