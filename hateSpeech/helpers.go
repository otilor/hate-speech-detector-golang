package hateSpeech

import "fmt"

func contains (needle interface{}, haystack map[interface{}]interface{}) bool {
	sizeOfHaystack := len(haystack)

	hashMap := make(map[interface{}]interface{})

	// Store each element in a hashmap
	for i:= 0; i < sizeOfHaystack; i++ {
			hashMap[i] = haystack
	}

	fmt.Println(hashMap)
	// loop through the hashmap and check if an element does not exist it it.

	for i := range haystack {
		if needle != haystack[i] {
			return false
		}
	}

	return true
}
