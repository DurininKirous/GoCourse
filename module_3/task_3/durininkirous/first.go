package durininkirous

import "task_3/wordz"

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
