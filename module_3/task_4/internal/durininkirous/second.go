package durininkirous

import "module_3/task_4/pkg/wordz"

func Digit() string {
	cities := map[string]string{
		"one":   "One",
		"two":   "Two",
		"three": "Three",
		"four":  "Four",
		"five":  "Five",
		"six":   "Six",
		"seven": "Seven",
	}
	return cities[wordz.Random()]
}
