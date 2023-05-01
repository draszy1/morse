package morse

import "strings"

var morseCodeToLetter = map[string]string{".-": "A", "-...": "B", "-.-.": "C", "-..": "D", ".": "E", "..-.": "F", "--.": "G", "....": "H", "..": "I", ".---": "J", "-.-": "K", ".-..": "L", "--": "M", "-.": "N", "---": "O", ".--.": "P", "--.-": "Q", ".-.": "R", "...": "S", "-": "T", "..-": "U", "...-": "V", ".--": "W", "-..-": "X", "-.--": "Y", "--..": "Z", ".-.-.-": ".", "--..--": ",", "..--..": "?", "-.-.--": "!", " ": " ", "/": " "}

func IsMorseCode(str string) bool {
	for _, char := range str {
		if char != '.' && char != '-' && char != ' ' && char != '\n' && char != '/' {
			return false
		}
	}
	return true
}

func MorseToText(str string) string {
	var result strings.Builder
	words := strings.Split(strings.TrimSpace(str), "/")
	for i, word := range words {
		letters := strings.Split(word, " ")
		for _, letter := range letters {
			if char, ok := morseCodeToLetter[letter]; ok {
				result.WriteString(char)
			}
		}
		if i < len(words)-1 {
			result.WriteString(" ")
		}
	}
	return result.String()
}
