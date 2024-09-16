package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// functie om de juiste bestanden te kunnen verplaatsen naar de juiste map
func moveFile(srcPath, destDir string) error {
	// Zorgt ervoor dat de doelmap bestaat
	if _, err := os.Stat(destDir); os.IsNotExist(err) {
		err = os.MkdirAll(destDir, 0755)
		if err != nil {
			return fmt.Errorf("fout bij het maken van map: %v", err)
		}
	}

	destPath := filepath.Join(destDir, filepath.Base(srcPath))

	err := os.Rename(srcPath, destPath)
	if err != nil {
		return fmt.Errorf("fout bij het verplaatsen van bestand: %v", err)
	}

	fmt.Printf("Bestand verplaatst naar: %s\n", destPath)
	return nil
}

func sortFiles(srcDir string) error {
	entries, err := os.ReadDir(srcDir)
	if err != nill {
		return fmt.Errorf("Fout bij het lezen van map: %v", err)
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			srcPath := filepath.Join(srcDir, entry.Name())
			ext := strings.ToLower(filepath.Ext(entry.Name()))

			// Hier wordt bepaal welke betandstype waar moet komen
			switch ext {
			case ".jpg", ".jpeg", ".png", ".gif":
				moveFile(srcPath, filepath.Join(srcDir, "Afbeeldingen"))
			case ".mp4":
				moveFile(srcPath, filepath.Join(srcDir, "Videos"))
			}
		}
	}

	return nil
}

func main() {
	// Hier kan je aangeven in welke map alles gesorteerd moet worden
	srcDir := "C:\\Users\\TimovG.DESKTOP-238SGRN\\Documents\\Fontys\\Semester 2-2\\Applicatie\\TEST"

	err := sortFiles(srcDir)
	if err != nil {
		fmt.Printf("Er is een fout opgetreden: %v\n", err)
	} else {
		fmt.Println("Bestanden succesvol gesorteerd.")
	}
}
