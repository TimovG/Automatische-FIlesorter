package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// Functie om een bestand te kopiëren van bron naar bestemming
func copyFile(srcPath, destPath string) error {
	srcFile, err := os.Open(srcPath)
	if err != nil {
		return fmt.Errorf("fout bij het openen van bronbestand: %v", err)
	}
	defer srcFile.Close()

	destFile, err := os.Create(destPath)
	if err != nil {
		return fmt.Errorf("fout bij het maken van doelbestand: %v", err)
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		return fmt.Errorf("fout bij het kopiëren van bestand: %v", err)
	}

	return nil
}

// Functie om de juiste bestanden te verplaatsen naar de juiste map
func moveFile(srcPath, destDir string) error {
	// Zorgt ervoor dat de doelmap bestaat
	if _, err := os.Stat(destDir); os.IsNotExist(err) {
		err = os.MkdirAll(destDir, 0755)
		if err != nil {
			return fmt.Errorf("fout bij het maken van map: %v", err)
		}
	}

	destPath := filepath.Join(destDir, filepath.Base(srcPath))

	// Kopieer het bestand naar de doelmap
	err := copyFile(srcPath, destPath)
	if err != nil {
		return err
	}

	// Verwijder het bronbestand nadat het is gekopieerd
	err = os.Remove(srcPath)
	if err != nil {
		return fmt.Errorf("fout bij het verwijderen van bronbestand: %v", err)
	}

	fmt.Printf("Bestand verplaatst naar: %s\n", destPath)
	return nil
}

func sortFiles(srcDir string) error {
	entries, err := os.ReadDir(srcDir)
	if err != nil {
		return fmt.Errorf("Fout bij het lezen van map: %v", err)
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			srcPath := filepath.Join(srcDir, entry.Name())
			ext := strings.ToLower(filepath.Ext(entry.Name()))

			// Hier wordt bepaald welk bestandstype waar moet komen
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
	srcDir := "C:\\Users\\TimovG.DESKTOP-238SGRN\\Documents\\Fontys\\Semester 2-2\\Applicatie\\Automatische-FIlesorter\\Test"

	err := sortFiles(srcDir)
	if err != nil {
		fmt.Printf("Er is een fout opgetreden: %v\n", err)
	} else {
		fmt.Println("Bestanden succesvol gesorteerd.")
	}
}
