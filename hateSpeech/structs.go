package hateSpeech

import "fmt"

type SpeechInterface interface {
	isRelatedToPresident(text string) bool
}

type speech struct {
	body string
}

func (s speech) isRelatedToPresident() bool {
	var hayStack = make(map[interface{}]interface{})
	hayStack["femi"] = "James"
	fmt.Println(hayStack)
	contains("Buhari", hayStack)
	if s.body == "Buhari" {
		return true
	}
	return false
}