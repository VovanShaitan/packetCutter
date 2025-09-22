package generators

import (
	"testing"
)

func TestGeneratorService_GenerateRandomString(t *testing.T) {
	generator := NewGeneratorService()

	str := generator.GenerateRandomVariant()
	if len(str) != 15 {
		t.Errorf("Expected string length 15, got %d", len(str))
	}

	// Проверяем, что строка содержит только допустимые символы
	for _, char := range str {
		if char != '1' && char != 'X' && char != '2' {
			t.Errorf("Invalid character in generated string: %c", char)
		}
	}
}

func TestGeneratorService_GenerateRandomSlice(t *testing.T) {
	generator := NewGeneratorService()

	length := 10
	slice := generator.GenerateRandomPredict(length)
	if len(slice) != length {
		t.Errorf("Expected slice length %d, got %d", length, len(slice))
	}

	// Проверяем, что все элементы содержат только допустимые символы
	for _, item := range slice {
		for _, char := range item {
			if char != '1' && char != 'X' && char != '2' {
				t.Errorf("Invalid character in generated slice item: %c", char)
			}
		}
	}
}

func TestGeneratorService_GenerateSequenceConfigs(t *testing.T) {
	generator := NewGeneratorService()

	configs := generator.GeneratePredictionConfigs()
	if len(configs) != 10 {
		t.Errorf("Expected 10 configs, got %d", len(configs))
	}

	for i, config := range configs {
		if len(config.Pred) != 15 {
			t.Errorf("Config %d: expected sequence length 15, got %d", i, len(config.Pred))
		}
	}
}
