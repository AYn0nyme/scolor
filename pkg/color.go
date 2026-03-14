package pkg

import "fmt"

var Colors = map[string]string{
	"red":     "\033[31m",
	"green":   "\033[32m",
	"yellow":  "\033[33m",
	"blue":    "\033[34m",
	"magenta": "\033[35m",
	"cyan":    "\033[36m",
	"gray":    "\033[37m",
	"white":   "\033[97m",
}
var Reset = "\033[0m"

func DisplayText(color string, text string) {
	if _, ok := Colors[color]; !ok {
		panic("Color not found")
	}

	fmt.Printf("%s%s%s", Colors[color], text, Reset)

}
