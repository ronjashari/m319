package main

import "fmt"

func main() {
	// AUFGABE: Fügen Sie Ihre Variablen hier ein!

	var anzahlSiegel = 7              // int
	var kristallReinheit = 0.85       // float64
	var istAktiviert = true           // bool
	var zauberspruch = "Lumen Logica" // string

	// AB HIER NICHTS MEHR AENDERN!
	// Aufruf der Prüffunktion inklusive des Zauberspruchs
	pruefeTor(anzahlSiegel, kristallReinheit, istAktiviert, zauberspruch)
}

// pruefeTor nutzt Kontrollstrukturen, um das Schicksal der Helden zu bestimmen
func pruefeTor(siegel int, reinheit float64, status bool, spruch string) {
	if siegel == 7 && reinheit >= 0.8 && status == true && spruch == "Lumen Logica" {
		fmt.Println("Die Kristalle leuchten hell auf und der Code vibriert...")
		fmt.Println("Das Tor schwingt auf!")
	} else {
		fmt.Println("Ein tiefes Grollen ertönt... Das Tor bleibt fest verschlossen.")

		// Logik-Hilfe
		if siegel != 7 {
			fmt.Println("Hinweis: Die Anzahl der Siegel ist inkorrekt.")
		}
		if reinheit < 0.8 {
			fmt.Println("Hinweis: Die Energie der Kristalle ist zu schwach.")
		}
		if !status {
			fmt.Println("Hinweis: Das System wurde nicht aktiviert.")
		}
		if spruch != "Lumen Logica" {
			fmt.Println("Hinweis: Der Zauberspruch verhallt wirkungslos. Die Syntax des Wortes ist falsch.")
		}
	}
}
