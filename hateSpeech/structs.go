package hateSpeech

type SpeechInterface interface {
	isRelatedToPresident(text string) bool
}

type speech struct {
	body string
}

func (s speech) isRelatedToPresident() bool {
	if s.body == "Buhari" {
		return true
	}
	return false
}