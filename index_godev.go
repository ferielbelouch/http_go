package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"text/tabwriter"
)

// Module représente la structure d'un module dans la réponse JSON
type Module struct {
	Path    string `json:"path"`
	Version string `json:"version"`
}

// ForgeStats représente les statistiques d'une forge
type ForgeStats struct {
	Forge       string
	Modules     int
	Versions    int
	VersionList []string
}

func main() {
	client := http.Client{}
	req, err := http.NewRequest("GET", "https://index.golang.org/index", nil)
	if err != nil {
		fmt.Println("Erreur lors de la création de la requête HTTP:", err)
		os.Exit(1)
	}

	req.Header.Set("Disable-Module-Fetch", "true")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Erreur lors de l'envoi de la requête HTTP:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// Lire la réponse JSON ligne par ligne
	var modules []Module
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		line := scanner.Text()

		// Désérialiser chaque ligne JSON individuellement
		var module Module
		err := json.Unmarshal([]byte(line), &module)
		if err != nil {
			fmt.Println("Erreur lors de la désérialisation JSON:", err)
			os.Exit(1)
		}

		modules = append(modules, module)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Erreur lors de la lecture de la réponse HTTP:", err)
		os.Exit(1)
	}

	// Afficher uniquement la liste des modules et versions
	modules, forgeStats := calculateForgeStats(modules)
	printStats(modules, forgeStats)
}

func printStats(modules []Module, forgeStats map[string]ForgeStats) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.AlignRight|tabwriter.Debug)
	defer w.Flush()

	// Header du tableau
	fmt.Fprintln(w, "Forge\tModule\tVersion")

	// Afficher la liste des modules et versions
	for _, module := range modules {
		fmt.Fprintf(w, "%s\t%s\t%s\n", getForgeFromPath(module.Path), module.Path, module.Version)
	}

	// Afficher les statistiques par forge
	fmt.Println("\nStatistiques par Forge:")
	fmt.Fprintln(w, "Forge\tModules\tVersions")
	for forge, stats := range forgeStats {
		fmt.Fprintf(w, "%s\t%d\t%d\n", forge, stats.Modules, stats.Versions)
	}
}

func calculateForgeStats(modules []Module) ([]Module, map[string]ForgeStats) {
	forgeStats := make(map[string]ForgeStats)

	// Calculer les statistiques par forge
	for _, module := range modules {
		forge := getForgeFromPath(module.Path)
		stats, exists := forgeStats[forge]
		if !exists {
			stats.Forge = forge
		}
		stats.Modules++
		if !versionExists(stats, module.Version) {
			stats.Versions++
			stats.VersionList = append(stats.VersionList, module.Version)
		}
		forgeStats[forge] = stats
	}

	return modules, forgeStats
}

func getForgeFromPath(path string) string {
	// Extraire le nom de la forge à partir du champ "path"
	// Diviser le chemin par "/" et prendre le premier élément
	forgeParts := strings.Split(path, "/")
	if len(forgeParts) > 0 {
		return forgeParts[0]
	}
	return ""
}

func versionExists(stats ForgeStats, version string) bool {
	// Vérifie si la version existe déjà dans les statistiques
	for _, v := range stats.VersionList {
		if v == version {
			return true
		}
	}
	return false
}
