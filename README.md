# mailcheck

Tool for checking the existence of MX, SPF and DMARC records for domain names.

## Prerequisites

- [Go](https://golang.org/)

## Installation

```
$ go get -v github.com/dstotijn/mailcheck
```

## Usage

Outputting results to stdout:

```
$ mailcheck < domains.txt
```

Outputting results to file:

```
$ mailcheck < domains.txt > output.csv
```

### Reading result output

The first line has CSV-headers:

`domain,hasMX,hasSPF,spfRecord,hasDMARC,dmarcRecord`

Each subsequent line has 6 comma separated fields:

- `domain`: Domain name
- `hasMX`: Domain has at least one MX record? (`true` or `false`)
- `hasSPF`: Domain has SPF record? (`true` or `false`)
- `spfRecord`: SPF record contents
- `hasDMARC`: Domain has DMARC record? (`true` or `false`)
- `dmarcRecord`: DMARC record contents
