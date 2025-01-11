package main

import (
	"fmt"
	"os/exec"
)

func speak(word string) {
	// Google TTS API (fetches an MP3 file)
	url := fmt.Sprintf("https://translate.google.com/translate_tts?ie=UTF-8&client=tw-ob&q=%s&tl=en", word)
	cmd := exec.Command("mpg123", url)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error playing audio:", err)
	}
}

func main() {
	word := "Since you're building a Dictionary CLI in Go, you're already working with API calls, JSON handling, and CLI interactions. Once you finish this, here’s what you should learn next in Go to level up:"
	speak(word)
}
