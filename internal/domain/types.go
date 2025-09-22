package domain

type Prediction []string

type PredictionConfig struct {
	Pred      Prediction
	BorderMin int
	BorderMax int
}

type VariantData struct {
	HexPredMatches string
	Variant        string
	IsValid        bool
}
