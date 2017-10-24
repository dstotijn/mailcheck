# spfcheck

Tool for checking the existence of a SPF record for domain names.

## Prerequisites

* [Go](https://golang.org/)

## Installation

```
$ go get -v github.com/dstotijn/spfcheck
```

## Usage

Outputting results to stdout:
```
$ spfcheck < domains.txt
```

Outputting results to file:
```
$ spfcheck < domains.txt > output.csv
```

### Reading results

Each line has 3 comma separated fields, first one is the domain name, second
is the result of the SPF check (`true` or `false`) and the last field has the
SPF record contents.
