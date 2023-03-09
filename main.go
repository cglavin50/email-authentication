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
	fmt.Printf("domain, hasMX, hasSPF, spfRecord, hasDMARC, demarcRecord\n")

	for scanner.Scan() {
		checkDomain(scanner.Text())
		fmt.Printf("domain, hasMX, hasSPF, spfRecord, hasDMARC, demarcRecord\n")
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}

func checkDomain(domain string) {
	var hasMX, hasSPF, hasDMARC bool
	var spfRec, dmarcRec string

	// checking mx records
	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		log.Println(err)
	}
	if len(mxRecords) > 0 {
		hasMX = true
	}

	//using DNS txt to lookup SPF records
	txtRec, err := net.LookupTXT(domain)
	if err != nil {
		panic(err)
	}

	for _, record := range txtRec {
		if strings.HasPrefix(record, "v=spf") {
			hasSPF = true
			spfRec = record
			break
		}
	}

	//
	dmarcRecords, err := net.LookupTXT("_demarc" + domain)
	if err != nil {
		log.Println(err)
	}

	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRec = record
			break
		}
	}

	if !hasDMARC && !hasMX && !hasSPF {
		fmt.Printf("No records found for this domain: %v\n", domain)
	} else {
		fmt.Printf("%v, %v, %v, %v, %v, %v\n\n", domain, hasMX, hasSPF, spfRec, hasDMARC, dmarcRec)
	}
}
