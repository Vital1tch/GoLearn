package loops

import "fmt"

func ReverseString(sentence string) {

	sentenceReversed := make([]byte, 0, len(sentence))
	for i := len(sentence) - 1; i >= 0; i-- {
		sentenceReversed = append(sentenceReversed, sentence[i])
		fmt.Println(string(sentence[i]))
	}
	fmt.Println(string(sentenceReversed))
}
