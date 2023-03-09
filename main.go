package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// Controleren of er juist één argument is opgegeven
	if len(os.Args) != 2 {
		log.Fatal("Geef het aantal keren op dat de tekst weergegeven dient te worden.")
	}

	// Converteer het argument naar een integer
	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal("Ongeldige invoer. Geef een positief geheel getal op.")
	}

	// Controleren of het getal positief is
	if num <= 0 {
		log.Fatal("Ongeldige invoer. Geef een positief geheel getal op.")
	}

	// Weergeven van "Alarm x!" zoveel keer als opgegeven
	for i := 1; i <= num; i++ {
		fmt.Printf("Alarm %d!\n", i)
	}
}
