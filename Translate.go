package main

import (
    "fmt"
    "strconv"
)

func main() {
    var hexCode string
    fmt.Println("Enter a hex color code (without the # symbol):")
    fmt.Scanln(&hexCode)

    r, _ := strconv.ParseUint(hexCode[0:2], 16, 0)
    g, _ := strconv.ParseUint(hexCode[2:4], 16, 0)
    b, _ := strconv.ParseUint(hexCode[4:6], 16, 0)

    ansiEscape := fmt.Sprintf(`\033[38;2;%d;%d;%dm`, r, g, b)
    fmt.Printf("Hex code #%s translates to ANSI escape code %s\n", hexCode, ansiEscape)
}
