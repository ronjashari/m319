package main

import "fmt"

func main() {
	fmt.Println("=== Logik-Viadukt von Null-Pointer-City ===")

	// AUFGABE 1: Initialisierung der Umgebungsvariablen
	// Erstellen Sie drei Variablen mit den passenden Datentypen:
	// 1. 'wasserstand': Ein Kommazahl, der den Füllstand in Prozent angibt.
	// 2. 'temperatur': Ein ganzzahliger Wert, für die Grad Celsius.
	// 3. 'stadtviertel': Ein Textwert, der das Ziel angibt (z.B. "Markt", "Wohngebiet" oder "Oasenpark").

	var wasserstand float64 = 75.0
	var temperatur int = 28
	var stadtviertel string = "Markt"

	fmt.Printf("\nStatus-Check: %.1f%% Wasser, %d°C Außentemperatur\n", wasserstand, temperatur)
	fmt.Println("-------------------------------------------------")

	// AUFGABE 2: Separate Sicherheits-Checks
	// Das System soll nur funktionieren, wenn es sicher ist.
	// Standardmäßig setzen wir die Sicherheit auf 'true'.
	sicherheitGewaehrleistet := true

	// Setzen Sie hier die korrekte Kontrollbedingung für die If-Abfrage ein.
	// Die Warnung soll erscheinen, wenn der 'wasserstand' GRÖSSER als 90 ist.
	if wasserstand > 90 {
		fmt.Println("[WARNUNG]: Wasserstand kritisch! Überlauf-Gefahr.")
		sicherheitGewaehrleistet = false
	}

	// Fügen Sie hier um das Print Statement herum eine If-Abfrage ein.
	// Die Bedingung soll prüfen, ob die 'temperatur' GRÖSSER als 30 ist.
	// Falls ja, geben Sie die untenstehende Warnung aus und setzen Sie 'sicherheitGewaehrleistet' auf false.
	if temperatur > 30 {
		fmt.Println("[WARNUNG]: Hitze-Alarm! Hohe Verdunstungsrate.")
		sicherheitGewaehrleistet = false
	}

	// AUFGABE 3: Wasserverteilung nur bei Sicherheit
	// Prüfen Sie, ob 'sicherheitGewaehrleistet' wahr (true) ist.
	// Falls ja: Nutzen Sie eine Switch-Anweisung für die Variable 'stadtviertel'.
	// - Case "Markt": "Wasser zum Markt geleitet."
	// - Case "Wohngebiet": "Wasser zum Wohngebiet geleitet."
	// - Default: "Unbekanntes Viertel!"
	// Falls die Sicherheit NICHT gewährleistet ist, geben Sie "System gesperrt!" aus.
	if sicherheitGewaehrleistet {
		switch stadtviertel {
		case "Markt":
			fmt.Println("Wasser zum Markt geleitet.")
		case "Wohngebiet":
			fmt.Println("Wasser zum Wohngebiet geleitet.")
		default:
			fmt.Println("Unbekanntes Viertel!")
		}
	} else {
		fmt.Println("System gesperrt!")
	}
	fmt.Println("-------------------------------------------------")
	fmt.Println("=== Viadukt-Steuerung beendet ===")
}
