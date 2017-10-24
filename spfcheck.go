package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		var hasSPF bool
		var spfRecord string

		domain := scanner.Text()
		records, err := net.LookupTXT(domain)
		if err != nil {
			log.Printf("Error: %v\n", err)
		}

		for _, record := range records {
			if strings.HasPrefix(record, "v=spf1") {
				hasSPF = true
				spfRecord = record
			}
		}

		fmt.Printf("%v,%v,%v\n", domain, hasSPF, spfRecord)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error: Could not read from input: %v\n", err)
	}
}
