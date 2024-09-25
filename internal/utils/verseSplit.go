package utils

import (
	"strings"
)

func VerseSplit(text string) []string {

	verses := strings.Split(text, "\n")
	counter := 0
	for i:=0; i < len(verses); i+=4 {
		if i + 3 > len(verses) {
			break
		}
		verses[counter] = verses[i] + verses[i+1] + verses[i+2] + verses[i+3]
		counter++
	}
	
	return verses[:counter]
}