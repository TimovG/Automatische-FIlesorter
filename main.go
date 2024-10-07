package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Folder struct {
	FolderName     string   `json:"folderName"`
	FileExtensions []string `json:"fileExtensions"`
}

type Config struct {
	Folders []Folder `json:"folders"`
}

func loadConfig(filePath string) (Config, error) {
	configFile, err := os.Open(filePath)
	if err != nil {
		return Config{}, fmt.Errorf("fout bij het openen van json bestand: %v", err)
	}
	defer configFile.Close()

	var config Config
	err = json.NewDecoder(configFile).Decode(&config)
	if err != nil {
		return Config{}, fmt.Errorf("fout bij het decoderen van het JSON-bestand: %v", err)
	}

	return config, nil
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

	// Verplaats het bestand naar de juiste map
	err := os.Rename(srcPath, destPath)
	if err != nil {
		return fmt.Errorf("fout bij het verplaatsen van bestand: %v", err)
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
			ext := strings.ToLower(filepath.Ext(entry.Name()))

			// Controleer bij alle gedefinieerde mappen
			for _, folder := range config.Folders {
				for _, fileExt := range folder.FileExtensions {
					if ext == fileExt {
						destDir := filepath.Join(srcDir, folder.FolderName)
						err := moveFile(srcPath, destDir)
						if err != nil {
							return err
						}
						break
					}
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
