package morse

import "strings"

var letterToMorseCode = map[rune]string{'A': ".-", 'B': "-...", 'C': "-.-.", 'D': "-..", 'E': ".", 'F': "..-.", 'G': "--.", 'H': "....", 'I': "..", 'J': ".---", 'K': "-.-", 'L': ".-..", 'M': "--", 'N': "-.", 'O': "---", 'P': ".--.", 'Q': "--.-", 'R': ".-.", 'S': "...", 'T': "-", 'U': "..-", 'V': "...-", 'W': ".--", 'X': "-..-", 'Y': "-.--", 'Z': "--..", '.': ".-.-.-", ',': "--..--", '?': "..--..", '!': "-.-.--", ' ': "/"}

func TextToMorse(str string) string {
	var result strings.Builder
	str = strings.ToUpper(str)
	for _, char := range str {
		if code, ok := letterToMorseCode[char]; ok {
			result.WriteString(code)
			result.WriteString(" ")
		}
	}
	return strings.TrimSpace(result.String())
}
