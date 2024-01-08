package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"os/exec"
	"sync"
)

var (
	clear   = exec.Command("clear") // clear command to clear the terminal
	domains []string
)

const (
	green   = "\033[92m"
	reset   = "\033[0m"
	red     = "\033[91m"
	workers = 1000
)

func clearTerminal() {
	clear.Stdout = os.Stdout
	_ = clear.Run()
}

func checkPing(host string) bool {
	_, err := net.LookupHost(host)
	return err == nil
}

func checkSubdomain(subdomain, domain, outputFile string, wg *sync.WaitGroup) {
	defer wg.Done()
	fullDomain := fmt.Sprintf("%s.%s", subdomain, domain)
	if checkPing(fullDomain) {
		saveToFile(outputFile, fullDomain)
		fmt.Printf("[ %sLIVE%s ]:%s%s%s\n", red, reset, green, fullDomain, reset)
	}
}

func saveToFile(outputFile, subdomain string) {
	file, err := os.OpenFile(outputFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error opening or creating the output file: %s\n", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(subdomain + "\n")
	if err != nil {
		fmt.Printf("Error writing to the output file: %s\n", err)
	}
}

func checkAndSaveSubdomains(domain, subdomainsFile, outputFile string) {
	file, err := os.Open(subdomainsFile)
	if err != nil {
		fmt.Printf("Error opening subdomains file: %s\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var wg sync.WaitGroup

	for scanner.Scan() {
		subdomain := scanner.Text()
		wg.Add(1)
		go checkSubdomain(subdomain, domain, outputFile, &wg)
	}

	wg.Wait()
}

func main() {
	clearTerminal()
	fmt.Println("=====================================")
	fmt.Println("           Sub_Domain Bruteforce     ")
	fmt.Println("            Author 0xShaheen         ")
	fmt.Println("=====================================")

	var domain string
	var domainList string
	var input string
	var output string

	flag.StringVar(&domain, "d", "", "Single domain to check")
	flag.StringVar(&domainList, "dl", "", "File containing a list of domains")
	flag.StringVar(&input, "i", "", "Input subdomains file")
	flag.StringVar(&output, "o", "", "Output file to save found subdomains")
	flag.Parse()

	if input == "" || output == "" {
		fmt.Println("Please provide input file (-i) and output file (-o)")
		return
	} else if domain != "" && domainList != "" {
		fmt.Println("Please provide either a single domain (-d) or a file containing a list of domains (-dl), not both.")
		return
	} else if domain == "" && domainList == "" {
		fmt.Println("Please provide either a single domain (-d) or a file containing a list of domains (-dl)")
		return
	}

	if domainList != "" {
		file, err := os.Open(domainList)
		if err != nil {
			fmt.Printf("Error opening domain list file: %s\n", err)
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			domain := scanner.Text()
			domains = append(domains, domain)
		}
	} else if domain != "" {
		domains = append(domains, domain)
	}

	for _, domain := range domains {
		checkAndSaveSubdomains(domain, input, output)
	}
}
