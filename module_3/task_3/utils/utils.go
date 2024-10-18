package utils

func Contains(a string, slices []string) bool {
	var res bool = false
	for _, slice := range slices {
		if slice == a {
			res = true
			break
		}
	}
	return res
}
