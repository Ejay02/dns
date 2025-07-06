# DNS Lookup CLI Tool

A command-line DNS lookup utility written in Go that supports multiple record types.

## Features

- Lookup various DNS record types (A, AAAA, MX, TXT, CNAME, NS)
- Simple command-line interface
- Clear output formatting
- Error handling and validation

## Installation

1. Ensure you have [Go installed](https://golang.org/doc/install)
2. Clone this repository
3. Build and run:

```bash
go build -o dns-lookup
./dns-lookup -host example.com -type MX
```

Or run directly:
```bash
go run main.go -host example.com -type MX
```

## Usage

```
dns-lookup -host <hostname> [-type <record-type>]
```

### Options

- `-host`: Domain name to query (required)
- `-type`: DNS record type (default: A)
  - Supported types: A, AAAA, MX, TXT, CNAME, NS

### Examples

```bash
# Lookup A records (IPv4 addresses)
dns-lookup -host google.com

# Lookup MX records (mail servers)
dns-lookup -host microsoft.com -type MX

# Lookup TXT records
dns-lookup -host github.com -type TXT
```

## License

MIT