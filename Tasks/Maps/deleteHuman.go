package Maps

func DeleteHuman(human map[string]int, name string) {
	if _, ok := human[name]; ok {
		delete(human, name)
	}
}
