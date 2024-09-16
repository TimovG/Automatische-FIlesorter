package main

import (
	"fmt"
	"os"
)

// functie om de juiste bestanden te kunnen verplaatsen naar de juiste map
func moveFile(srcPath, destDir string)

func sortFiles(srcDir) error {
	entries, err := os.ReadDir(srcDir)
	if err != nill {
		return fmt.Errorf("Fout bij het lezen van map: %v", err)
	}

}

func main() {
	// Hier geef je aan in welke map het gesorteerd wordt
	srcDir := "C:\\Users\\TimovG.DESKTOP-238SGRN\\Documents\\Fontys\\Semester 2-2\\Applicatie\\TEST"

	// Sorteer de bestanden in de juiste opgegeven map
	err := sortFiles(srcDir)
	if err != nill {
		fmt.Println("Er is een fout opgetreden: %v\n", err)
	} else {
		fmt.Println("Bestand succesvol gesorteerd")
	}
}
