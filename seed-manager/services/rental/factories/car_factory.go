package factories

import (
	"fmt"
	"math/rand"
	"time"

	"seeder-manager/models"
	"seeder-manager/reference_store"

	"github.com/go-faker/faker/v4"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type CarFactory struct {
	refStore *reference_store.InMemoryStore
}

func NewCarFactory(refStore *reference_store.InMemoryStore) *CarFactory {
	return &CarFactory{
		refStore: refStore,
	}
}

func (f *CarFactory) Create() (*models.CarModel, error) {
	// Get random images for the car
	imageIDs := f.refStore.GetRandomCarImageIDs(4) // Get 4 random images
	if len(imageIDs) == 0 {
		return nil, fmt.Errorf("no car images available in the store")
	}

	// Generate random tags
	possibleTags := []string{"luxury", "sports", "suv", "electric", "hybrid", "compact", "family", "performance"}
	numTags := rand.Intn(3) + 1 // 1 to 3 tags
	tags := make([]string, 0, numTags)
	for i := 0; i < numTags; i++ {
		tags = append(tags, possibleTags[rand.Intn(len(possibleTags))])
	}

	randInt, err := faker.RandomInt(0, 4, 1)
	if err != nil {
		return nil, err
	}
	randInt2, err := faker.RandomInt(50, 300, 1)
	if err != nil {
		return nil, err
	}

	brand := faker.Word()
	model := faker.Word()
	year := 2020 + randInt[0]

	featuredImageId := imageIDs[0]
	additionalImageIds := imageIDs

	return &models.CarModel{
		Heading:            fmt.Sprintf("%s %s %d", brand, model, year),
		ShortDescription:   faker.Sentence(),
		FeaturedImageUrl:   featuredImageId,
		UrlHandle:          fmt.Sprintf("%s-%s-%d", brand, model, year),
		Horsepower:         fmt.Sprintf("%d", 150+rand.Intn(400)),
		YearOfProduction:   year,
		EngineDetails:      fmt.Sprintf("%.1fL %s", 1.6+float64(rand.Intn(30))/10.0, []string{"Turbo", "Naturally Aspirated", "Hybrid", "Electric"}[rand.Intn(4)]),
		DriveDetails:       []string{"FWD", "RWD", "AWD", "4WD"}[rand.Intn(4)],
		GearboxDetails:     []string{"Manual 6-speed", "Automatic 8-speed", "CVT", "Dual-clutch 7-speed"}[rand.Intn(4)],
		Visible:            true,
		OneNormalDayPrice:  float64(randInt2[0]),
		OneWeekendDayPrice: float64(randInt2[0]) * 1.3,
		OneWeekPrice:       float64(randInt2[0]) * 6,
		OneMonthPrice:      float64(randInt2[0]) * 25,
		Tags:               tags,
		ImageUrls:          generateImageUrls(additionalImageIds),
	}, nil
}

func generateImageUrls(imageIds []string) []string {
	urls := make([]string, len(imageIds))
	for i, id := range imageIds {
		urls[i] = fmt.Sprintf("/api/files/%s", id)
	}
	return urls
}

func (f *CarFactory) CreateMany(count int) ([]*models.CarModel, error) {
	cars := make([]*models.CarModel, 0, count)
	for i := 0; i < count; i++ {
		car, err := f.Create()
		if err != nil {
			return nil, err
		}
		cars = append(cars, car)
	}
	return cars, nil
}
