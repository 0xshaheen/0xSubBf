# Subdomain Bruteforce Tool

**Author**: 0xShaheen

## Overview

This Subdomain Bruteforce Tool is a command-line utility designed for identifying live subdomains of a given domain through DNS resolution. It is written in Go and uses concurrent workers to efficiently check the availability of subdomains.

## Features

- Single domain or multiple domains from a list can be checked.
- Utilizes DNS resolution to identify live subdomains.
- Concurrent workers for improved performance.

## Usage

### Prerequisites

Make sure you have Go installed on your machine.

### Installation

1. Clone the repository:

   ```
   git clone https://github.com/0xshaheen/0xSubBf.git
2. Navigate to the project directory:
    ```bash
    cd 0xSubBf
3. Build the executable:
    ```
    go build subbf.go
    
### Usage Example

1. Single domain:
    ```
    ./0xSubBf -i name.txt -d google.com -o output.txt
2. Multiple domains from a list:
    ```
    ./0xSubBf -i name.txt -d domains.txt -o output.txt
### Options
```
-d: Specify a single domain to check.
-dl: Specify a file containing a list of domains.
-i: Specify the input file containing subdomains to check.
-o: Specify the output file to save found subdomains.
```
### Acknowledgments
``` 
This tool was created by 0xShaheen.
Inspired by the need for an efficient subdomain enumeration tool.
```


