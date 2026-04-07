// definiert ... (unleserlich)
package main

// importiert ... (unleserlich)
import "fmt"

// =================================================================
// ARCHITEKTUR-BEREICH (Fortgeschrittene Magie - Bitte nicht ändern)
// =================================================================

// Fahrzeug ist die Basisklasse (Struktur) für alle Fortbewegungsmittel
type Fahrzeug struct {
	Name string
}

// Barke erbt Eigenschaften vom Fahrzeug und erweitert diese
type Barke struct {
	Fahrzeug
	Tragkraft  float64
	Passagiere int
}

// Methode zum Materialisieren der Barke
func (b Barke) beschwoere() {
	fmt.Printf("Die Barke '%s' beginnt zu schimmern...\n", b.Name)
	// Ein Loop der alten Zivilisation zur Visualisierung des Code-Rumpfes der 3x "DATA-STREAM-STABILIZING" ausgibt
	for i := 0; i < 3; i++ {
		fmt.Println("... [DATA-STREAM-STABILIZING] ...")
	}
}

// =================================================================
// IHRE AUFGABE: DIE LOGIK-WERKSTATT
// =================================================================

// Die .... (unleserlich)
func main() {
	fmt.Println("=== Die Beschwörung der Logik-Barke ===")

	// AUFGABE 1: Variablen-Initialisierung
	// TODO: Deklarieren Sie die Variable 'tragkraft' als float64 mit dem Wert 500.5
	var tragkraft float64 = 500.5
	// TODO: Deklarieren Sie die Variable 'anzahlPassagiere' als int mit dem Wert 5
	var anzahlPassagiere int = 5

	// Erklären Sie hier kurz mit einem einzeiligen Kommentar,
	// warum float64 für die Tragkraft gewählt wurde:
	// Weil die Gewichtswerte Dezimalstellen haben

	/* FIXME: Der nachfolgende Block berechnet die Stabilität.
	   Die Logik der alten Meister ist korrupt.
	   Bedingung: Wenn die anzahlPassagiere größer als 6 ODER
	   die tragkraft kleiner als 400 ist, bricht die Beschwörung ab.
	*/

	istStabil := false

	// Füge hier nachfolgend die logische Bedingung ein
	if anzahlPassagiere > 6 || tragkraft < 400 {
		fmt.Println("[WARNUNG]: Die Logik-Parameter sind instabil!")
		istStabil = false
	} else {
		fmt.Println("[CHECK]: Auftriebs-Logik stabil.")
		istStabil = true
	}

	// .... (unleserlich)
	fmt.Println("---------------------------------------")

	// AUFGABE 2: Navigationstyp festlegen
	// Nutzen Sie einen Switch, um je nach 'anzahlPassagiere' einen Modus auszugeben:
	// 1-3 Passagiere: "Modus: Kleines Boot"
	// 4-6 Passagiere: "Modus: Gruppen-Barke"
	// Default: "Modus: Unbekannt"

	// TODO: Fügen Sie hier den Switch-Block ein
	switch {
	case anzahlPassagiere >= 1 && anzahlPassagiere <= 3:
		fmt.Println("Modus: Kleines Boot")
	case anzahlPassagiere >= 4 && anzahlPassagiere <= 6:
		fmt.Println("Modus: Gruppen-Barke")
	default:
		fmt.Println("Modus: Unbekannt")
	}

	// AUFGABE 3: Fügen Sie hier einen mehrzeiligen Kommentar ein, der das Switch Statement oben erklärt.
	/*Das Switch-Statement bestimmt den Fahrmodus der Barke
	  basierend auf der Anzahl der Passagiere.
	  - Bei 1 bis 3 Passagieren wird "Kleines Boot" ausgegeben.
	  - Bei 4 bis 6 Passagieren wird "Gruppen-Barke" ausgegeben.
	  - Bei allen anderen Werten greift der Default-Fall: "Unbekannt".
	*/

	// FINALER SCHRITT: Wenn alles stabil ist, wird die Barke erzeugt
	if istStabil {
		// Hier wird ein neues Objekt der Klasse 'Barke' erstellt
		meineBarke := Barke{
			Fahrzeug:   Fahrzeug{Name: "Lumen Logica"},
			Tragkraft:  tragkraft,
			Passagiere: anzahlPassagiere,
		}

		meineBarke.beschwoere()
		fmt.Println("Erfolg: Die Barke ist bereit für das Bytestromdelta!")
	} else {
		fmt.Println("Fehler: Die Beschwörung ist fehlgeschlagen.")
	}
}
