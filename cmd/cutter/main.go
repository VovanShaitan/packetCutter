package main

import (
	"fmt"
	"packetCutter/internal/generators"
	"packetCutter/internal/services"
	"packetCutter/internal/storage"
	"packetCutter/internal/workers"
	"runtime"
	"time"
)

func main() {
	fmt.Printf("🚀 Запуск Awesome Matcher на %d ядрах\n", runtime.NumCPU())

	// Инициализация зависимостей
	collection := storage.NewResultCollection()
	generator := generators.NewGeneratorService()
	matcher := services.NewMatcherService(collection)
	workerPool := workers.NewWorkerPool(matcher)

	// Генерация тестовых данных
	targetCount := 100000
	fmt.Printf("🎲 Генерация %d targets...\n", targetCount)

	targets := generator.GenerateVariants(targetCount)
	predictionsConfigs := generator.GeneratePredictionConfigs()

	// Обработка
	fmt.Printf("⚡ Обработка %d targets...\n", targetCount)
	start := time.Now()

	workerPool.SubmitBatch(targets, &predictionsConfigs)
	errors := workerPool.Wait()

	elapsed := time.Since(start)

	// Вывод результатов
	printResults(workerPool, elapsed, errors, targetCount)

	// Применяем фильтр и показываем результаты
	fmt.Println("\n🔍 Фильтрация результатов (только single targets):")
	filteredCollection := storage.FilterSingleTargets(workerPool.GetCollection())
	printFilteredResults(filteredCollection)
}

func printResults(pool *workers.WorkerPool, elapsed time.Duration, errors []error, total int) {
	stats := pool.GetStats()

	fmt.Printf("\n📊 Результаты обработки:\n")
	fmt.Printf("   Время: %v\n", elapsed)
	fmt.Printf("   Скорость: %.0f ops/sec\n", float64(total)/elapsed.Seconds())
	fmt.Printf("   Обработано: %d targets\n", stats["results_count"])
	fmt.Printf("   Уникальных ключей: %d\n", stats["unique_keys"])
	fmt.Printf("   Ошибок: %d\n", len(errors))
	fmt.Printf("   Воркеров: %d\n", stats["workers"])
}

func printFilteredResults(collection *storage.ResultCollection) {
	fmt.Printf("   Single targets: %d keys\n", collection.Len())
	fmt.Printf("   Всего single targets: %d\n", collection.Count())
}
