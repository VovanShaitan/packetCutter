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

func (s *MatcherService) CountMatchesForSequence(target string, seq domain.Sequence) (int, error) {
	targetLen := len(target)
	if len(seq) != targetLen {
		return 0, domain.ErrSequenceLength
	}

	count := 0
	for i := 0; i < targetLen; i++ {
		for j := 0; j < len(seq[i]); j++ {
			if seq[i][j] == target[i] {
				count++
				break
			}
		}
	}
	return count, nil
}

func (s *MatcherService) ProcessTarget(target string, configs []domain.SequenceConfig) error {
	return s.countMatchesGenericUltraOptimized(target, configs, s.collection)
}

func (s *MatcherService) countMatchesGenericUltraOptimized(
	target string,
	configs []domain.SequenceConfig,
	collection *storage.ResultCollection,
) error {
	if len(target) != 15 {
		return domain.ErrInvalidTargetLength
	}

	if len(configs) != 10 {
		return domain.ErrInvalidConfigLength
	}

	result := make([]byte, 10)
	allInRange := true
	var firstError error

	for i, config := range configs {
		if config.BorderMin > config.BorderMax {
			result[i] = '0'
			allInRange = false
			if firstError == nil {
				firstError = domain.ErrInvalidBorders
			}
			continue
		}

		count, err := s.CountMatchesForSequence(target, config.Seq)
		if err != nil {
			result[i] = '0'
			allInRange = false
			if firstError == nil {
				firstError = err
			}
			continue
		}

		result[i] = s.byteToHex(count)
		if count < config.BorderMin || count > config.BorderMax {
			allInRange = false
		}
	}

	if allInRange {
		hexResult := string(result)
		collection.Add(hexResult, target)
	}

	return firstError
}

func (s *MatcherService) byteToHex(n int) byte {
	if n < 10 {
		return byte('0' + n)
	}
	return byte('A' + n - 10)
}
