package services

import (
	"packetCutter/internal/domain"
	"packetCutter/internal/storage"
)

type MatcherService struct {
	collection *storage.ResultCollection
}

func NewMatcherService(collection *storage.ResultCollection) *MatcherService {
	return &MatcherService{collection: collection}
}

func (s *MatcherService) CountMatchesForPrediction(variant string, pred domain.Prediction) (int, error) {
	targetLen := len(variant)
	if len(pred) != targetLen {
		return 0, domain.ErrPredictLength
	}

	count := 0
	for i := range targetLen {
		for j := 0; j < len(pred[i]); j++ {
			if pred[i][j] == variant[i] {
				count++
				break
			}
		}
	}
	return count, nil
}

func (s *MatcherService) CountMatchesForPredSequence(variant string, configs *[]domain.PredictionConfig) error {
	if len(variant) != 15 {
		return domain.ErrInvalidVariantLength
	}

	if len(*configs) != 10 {
		return domain.ErrInvalidPredSeqLength
	}

	predictionMatches := make([]byte, 10)
	allInRange := true
	var firstError error

	for i, config := range *configs {
		if config.BorderMin > config.BorderMax {
			predictionMatches[i] = '0'
			allInRange = false
			if firstError == nil {
				firstError = domain.ErrInvalidBorders
			}
			continue
		}

		count, err := s.CountMatchesForPrediction(variant, config.Pred)
		if err != nil {
			predictionMatches[i] = '0'
			allInRange = false
			if firstError == nil {
				firstError = err
			}
			continue
		}

		predictionMatches[i] = s.byteToHex(count)
		if count < config.BorderMin || count > config.BorderMax {
			allInRange = false
		}
	}

	if allInRange {
		hexResult := string(predictionMatches)
		s.collection.Add(hexResult, variant)
	}

	return firstError
}

func (s *MatcherService) byteToHex(n int) byte {
	if n < 10 {
		return byte('0' + n)
	}
	return byte('A' + n - 10)
}
