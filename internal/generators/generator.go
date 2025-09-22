package generators

import (
	"math/rand"
	"packetCutter/internal/domain"
)

var AllWaysList = []string{"1", "X", "2", "1X", "12", "X2", "1X2"}

type GeneratorService struct{}

func NewGeneratorService() *GeneratorService {
	return &GeneratorService{}
}

func (g *GeneratorService) GenerateRandomPredict(length int) domain.Prediction {
	result := make(domain.Prediction, length)
	for i := range length {
		result[i] = AllWaysList[rand.Intn(len(AllWaysList))]
	}
	return result
}

func (g *GeneratorService) GenerateRandomVariant() string {
	result := make([]byte, 15)
	for i := range 15 {
		chars := []byte{'1', 'X', '2'}
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}

func (g *GeneratorService) GeneratePredictionConfigs() []domain.PredictionConfig {
	configs := make([]domain.PredictionConfig, 10)
	for i := range 10 {
		randomNunber := rand.Intn(8)
		configs[i] = domain.PredictionConfig{
			Pred:      g.GenerateRandomPredict(15),
			BorderMin: randomNunber,
			BorderMax: randomNunber + 8,
		}
	}
	return configs
}

func (g *GeneratorService) GenerateVariants(count int) []string {
	targets := make([]string, count)
	for i := range count {
		targets[i] = g.GenerateRandomVariant()
	}
	return targets
}
