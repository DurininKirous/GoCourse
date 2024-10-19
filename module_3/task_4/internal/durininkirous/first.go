package durininkirous

import "module_4/task_4/pkg/wordz"

func City() string {
	cities := map[string]string{
		"one":   "Moscow",
		"two":   "Kiev",
		"three": "New-York",
		"four":  "Tokyo",
		"five":  "Rom",
		"six":   "Oslo",
		"seven": "St-Petersbug",
	}
	return cities[wordz.Random()]
}
