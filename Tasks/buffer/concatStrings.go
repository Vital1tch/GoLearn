package buffer

import (
	"bytes"
	"fmt"
)

func ConcatStrings(sl []string, separator string) {
	var buf bytes.Buffer

	for index, value := range sl {

		if value == "" {
			value = "null"
		}
		buf.WriteString(value)

		if index < len(sl)-1 {
			buf.WriteString(separator)
		} else {
			buf.WriteString(".")
		}

	}

	fmt.Println(buf.String())
}
