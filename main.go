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

	if len(os.Args) > 3 {
		fmt.Println("Parece que você está utilizando o programa de maneira errada.")
		fmt.Println("")
		helpMessage()
		os.Exit(0)
	}

	action := os.Args[1]
	switch action {
	case "help":
		helpMessage()
		os.Exit(0)
	case "about":
		aboutMessage()
		os.Exit(0)
	case "verify":
		domain := os.Args[2]
		checkDomain(domain)
		os.Exit(0)
	}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("domínio, possuiMX, possuiSPF, spf, possuiDMARC, DMARC\n")

	for scanner.Scan() {
		checkDomain(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Erro: %v\n", err.Error())
	}

}

func helpMessage() {
	fmt.Println("O programa aceita três possíveis comandos.")
	fmt.Println("verify\t para verificar domínios de email (necessário passar o domínio)")
	fmt.Println("about\t para entender mais sobre a ferramenta")
	fmt.Println("help\t para entender o funcionamento (essa mensagem)")
	fmt.Println("")
	fmt.Println("Execute novamente com um destes comandos.")
}

func aboutMessage() {
	fmt.Println("Sobre")
}

func checkDomain(domain string) {
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, DMARCRecord string

	fmt.Println("Verificando o domínio...")
	hasMX = verifyMXRecords(domain)
	hasSPF, spfRecord = verifySPFRecords(domain)
	hasDMARC, DMARCRecord = verifyDMARCRecords(domain)

	printResults(domain, hasMX, hasSPF, hasDMARC, spfRecord, DMARCRecord)
}

func verifyMXRecords(domain string) bool {
	mxRecords, err := net.LookupMX(domain)

	if err != nil {
		log.Printf("Erro: %v\n", err.Error())
		os.Exit(0)
	}

	if len(mxRecords) > 0 {
		return true
	}

	return false
}

func verifySPFRecords(domain string) (bool, string) {
	txtRecords, err := net.LookupTXT(domain)

	if err != nil {
		log.Printf("Erro: %v\n", err.Error())
	}

	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			return true, record
		}
	}

	return false, ""
}

func verifyDMARCRecords(domain string) (bool, string) {
	DMARCRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("Erro: %v\n", err.Error())
	}

	for _, record := range DMARCRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			return true, record
		}
	}

	return false, ""
}

func printResults(domain string, hasMX bool, hasSPF bool, hasDMARC bool, spfRecord string, DMARCRecord string) {
	fmt.Printf("\nDomínio: %v\n", domain)

	fmt.Println("")
	if hasMX {
		fmt.Println("Domínio possui um servidor de email para receber mensagens (MX Record).")
	} else {
		fmt.Println("Domínio não possui servidor de email (MX Record).")
	}

	fmt.Println("")
	if hasSPF {
		fmt.Println("Possui um SPF.")
		fmt.Printf("SPF: %v\n", spfRecord)
	} else {
		fmt.Println("Domínio não possui SPF.")
	}

	fmt.Println("")
	if hasDMARC {
		fmt.Println("Possui DMARC.")
		fmt.Printf("DMARC: %v\n", DMARCRecord)
	} else {
		fmt.Println("Domínio não possui DMARC.")
	}
}
