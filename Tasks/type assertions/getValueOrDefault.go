package type_assertions

func getValueOrDefault(m map[string]int, key string) int {
	if value, ok := m[key]; ok {
		return value
	}

	return -1
}
