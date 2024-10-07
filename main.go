package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type Config struct {
	Afbeeldingen []string `json:"Afbeeldingen"`
	Videos       []string `json:"Videos"`
	Documents    []string `json:"Documents"`
}

func loadConfig(filePath string) (Config, error) {
	var config Config

	// Open en lees het JSON-bestand
	configFile, err := os.Open(filePath)
	if err != nil {
		return config, fmt.Errorf("Fout bij het openen van JSON-bestand: %v", err)
	}
	defer configFile.Close()

	// Decodeer het JSON-bestand naar de Config struct
	err = json.NewDecoder(configFile).Decode(&config)
	if err != nil {
		return config, fmt.Errorf("Fout bij het decoderen van het JSON-bestand: %v", err)
	}

	return config, nil
}

func getTargetDirectory(config Config, srcDir, fileName string) string {
	ext := strings.ToLower(filepath.Ext(fileName))

	for _, imageExt := range config.Afbeeldingen {
		if ext == imageExt {
			return filepath.Join(srcDir, "Afbeeldingen")
		}
	}
	for _, videoExt := range config.Videos {
		if ext == videoExt {
			return filepath.Join(srcDir, "Videos")
		}
	}
	for _, docExt := range config.Documents {
		if ext == docExt {
			return filepath.Join(srcDir, "Documents")
		}
	}
	// Als geen match, return een lege string
	return ""
}

// Functie om een bestand te kopiëren van huidige plek naar bedoelde plek
func copyFile(srcPath, destPath string) error {
	srcFile, err := os.Open(srcPath)
	if err != nil {
		return fmt.Errorf("fout bij het openen van bestand: %v", err)
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

	// Kopieer het bestand naar de juiste map
	err := copyFile(srcPath, destPath)
	if err != nil {
		return err
	}

	// Verwijder het orgiginele nadat het is gekopieerd
	err = os.Remove(srcPath)
	if err != nil {
		return fmt.Errorf("fout bij het verwijderen van bronbestand: %v", err)
	}

	fmt.Printf("Bestand verplaatst naar: %s\n", destPath)
	return nil
}

// functie om de bestanden te sorteren naar de juiste map
func sortFiles(srcDir string, config Config) error {
	entries, err := os.ReadDir(srcDir)
	if err != nil {
		return fmt.Errorf("fout bij het lezen van map: %v", err)
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			srcPath := filepath.Join(srcDir, entry.Name())

			destDir := getTargetDirectory(config, srcDir, entry.Name())
			if destDir != "" {
				err := moveFile(srcPath, destDir)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func main() {
	// Hier kan je aangeven in welke map alles gesorteerd moet worden
	srcDir := "C:\\Users\\TimovG.DESKTOP-238SGRN\\Documents\\Fontys\\Semester 2-2\\Applicatie\\Automatische-FIlesorter\\Test"

	configFilePath := "config.json"

	config, err := loadConfig(configFilePath)
	if err != nil {
		fmt.Printf("Er is een fout opgetreden bij het laden van de configuratie: %v\n", err)
		return
	}

	err = sortFiles(srcDir, config)
	if err != nil {
		fmt.Printf("Er is een fout opgetreden: %v\n", err)
	} else {
		fmt.Println("Bestanden succesvol gesorteerd.")
	}
}
