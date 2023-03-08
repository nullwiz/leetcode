package main

func CaesarCipherEncryptor(str string, key int) string {
	// Write your code here.
	if key > 26 {
		key -= 26
	}
	//
	slice := make([]byte, len(str))
	for k, v := range str {
		char := int(v) + key
		if char > 122 {
			char -= 26
		}
		slice[k] = byte(char)
	}
	return string(slice)
}
