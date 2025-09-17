package domain

type Sequence []string

type SequenceConfig struct {
	Seq       Sequence
	BorderMin int
	BorderMax int
}

type MatchResult struct {
	HexResult string
	Target    string
	IsValid   bool
}
