package solitaire

func printCipherText(ciphertext []byte) {
	// print the ciphertext in blocks of 5 characters
	for i := 0; i < len(ciphertext); i += 5 {
		if i+5 < len(ciphertext) {
			print(string(ciphertext[i : i+5]))
		} else {
			print(string(ciphertext[i:]))
		}
		print(" ")
	}
	// print a newline at the end
	println()
}
