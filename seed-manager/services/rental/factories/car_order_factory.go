package factories

import (
	"math/rand"
	"time"

	"seeder-manager/models"
	"seeder-manager/reference_store"

	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
)

type CarOrderFactory struct {
	refStore *reference_store.InMemoryStore
}

func NewCarOrderFactory(refStore *reference_store.InMemoryStore) *CarOrderFactory {
	return &CarOrderFactory{
		refStore: refStore,
	}
}

func (f *CarOrderFactory) Create() (*models.CarOrderModel, error) {
	return f.CreateWithCar(uuid.New().String())
}

func (f *CarOrderFactory) CreateWithCar(carID string) (*models.CarOrderModel, error) {
	randInt1, err := faker.RandomInt(1, 90, 1)
	if err != nil {
		return nil, err
	}
	randInt2, err := faker.RandomInt(1, 14, 1)
	if err != nil {
		return nil, err
	}
	randInt3, err := faker.RandomInt(100, 1000, 1)
	if err != nil {
		return nil, err
	}
	// Generate random dates within next 3 months
	startDate := time.Now().AddDate(0, 0, randInt1[0])
	endDate := startDate.AddDate(0, 0, randInt2[0])

	availableStatuses := []string{
		"PENDING",
		"PREPARING",
		"READY",
		"ACTIVE",
		"COMPLETED",
		"CANCELLED",
		"ARCHIVED",
	}

	randomStatus := availableStatuses[rand.Intn(len(availableStatuses))]

	return &models.CarOrderModel{
		CarOfferId:       carID,
		StartDate:        startDate.Format(time.RFC3339),
		EndDate:          endDate.Format(time.RFC3339),
		DeliveryLocation: faker.GetRealAddress().Address,
		ReturnLocation:   faker.GetRealAddress().Address,
		NumOfDrivers:     2,
		Status:           randomStatus,
		TotalCost:        float64(randInt3[0]),
	}, nil
}
