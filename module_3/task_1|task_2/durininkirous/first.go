package durininkirous

import "module_3/task_1/wordz"

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