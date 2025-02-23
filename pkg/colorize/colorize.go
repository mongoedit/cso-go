package colorize

import (
	"runtime"
)

var isTrashyWindows bool

var colorMap = map[string]string{
	// esc := "\u001b"
	"reset": "\u001b[0m",

	"red":    "\u001b[31m",
	"green":  "\u001b[32m",
	"yellow": "\u001b[33m",
	"blue":   "\u001b[34m",
	"purple": "\u001b[35m",
	"cyan":   "\u001b[36m",
	"white":  "\u001b[37m",

	"redBright":    "\u001b[31;1m",
	"greenBright":  "\u001b[32;1m",
	"yellowBright": "\u001b[33;1m",
	"blueBright":   "\u001b[34;1m",
	"purpleBright": "\u001b[35;1m",
	"cyanBright":   "\u001b[36;1m",
	"whiteBright":  "\u001b[37;1m",

	"redDim":    "\u001b[31;2m",
	"greenDim":  "\u001b[32;2m",
	"yellowDim": "\u001b[33;2m",
	"blueDim":   "\u001b[34;2m",
	"purpleDim": "\u001b[35;2m",
	"cyanDim":   "\u001b[36;2m",
	"whiteDim":  "\u001b[37;2m",
}

func init() {
    if runtime.GOOS == "windows" {
        k, err := registry.OpenKey(registry.CURRENT_USER, `Console`, registry.QUERY_VALUE|registry.SET_VALUE)
        if err != nil {
			isTrashyWindows = true
            return
        }
     	defer k.Close()
        id, _, err := k.GetIntegerValue("VirtualTerminalLevel")
        if err != nil{
            // fmt.Println("Key is missing trashy set to true")
            isTrashyWindows = true
			return 
        }
        if id == 0 {
            // fmt.Println("key exists trashy set to false", err, id)
            isTrashyWindows = true

        }
    }

}

func colorize(text string, color string) string {

	colorVal, ok := colorMap[color]
	if !ok {
		return text
	}

	if isTrashyWindows {
		return text
	}

	coloredText := string(colorVal) + text + string(colorMap["reset"])
	return coloredText
}

// Standard colors with increased brightness
func RedBright(text string) string {
	return colorize(text, "redBright")
}
func GreenBright(text string) string {
	return colorize(text, "greenBright")
}
func BlueBright(text string) string {
	return colorize(text, "blueBright")
}
func YellowBright(text string) string {
	return colorize(text, "yellowBright")
}
func PurpleBright(text string) string {
	return colorize(text, "purpleBright")
}
func CyanBright(text string) string {
	return colorize(text, "cyanBright")
}
func WhiteBright(text string) string {
	return colorize(text, "whiteBright")
}

// Standard colors with decreased brightness
func RedDim(text string) string {
	return colorize(text, "redDim")
}
func GreenDim(text string) string {
	return colorize(text, "greenDim")
}
func BlueDim(text string) string {
	return colorize(text, "blueDim")
}
func YellowDim(text string) string {
	return colorize(text, "yellowDim")
}
func PurpleDim(text string) string {
	return colorize(text, "purpleDim")
}
func CyanDim(text string) string {
	return colorize(text, "cyanDim")
}
func WhiteDim(text string) string {
	return colorize(text, "whiteDim")
}

// Standard Colors
func Red(text string) string {
	return colorize(text, "red")
}
func Green(text string) string {
	return colorize(text, "green")
}
func Blue(text string) string {
	return colorize(text, "blue")
}
func Yellow(text string) string {
	return colorize(text, "yellow")
}
func Purple(text string) string {
	return colorize(text, "purple")
}
func Cyan(text string) string {
	return colorize(text, "cyan")
}
func White(text string) string {
	return colorize(text, "white")
}
