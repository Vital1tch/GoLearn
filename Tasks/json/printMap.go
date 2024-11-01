package json

import "fmt"

func printMap(data map[string]interface{}, indent string) {
	for key, value := range data {
		if innerMap, ok := value.(map[string]interface{}); ok {
			fmt.Printf("%s%s:\n", indent, key)
			printMap(innerMap, indent+"  ") // Углубляем отступы
		} else {
			fmt.Printf("%s%s: %v\n", indent, key, value)
		}
	}
}
