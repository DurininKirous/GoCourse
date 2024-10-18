package utils

func Contains(str string, slices []string) bool {
	var res bool = false
	for _, slice := range slices {
		if slice == str {
			res = true
			break
		}
	}
	return res
}
