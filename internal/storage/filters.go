package storage

// FilterByTargetCount возвращает новую коллекцию с элементами,
// у которых количество targets в указанном диапазоне
func FilterByTargetCount(rc *ResultCollection, minCount, maxCount int) *ResultCollection {
	filtered := NewResultCollection()

	data := rc.GetAll()
	for hexResult, targets := range data {
		count := len(targets)
		if count >= minCount && count <= maxCount {
			for _, target := range targets {
				filtered.Add(hexResult, target)
			}
		}
	}

	return filtered
}

// FilterSingleTargets возвращает новую коллекцию только с элементами,
// у которых ровно один target в значении
func FilterSingleTargets(rc *ResultCollection) *ResultCollection {
	filtered := NewResultCollection()

	data := rc.GetAll()
	for hexResult, targets := range data {
		if len(targets) == 1 {
			filtered.Add(hexResult, targets[0])
		}
	}

	return filtered
}
