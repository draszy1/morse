package main

import (
	"bufio"
	"fmt"
	"os"
	"testgo/morse"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')

	if morse.IsMorseCode(str) {
		text := morse.MorseToText(str)
		fmt.Println(text)
	} else {
		morseCode := morse.TextToMorse(str)
		fmt.Println(morseCode)
	}
}
