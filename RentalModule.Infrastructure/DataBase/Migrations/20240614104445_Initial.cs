using System;
using Microsoft.EntityFrameworkCore.Migrations;

#nullable disable

namespace RentalModule.Infrastructure.DataBase.Migrations
{
    /// <inheritdoc />
    public partial class Initial : Migration
    {
        /// <inheritdoc />
        protected override void Up(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.CreateTable(
                name: "CarOffers",
                columns: table => new
                {
                    Id = table.Column<Guid>(type: "uniqueidentifier", nullable: false),
                    Heading = table.Column<string>(type: "nvarchar(200)", maxLength: 200, nullable: false),
                    ShortDescription = table.Column<string>(type: "nvarchar(500)", maxLength: 500, nullable: false),
                    FeaturedImageUrl = table.Column<string>(type: "nvarchar(200)", maxLength: 200, nullable: false),
                    UrlHandle = table.Column<string>(type: "nvarchar(100)", maxLength: 100, nullable: false),
                    Horsepower = table.Column<string>(type: "nvarchar(50)", maxLength: 50, nullable: false),
                    YearOfProduction = table.Column<int>(type: "int", nullable: false),
                    EngineDetails = table.Column<string>(type: "nvarchar(100)", maxLength: 100, nullable: false),
                    DriveDetails = table.Column<string>(type: "nvarchar(100)", maxLength: 100, nullable: false),
                    GearboxDetails = table.Column<string>(type: "nvarchar(100)", maxLength: 100, nullable: false),
                    CarDeliverylocation = table.Column<string>(type: "nvarchar(200)", maxLength: 200, nullable: false),
                    CarReturnLocation = table.Column<string>(type: "nvarchar(200)", maxLength: 200, nullable: false),
                    PublishedDate = table.Column<DateTime>(type: "datetime2", nullable: false),
                    Visible = table.Column<bool>(type: "bit", nullable: false)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_CarOffers", x => x.Id);
                });

            migrationBuilder.CreateTable(
                name: "CarOrders",
                columns: table => new
                {
                    Id = table.Column<Guid>(type: "uniqueidentifier", nullable: false),
                    UserId = table.Column<Guid>(type: "uniqueidentifier", nullable: false),
                    CarOfferId = table.Column<Guid>(type: "uniqueidentifier", nullable: false),
                    StartDate = table.Column<DateTime>(type: "datetime2", nullable: false),
                    EndDate = table.Column<DateTime>(type: "datetime2", nullable: false),
                    Notes = table.Column<string>(type: "nvarchar(1000)", maxLength: 1000, nullable: true),
                    NumOfDrivers = table.Column<int>(type: "int", nullable: false),
                    TotalCost = table.Column<double>(type: "float", nullable: false)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_CarOrders", x => x.Id);
                });

            migrationBuilder.CreateTable(
                name: "CarTags",
                columns: table => new
                {
                    Id = table.Column<Guid>(type: "uniqueidentifier", nullable: false),
                    Name = table.Column<string>(type: "nvarchar(100)", maxLength: 100, nullable: false),
                    CarOfferId = table.Column<Guid>(type: "uniqueidentifier", nullable: false)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_CarTags", x => x.Id);
                    table.ForeignKey(
                        name: "FK_CarTags_CarOffers_CarOfferId",
                        column: x => x.CarOfferId,
                        principalTable: "CarOffers",
                        principalColumn: "Id",
                        onDelete: ReferentialAction.Cascade);
                });

            migrationBuilder.CreateTable(
                name: "CarTariff",
                columns: table => new
                {
                    Id = table.Column<Guid>(type: "uniqueidentifier", nullable: false),
                    CarOfferId = table.Column<Guid>(type: "uniqueidentifier", nullable: false),
                    OneNormalDayPrice = table.Column<double>(type: "float", nullable: false),
                    OneWeekendDayPrice = table.Column<double>(type: "float", nullable: false),
                    FullWeekendPrice = table.Column<double>(type: "float", nullable: false),
                    OneWeekPrice = table.Column<double>(type: "float", nullable: false),
                    OneMonthPrice = table.Column<double>(type: "float", nullable: false)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_CarTariff", x => x.Id);
                    table.ForeignKey(
                        name: "FK_CarTariff_CarOffers_CarOfferId",
                        column: x => x.CarOfferId,
                        principalTable: "CarOffers",
                        principalColumn: "Id",
                        onDelete: ReferentialAction.Cascade);
                });

            migrationBuilder.CreateTable(
                name: "ImageUrls",
                columns: table => new
                {
                    Id = table.Column<Guid>(type: "uniqueidentifier", nullable: false),
                    CarOfferId = table.Column<Guid>(type: "uniqueidentifier", nullable: false),
                    Url = table.Column<string>(type: "nvarchar(500)", maxLength: 500, nullable: false)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_ImageUrls", x => x.Id);
                    table.ForeignKey(
                        name: "FK_ImageUrls_CarOffers_CarOfferId",
                        column: x => x.CarOfferId,
                        principalTable: "CarOffers",
                        principalColumn: "Id",
                        onDelete: ReferentialAction.Cascade);
                });

            migrationBuilder.CreateTable(
                name: "TimePeriodEntity",
                columns: table => new
                {
                    Id = table.Column<Guid>(type: "uniqueidentifier", nullable: false),
                    CarOfferId = table.Column<Guid>(type: "uniqueidentifier", nullable: false),
                    StartDate = table.Column<DateTime>(type: "datetime2", nullable: false),
                    EndDate = table.Column<DateTime>(type: "datetime2", nullable: false)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_TimePeriodEntity", x => x.Id);
                    table.ForeignKey(
                        name: "FK_TimePeriodEntity_CarOffers_CarOfferId",
                        column: x => x.CarOfferId,
                        principalTable: "CarOffers",
                        principalColumn: "Id",
                        onDelete: ReferentialAction.Cascade);
                });

            migrationBuilder.CreateIndex(
                name: "IX_CarTags_CarOfferId",
                table: "CarTags",
                column: "CarOfferId");

            migrationBuilder.CreateIndex(
                name: "IX_CarTariff_CarOfferId",
                table: "CarTariff",
                column: "CarOfferId",
                unique: true);

            migrationBuilder.CreateIndex(
                name: "IX_ImageUrls_CarOfferId",
                table: "ImageUrls",
                column: "CarOfferId");

            migrationBuilder.CreateIndex(
                name: "IX_TimePeriodEntity_CarOfferId",
                table: "TimePeriodEntity",
                column: "CarOfferId");
        }

        /// <inheritdoc />
        protected override void Down(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.DropTable(
                name: "CarOrders");

            migrationBuilder.DropTable(
                name: "CarTags");

            migrationBuilder.DropTable(
                name: "CarTariff");

            migrationBuilder.DropTable(
                name: "ImageUrls");

            migrationBuilder.DropTable(
                name: "TimePeriodEntity");

            migrationBuilder.DropTable(
                name: "CarOffers");
        }
    }
}
