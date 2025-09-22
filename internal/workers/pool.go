package workers

import (
	"packetCutter/internal/domain"
	"packetCutter/internal/services"
	"packetCutter/internal/storage"
	"runtime"
	"sync"
)

type WorkerPool struct {
	matcherService *services.MatcherService
	collection     *storage.ResultCollection
	workerCount    int
	taskCh         chan Task
	wg             sync.WaitGroup
	errCh          chan error
}

func NewWorkerPool(matcherService *services.MatcherService) *WorkerPool {
	workerCount := max(runtime.NumCPU(), 2)

	wp := &WorkerPool{
		matcherService: matcherService,
		collection:     storage.NewResultCollection(),
		workerCount:    workerCount,
		taskCh:         make(chan Task, workerCount*10),
		errCh:          make(chan error, workerCount*100),
	}

	wp.wg.Add(workerCount)
	for range workerCount {
		go wp.worker()
	}

	return wp
}

func (wp *WorkerPool) worker() {
	defer wp.wg.Done()
	for task := range wp.taskCh {
		err := wp.matcherService.CountMatchesForPredSequence(task.Target, task.Configs)
		if err != nil {
			select {
			case wp.errCh <- err:
			default:
				// Пропускаем ошибку если канал переполнен
			}
		}
	}
}

func (wp *WorkerPool) Submit(target string, configs *[]domain.PredictionConfig) {
	wp.taskCh <- Task{Target: target, Configs: configs}
}

func (wp *WorkerPool) SubmitBatch(targets []string, configs *[]domain.PredictionConfig) {
	for _, target := range targets {
		wp.Submit(target, configs)
	}
}

func (wp *WorkerPool) Wait() []error {
	close(wp.taskCh)
	wp.wg.Wait()
	close(wp.errCh)

	errors := make([]error, 0)
	for err := range wp.errCh {
		errors = append(errors, err)
	}
	return errors
}

func (wp *WorkerPool) GetCollection() *storage.ResultCollection {
	return wp.collection
}

func (wp *WorkerPool) GetStats() map[string]any {
	return map[string]any{
		"workers":       wp.workerCount,
		"results_count": wp.collection.Count(),
		"unique_keys":   wp.collection.Len(),
	}
}
