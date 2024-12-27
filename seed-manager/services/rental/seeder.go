package rental

import (
	"fmt"
	"log"
	"seeder-manager/api"
	"seeder-manager/config"
	"seeder-manager/models"
	"seeder-manager/reference_store"
	"seeder-manager/services/rental/factories"
)

type RentalSeeder struct {
	carFactory      *factories.CarFactory
	carOrderFactory *factories.CarOrderFactory
	apiClient       *api.APIClient
	cfg             *config.Config
}

func NewRentalSeeder(apiBaseURL string, store *reference_store.InMemoryStore) *RentalSeeder {
	return &RentalSeeder{
		carFactory:      factories.NewCarFactory(store),
		carOrderFactory: factories.NewCarOrderFactory(store),
		apiClient:       api.NewAPIClient(apiBaseURL),
		cfg:             config.GetConfig(),
	}
}

type CreateCarOfferRequest struct {
	Heading            string   `json:"heading" binding:"required" example:"Car Offer Title" swaggertype:"string" validate:"required"`
	ShortDescription   string   `json:"shortDescription" binding:"required" example:"Short description" swaggertype:"string" validate:"required"`
	FeaturedImageUrl   string   `json:"featuredImageUrl" example:"https://example.com/image.jpg" swaggertype:"string"`
	UrlHandle          string   `json:"urlHandle" binding:"required" example:"car-offer-title" swaggertype:"string" validate:"required"`
	Horsepower         string   `json:"horsepower" example:"300" swaggertype:"string"`
	YearOfProduction   int      `json:"yearOfProduction" example:"2023" swaggertype:"integer"`
	EngineDetails      string   `json:"engineDetails" example:"2.0L Turbo" swaggertype:"string"`
	DriveDetails       string   `json:"driveDetails" example:"AWD" swaggertype:"string"`
	GearboxDetails     string   `json:"gearboxDetails" example:"Automatic" swaggertype:"string"`
	Visible            bool     `json:"visible" example:"true" swaggertype:"boolean"`
	OneNormalDayPrice  float64  `json:"oneNormalDayPrice" example:"100.00" swaggertype:"number"`
	OneWeekendDayPrice float64  `json:"oneWeekendDayPrice" example:"150.00" swaggertype:"number"`
	OneWeekPrice       float64  `json:"oneWeekPrice" example:"600.00" swaggertype:"number"`
	OneMonthPrice      float64  `json:"oneMonthPrice" example:"2000.00" swaggertype:"number"`
	Tags               []string `json:"tags" example:"[\"luxury\",\"sports\"]" swaggertype:"array,string"`
	ImageUrls          []string `json:"imageUrls" example:"[\"https://example.com/image1.jpg\"]" swaggertype:"array,string"`
}

type CreateCarOrderRequest struct {
	CarOfferId       string  `json:"carOfferId" binding:"required" example:"123e4567-e89b-12d3-a456-426614174000" swaggertype:"string" validate:"required"`
	StartDate        string  `json:"startDate" binding:"required" example:"2024-03-15T14:30:00Z" swaggertype:"string" validate:"required,datetime,futuredate"`
	EndDate          string  `json:"endDate" binding:"required" example:"2024-03-20T12:00:00Z" swaggertype:"string" validate:"required,datetime,gtdate=StartDate"`
	DeliveryLocation string  `json:"deliveryLocation" example:"City Center" swaggertype:"string"`
	ReturnLocation   string  `json:"returnLocation" example:"City Center" swaggertype:"string"`
	NumOfDrivers     int     `json:"numOfDrivers" example:"2" swaggertype:"integer"`
	TotalCost        float64 `json:"totalCost" example:"750.00" swaggertype:"number"`
	Status           string  `json:"status" example:"pending" swaggertype:"string" validate:"required,validCarOrderStatus"`
}

func (s *RentalSeeder) createCarOffer(car *models.CarModel, token string) (string, error) {
	request := CreateCarOfferRequest{
		Heading:            car.Heading,
		ShortDescription:   car.ShortDescription,
		FeaturedImageUrl:   car.FeaturedImageUrl,
		UrlHandle:          car.UrlHandle,
		Horsepower:         car.Horsepower,
		YearOfProduction:   car.YearOfProduction,
		EngineDetails:      car.EngineDetails,
		DriveDetails:       car.DriveDetails,
		GearboxDetails:     car.GearboxDetails,
		Visible:            car.Visible,
		OneNormalDayPrice:  car.OneNormalDayPrice,
		OneWeekendDayPrice: car.OneWeekendDayPrice,
		OneWeekPrice:       car.OneWeekPrice,
		OneMonthPrice:      car.OneMonthPrice,
		Tags:               car.Tags,
		ImageUrls:          car.ImageUrls,
	}

	resp, err := s.apiClient.Post("/rental-api/api/car-offers", request, token)
	if err != nil {
		return "", fmt.Errorf("error creating car offer: %w", err)
	}

	if resp.Id == "" {
		return "", fmt.Errorf("car offer created but no ID returned")
	}

	return resp.Id, nil
}

func (s *RentalSeeder) createCarOrder(order *models.CarOrderModel, token string) error {
	request := CreateCarOrderRequest{
		CarOfferId:       order.CarOfferId,
		StartDate:        order.StartDate,
		EndDate:          order.EndDate,
		DeliveryLocation: order.DeliveryLocation,
		ReturnLocation:   order.ReturnLocation,
		NumOfDrivers:     order.NumOfDrivers,
		TotalCost:        order.TotalCost,
		Status:           order.Status,
	}

	_, err := s.apiClient.Post("/rental-api/api/car-orders", request, token)
	if err != nil {
		return fmt.Errorf("error creating car order: %w", err)
	}

	return nil
}

func (s *RentalSeeder) Seed(store *reference_store.InMemoryStore, token string) error {
	log.Printf("Starting to seed %d cars with %d orders each...", s.cfg.SeedCount.Cars, s.cfg.SeedCount.OrdersPerCar)

	// Create cars first
	cars, err := s.carFactory.CreateMany(s.cfg.SeedCount.Cars)
	if err != nil {
		return err
	}

	// Initialize car order factory with store
	s.carOrderFactory = factories.NewCarOrderFactory(store)

	// Create car offers and orders
	for _, car := range cars { // Create orders for all cars
		// Create car offer first
		carOfferId, err := s.createCarOffer(car, token)
		if err != nil {
			return err
		}

		// Create orders per car
		for i := 0; i < s.cfg.SeedCount.OrdersPerCar; i++ {
			order, err := s.carOrderFactory.CreateWithCar(carOfferId)
			if err != nil {
				return err
			}

			err = s.createCarOrder(order, token)
			if err != nil {
				return err
			}
		}
	}

	log.Printf("Completed seeding %d cars with %d orders each", s.cfg.SeedCount.Cars, s.cfg.SeedCount.OrdersPerCar)
	return nil
}

func (s *RentalSeeder) Cleanup() error {
	// TODO: Implement cleanup logic if needed
	return nil
}
