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

func (g *GeneratorService) GenerateRandomSlice(length int) domain.Sequence {
	result := make(domain.Sequence, length)
	for i := range length {
		result[i] = AllWaysList[rand.Intn(len(AllWaysList))]
	}
	return result
}

func (g *GeneratorService) GenerateRandomString() string {
	result := make([]byte, 15)
	for i := range 15 {
		chars := []byte{'1', 'X', '2'}
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}

func (g *GeneratorService) GenerateSequenceConfigs() []domain.SequenceConfig {
	configs := make([]domain.SequenceConfig, 10)
	for i := range 10 {
		configs[i] = domain.SequenceConfig{
			Seq:       g.GenerateRandomSlice(15),
			BorderMin: 8,
			BorderMax: 12,
		}
	}
	return configs
}

func (g *GeneratorService) GenerateManyTargets(count int) []string {
	targets := make([]string, count)
	for i := range count {
		targets[i] = g.GenerateRandomString()
	}
	return targets
}
