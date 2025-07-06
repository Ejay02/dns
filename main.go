package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// Parse command line arguments
	hostPtr := flag.String("host", "", "Hostname to look up")
	typePtr := flag.String("type", "A", "DNS record type (A, AAAA, MX, TXT, CNAME, NS, PTR)")
	flag.Parse()

	if *hostPtr == "" {
		fmt.Println("Error: host parameter is required")
		flag.Usage()
		os.Exit(1)
	}

	// Convert type to uppercase for case-insensitive comparison
	recordType := strings.ToUpper(*typePtr)

	// Perform DNS lookup based on record type
	switch recordType {
	case "A":
		records, err := net.LookupHost(*hostPtr)
		if err != nil {
			fmt.Printf("Error looking up A records for %s: %v\n", *hostPtr, err)
			os.Exit(1)
		}
		fmt.Printf("A records for %s:\n", *hostPtr)
		for _, record := range records {
			fmt.Println(record)
		}

	case "AAAA":
		records, err := net.LookupIP(*hostPtr)
		if err != nil {
			fmt.Printf("Error looking up AAAA records for %s: %v\n", *hostPtr, err)
			os.Exit(1)
		}
		fmt.Printf("AAAA records for %s:\n", *hostPtr)
		for _, record := range records {
			if record.To4() == nil { // Only IPv6 addresses
				fmt.Println(record)
			}
		}

	case "MX":
		records, err := net.LookupMX(*hostPtr)
		if err != nil {
			fmt.Printf("Error looking up MX records for %s: %v\n", *hostPtr, err)
			os.Exit(1)
		}
		fmt.Printf("MX records for %s:\n", *hostPtr)
		for _, record := range records {
			fmt.Printf("%s (pref %d)\n", record.Host, record.Pref)
		}

	case "TXT":
		records, err := net.LookupTXT(*hostPtr)
		if err != nil {
			fmt.Printf("Error looking up TXT records for %s: %v\n", *hostPtr, err)
			os.Exit(1)
		}
		fmt.Printf("TXT records for %s:\n", *hostPtr)
		for _, record := range records {
			fmt.Println(record)
		}

	case "CNAME":
		record, err := net.LookupCNAME(*hostPtr)
		if err != nil {
			fmt.Printf("Error looking up CNAME for %s: %v\n", *hostPtr, err)
			os.Exit(1)
		}
		fmt.Printf("CNAME for %s:\n%s\n", *hostPtr, record)

	case "NS":
		records, err := net.LookupNS(*hostPtr)
		if err != nil {
			fmt.Printf("Error looking up NS records for %s: %v\n", *hostPtr, err)
			os.Exit(1)
		}
		fmt.Printf("NS records for %s:\n", *hostPtr)
		for _, record := range records {
			fmt.Println(record.Host)
		}

	default:
		fmt.Printf("Unsupported record type: %s\n", recordType)
		fmt.Println("Supported types: A, AAAA, MX, TXT, CNAME, NS")
		os.Exit(1)
	}
}
