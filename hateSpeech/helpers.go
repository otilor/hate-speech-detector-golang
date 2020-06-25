package hateSpeech

func contains (needle interface{}, haystack map[interface{}]interface{}) bool {
	sizeOfHaystack := len(haystack)
	var hashMap map[interface{}]interface{}

	// Store each element in a hashmap
	for i:= 0; i < sizeOfHaystack; i++ {
			hashMap[i] = haystack[i]
	}

	// loop through the hashmap and check if an element does not exist it it.
	for i := range haystack {
		if needle != haystack[i] {
			return false
		}
	}
	return true
}
