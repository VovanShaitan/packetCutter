package workers

import "packetCutter/internal/domain"

type Task struct {
	Target  string
	Configs []domain.SequenceConfig
}

type TaskResult struct {
	HexResult string
	Target    string
	Error     error
}
