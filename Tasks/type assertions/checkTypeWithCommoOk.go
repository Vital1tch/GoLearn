package type_assertions

import "fmt"

func checkType(m map[string]interface{}) {
	for key, item := range m {
		switch value := item.(type) {
		case string:
			fmt.Printf("%s это строка со значением: %s\n", key, value)
		case int:
			fmt.Printf("%s это число с значением: %d\n", key, value)
		case float64:
			fmt.Printf("%s это вещественное число с значением: %.2f\n", key, value)
		default:
			fmt.Printf("Неизвестнчй тип %s\n", key)
		}
	}
}
