package main

import (
	"fmt"
	"time"
)

func main() {
	tijd := time.Now()
	tijdinuur := tijd.Hour()
	if tijdinuur >= 7 && tijdinuur < 12 {
		fmt.Println("Goedemorgen! Welkom bij Fonteyn Vakantieparken")
	} else if tijdinuur >= 12 && tijdinuur < 18 {
		fmt.Println("Goedemiddag! Welkom bij Fonteyn Vakantieparken")
	} else if tijdinuur > 18 && tijdinuur < 23 {
		fmt.Println("Goedenavond! Welkom bij Fonteyn Vakantieparken")
	} else if tijdinuur > 23 && tijdinuur < 7 {
		fmt.Println("Sorry, de parkeerplaats is ’s nachts gesloten")
	}
}
