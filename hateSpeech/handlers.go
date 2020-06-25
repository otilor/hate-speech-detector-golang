package hateSpeech

import "log"

func handleError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}