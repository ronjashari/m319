package main

import "fmt"

/*
 * TEIL 1: Die Basis-Auftriebsfunktion
 */
func berechneBasisAuftrieb(windgeschwindigkeit int, anflugwinkel int) int {
	ergebnis := windgeschwindigkeit + anflugwinkel
	return ergebnis
}

/*
 * TEIL 2: Der Energieverbrauch
 */
func berechneEnergieVerbrauch(gewicht int, strecke int) int {
	return (gewicht * strecke) / 10
}

/*
 * TEIL 3: Die finale Flug-Validierung (Multiple Return)
 */
func checkFlugStatus(auftrieb int, verbrauch int) (int, string) {
	differenz := auftrieb - verbrauch
	status := ""
	if differenz > 0 {
		status = "Stabil"
	} else {
		status = "Absturzgefahr"
	}
	return differenz, status
}

/*
 * TEIL 4: Das freudige Kreischen des Falkens
 */
func kreischen() {
	fmt.Println("KKKRREEEIIIISSSCCCHHHHH")
}

func main() {
	fmt.Println("--- Analyse der Variable-Falken startet ---")

	// 1. Aufruf der Basis-Funktion
	speed := 100
	winkel := 50
	aktuellerAuftrieb := berechneBasisAuftrieb(speed, winkel)
	fmt.Printf("Berechneter Auftrieb: %d Einheiten\n", aktuellerAuftrieb)

	// 2. Aufruf der Energie-Funktion
	gewicht := 10
	distanz := 10
	verbrauch := berechneEnergieVerbrauch(gewicht, distanz)
	fmt.Printf("Voraussichtlicher Verbrauch: %d Einheiten\n", verbrauch)

	// 3. Finale Prüfung (Auffangen von zwei Rückgabewerten)
	kraftReserve, flugZustand := checkFlugStatus(aktuellerAuftrieb, verbrauch)
	fmt.Println("-------------------------------------------")
	fmt.Printf("Analyse-Ergebnis: %s (Reserve: %d)\n", flugZustand, kraftReserve)
	if flugZustand == "Stabil" && kraftReserve > 0 {
		fmt.Println("Mission abgeschlossen: Der Falke hält seine Bahn!")
	} else {
		fmt.Println("Fehler: Die mathematische Kapselung ist noch instabil.")
	}

	// 4. Kreischen (3x aufrufen)
	kreischen()
	kreischen()
	kreischen()
}
