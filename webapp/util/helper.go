package util

func ContainsString(array []string, searchElem string) bool {
	for _, elem := range array {
		if elem == searchElem {
			return true
		}
	}
	return false
}
