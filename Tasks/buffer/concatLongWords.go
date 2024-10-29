package buffer

import (
	"bytes"
	"fmt"
)

func ConcatLongWords(sl []string, minLength int) {
	var buf bytes.Buffer

	for _, value := range sl {
		if len(value) >= minLength {
			if buf.Len() > 0 {
				buf.WriteString(" ")
			}
			buf.WriteString(value)
		}

	}
	buf.WriteString(".")

	fmt.Println(buf.String())
}
