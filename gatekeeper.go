package main

import (
	"fmt"
	"time"
)

func boodschap() {
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

func main() {
	kenteken := [10]string{"AB-123-CD", "XY-456-Z", "WZ-789-Y", "MN-111-P", "QA-222-B", "KL-333-J", "HI-444-N", "VG-555-M", "FC-666-T", "BU-777-K"}

	var input string
	fmt.Println("Hallo! Wat is uw Kenteken?")
	fmt.Scan(&input)

	staaterbij := false
	for i := 0; i < len(kenteken); i++ {
		if input == kenteken[i] {
			staaterbij = true
			break
		}
	}

	if staaterbij {
		boodschap()
	} else {
		fmt.Println("U heeft helaas geen toegang tot het parkeerterrein")
	}
}
