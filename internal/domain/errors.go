package domain

import "errors"

var (
	ErrInvalidVariantLength = errors.New("количество исходов в варианте должно быть равно 15")
	ErrInvalidPredSeqLength = errors.New("количество фильтров должно быть равно 10")
	ErrInvalidBorders       = errors.New("минимальная граница не может быть больше максимальной")
	ErrPredictLength        = errors.New("количество исходов в фильтре не совпадает с количеством исходов в варианте")
)
